// This program generates messages.go from messages.xml

// +build ignore

package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Definitions struct {
	Message []*Message
}

type Message struct {
	Name        string
	Type        string
	Description string
	Comment     string
	Firmware    string
	Class       Number   `xml:"Structure>Class"`
	Id          Number   `xml:"Structure>Id"`
	Length      string   `xml:"Structure>Length"` // of the form A + N * B, but varying syntax
	Blocks      []*Block `xml:"Structure>Payload>Block"`

	version int // different versions of the same message name
}

type Block struct {
	Cardinality string `xml:"type,attr"` // repeated or optional, in which case 'nested' is non nil
	CountField  string `xml:"name,attr"` // for repeated fields: name of the count field

	// non-repeated, non-optional fields
	Offset   string
	Name     string
	Type     string
	Comment  string
	Scale    string
	Unit     string
	Bitfield []*BitDef `xml:"Bitfield>Type"`

	Nested []*Block `xml:"Block"` // for repeated or optional blocks, this contains the subfields

	Message *Message `xml:"-"` // link back up
}

type BitDef struct {
	Index       string
	Type        string
	Name        string
	Description string

	Block *Block `xml:"-"` // link back up
}

// set Block.Message and BitDef.Block pointers
func (b *Block) Link(m *Message) {
	b.Message = m
	for _, v := range b.Nested {
		v.Link(m)
	}
	for _, v := range b.Bitfield {
		v.Block = b
	}
}

func (m *Message) TypeName() string {
	parts := strings.Split(strings.ToLower(m.Name), "-")
	for i, v := range parts {
		parts[i] = strings.Title(v)
	}
	if m.version > 0 {
		return fmt.Sprintf("%s%d", strings.Join(parts[1:], ""), m.version)
	}
	return strings.Join(parts[1:], "")
}

var (
	reUnit       = regexp.MustCompile(`^[a-zA-Z^/2]+$`)
	reScaleDec   = regexp.MustCompile(`^1e-\d+$`)
	reScaleLeft  = regexp.MustCompile(`^2\^-\d+$`)
	reScaleRight = regexp.MustCompile(`^2\^\d+$`)
	repl         = strings.NewReplacer("^", "", "/", "_")
)

// Return the fieldname to use for a Go struct.
// if the units are not too wild they are suffixed as lowercase, with '/' converted to _
// if the scaling is 1e-d, or 2^d, or 2^-d, we suffix ed, ld or rd resp. (Daedalean convention on variable naming)
func (b *Block) FieldName() string {
	n := strings.Title(b.Name)
	if u := reUnit.FindString(b.Unit); u != "" {
		n = n + "_" + repl.Replace(strings.ToLower(u))
		if s := reScaleDec.FindString(b.Scale); s != "" {
			n = n + "e" + s[3:]
		} else if s := reScaleLeft.FindString(b.Scale); s != "" {
			n = n + "l" + s[3:]
		} else if s := reScaleRight.FindString(b.Scale); s != "" {
			n = n + "r" + s[2:]
		}
	}
	return n
}

func (b *Block) FieldType() string {
	tp := b.Type
	if i := strings.Index(tp, "["); i >= 0 {
		tp = tp[:i]
	}
	switch tp {
	case "RU1_3":
		return "Float8"
	case "R4":
		return "float32"
	case "R8":
		return "float64"
	case "I1":
		return "int8"
	case "U1", "CH", "X1":
		return "byte"
	case "I2":
		return "int16"
	case "U2", "X2":
		return "uint16"
	case "I4":
		return "int32"
	case "U4", "X4":
		return "uint32"
	case "I8":
		return "int64"
	case "U8":
		return "uint64"
	}
	return tp // probably invalid
}

func (b *Block) ArraySpec() string {
	if i := strings.Index(b.Type, "["); i >= 0 {
		return b.Type[i:]
	}
	return ""
}

// numeric fields in any base, not just decimal
type Number uint64

func (v *Number) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var f string
	if err := d.DecodeElement(&f, &start); err != nil {
		return err
	}
	vv, err := strconv.ParseUint(f, 0, 64)
	if err != nil {
		return err
	}
	*v = Number(vv)
	return nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("msggen: ")
	flag.Parse()

	if len(flag.Args()) != 3 {
		log.Fatalf("Usage: %s code.tmpl messages.xml code.go", os.Args[0])
	}

	tmpl, err := template.New(filepath.Base(flag.Arg(0))).Funcs(tmplfuncs).ParseFiles(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if flag.Arg(1) != "-" {
		os.Stdin, err = os.Open(flag.Arg(1))
		if err != nil {
			log.Fatal(err)
		}
	}

	var definitions Definitions

	if err := xml.NewDecoder(os.Stdin).Decode(&definitions); err != nil {
		log.Fatal(err)
	}

	for _, v := range definitions.Message {
		for _, b := range v.Blocks {
			b.Link(v)
		}
	}

	dups := map[string]int{}
	for _, v := range definitions.Message {
		v.version = dups[v.Name]
		dups[v.Name]++
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Generated Code -- DO NOT EDIT.\n//go:generate go run msggen.go %s %s %s\n\n", flag.Arg(0), flag.Arg(1), flag.Arg(2))
	if err := tmpl.Execute(&buf, definitions); err != nil {
		log.Fatal(err)
	}
	b := buf.Bytes()

	// fix all &#xx; xml entities
	b = regexp.MustCompile(`&#[0-9]{2};`).ReplaceAllFunc(b, func(b []byte) []byte {
		v, _ := strconv.ParseUint(string(b[2:4]), 10, 8)
		return []byte{byte(v)}
	})
	b = regexp.MustCompile(`&gt;`).ReplaceAllLiteral(b, []byte(">"))
	b = regexp.MustCompile(`&lt;`).ReplaceAllLiteral(b, []byte("<"))
	b = regexp.MustCompile(`&amp;`).ReplaceAllLiteral(b, []byte("&"))

	// try to format as valid Go
	bb, err := format.Source(b)
	if err != nil {
		log.Println(err)
		bb = b
	}

	if flag.Arg(2) != "-" {
		os.Stdout, err = os.Create(flag.Arg(2))
		if err != nil {
			log.Fatal(err)
		}
		defer os.Stdout.Close()
	}

	if _, err := os.Stdout.Write(bb); err != nil {
		log.Fatal(err)
	}

}

// Helper functions for in the template
var tmplfuncs = template.FuncMap{
	"lower":  strings.ToLower,
	"upper":  strings.ToUpper,
	"title":  strings.Title,
	"notabs": notabs,
	"mask":   mask,
}

var wstospace = strings.NewReplacer("\t", " ", "\n", " ")

func notabs(s string) string { return wstospace.Replace(s) }

func msgtypename(s string) string {
	parts := strings.Split(strings.ToLower(s), "-")
	for i, v := range parts {
		parts[i] = strings.Title(v)
	}
	return strings.Join(parts[1:], "")
}

func mask(s string) string {
	parts := strings.Split(s, ":")
	if len(parts) == 2 {
		hi, _ := strconv.ParseUint(parts[1], 0, 8)
		lo, _ := strconv.ParseUint(parts[0], 0, 8)
		return fmt.Sprintf("0x%x", ((1<<(hi+1))-1)^((1<<(lo))-1))
	}
	i, _ := strconv.ParseUint(s, 0, 8)
	return fmt.Sprintf("0x%x", 1<<i)
}
