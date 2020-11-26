// This program generates messages.go from messages.xml
// TODO generate encoder/decoder depending on fixed, optional and variable size
// TODO generate string methods for all bitfield types

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
	"sort"
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
	LenField    string `xml:"name,attr"` // for repeated fields: name of the count field
	Name        string

	// non-repeated, non-optional fields
	Offset      string
	Type        string
	Comment     string
	Scale       string
	Unit        string
	BitfieldRef string    `xml:"Bitfield>Reference"`
	Bitfield    []*BitDef `xml:"Bitfield>Type"`

	Nested []*Block `xml:"Block"` // for repeated or optional blocks, this contains the subfields

	Message *Message `xml:"-"` // link back up
	LenFor  string   `xml:"-"` // for fields that are the CountField of a repeated field, name of the repeated field

}

type BitDef struct {
	Index       string
	Type        string
	Name        string
	Description string

	Block *Block `xml:"-"` // link back up
}

func (b *BitDef) Mask() string {
	parts := strings.Split(b.Index, ":")
	if len(parts) == 2 {
		hi, _ := strconv.ParseUint(parts[1], 0, 8)
		lo, _ := strconv.ParseUint(parts[0], 0, 8)
		if hi <= lo {
			log.Fatalf("hi<=lo in bit mask %q", b.Index)
		}
		return fmt.Sprintf("0x%x", ((1<<(hi+1))-1)^((1<<(lo))-1))
	}
	i, _ := strconv.ParseUint(b.Index, 0, 8)
	return fmt.Sprintf("0x%x", 1<<i)
}

type byNameAndLength []*Message

func (v byNameAndLength) Len() int      { return len(v) }
func (v byNameAndLength) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v byNameAndLength) Less(i, j int) bool {
	if v[i].Name != v[j].Name {
		return v[i].Name < v[j].Name
	}
	return v[i].MinSize() < v[j].MinSize()
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

func (m *Message) ClassIDName() string {
	parts := strings.Split(strings.ToLower(m.Name), "-")
	if len(parts) > 3 {
		parts = parts[:3]
	}
	return strings.Join(parts, "-")
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
		return "Float8" // defined in spec but not found in any message
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

func (b *Block) FieldSize() int {
	parts := strings.Split(b.Type, "[")
	sz := 1
	if len(parts) == 2 {
		v, err := strconv.ParseUint(parts[1][:len(parts[1])-1], 10, 8)
		if err != nil {
			log.Fatalf("invalid array size %#v", b.Type)
		}
		sz = int(v)
	}
	switch parts[0] {
	case "RU1_3", "I1", "U1", "CH", "X1":
		return sz
	case "I2", "U2", "X2":
		return sz * 2
	case "R4", "I4", "U4", "X4":
		return sz * 4
	case "R8", "I8", "U8":
		return sz * 8
	}
	return 0 // probably invalid
}

// the non optional non repeated fields at the begining
func (m *Message) MinSize() int {
	sz := 0
	for _, v := range m.Blocks {
		if v.Cardinality != "" {
			break
		}
		sz += v.FieldSize()
	}
	return sz
}

// minsize + the optional bit
func (m *Message) MaxFixSize() int {
	sz := 0
	for _, v := range m.Blocks {
		switch v.Cardinality {
		case "":
			sz += v.FieldSize()
		case "optional":
			for _, vv := range v.Nested {
				sz += vv.FieldSize()
			}
		case "repeated":
		}
	}
	return sz
}

// size of the (hopefulluy single) variable block
func (m *Message) VarSize() int {
	sz := 0
	for _, v := range m.Blocks {
		if v.Cardinality == "repeated" {
			for _, vv := range v.Nested {
				sz += vv.FieldSize()
			}
		}
	}
	return sz
}

// the UBX-INF-xxx messages are really just strings
func (m *Message) IsString() bool {
	if len(m.Blocks) != 1 {
		return false
	}
	return m.Blocks[0].IsString()
}

func (b *Block) IsString() bool {
	if b.Cardinality != "repeated" {
		return false
	}
	if len(b.Nested) != 1 {
		return false
	}
	if b.Nested[0].Type != "CH" && b.Nested[0].Type != "U1" {
		return false
	}
	return true
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

	sort.Stable(byNameAndLength(definitions.Message))

	for _, v := range definitions.Message {
		for _, b := range v.Blocks {
			b.Link(v)
		}
	}

	// name repeated and annotate 'count' fields
	for _, msg := range definitions.Message {
		for _, b := range msg.Blocks {
			if b.Cardinality == "repeated" {
				b.Name = "Items"
				if strings.HasPrefix(strings.ToLower(b.LenField), "num") {
					b.Name = strings.Title(b.LenField[3:])
				}

				// link back
				for _, bb := range msg.Blocks {
					if bb.Name == b.LenField {
						bb.LenFor = b.Name
						break
					}
				}

			}
		}
	}

	// sort by class/id, and length options
	msgs := map[string][]*Message{}
	for _, v := range definitions.Message {
		n := v.ClassIDName()
		v.version = len(msgs[n])
		msgs[n] = append(msgs[n], v)
	}

	// figure out if we can figure out the type from just class, id and size
	for k, v := range msgs {
		m := map[int]int{}
		for _, vv := range v {
			m[vv.MinSize()]++
			if vv.MaxFixSize() != vv.MinSize() {
				m[vv.MaxFixSize()]++
			}
		}
		for _, vv := range v {
			if sz := vv.VarSize(); sz != 0 {
				for kk, _ := range m {
					if (kk-vv.MinSize())%sz == 0 {
						m[vv.MinSize()]++
					}
				}
			}
		}
		for _, cnt := range m {
			if len(v) > 1 && cnt > 1 {
				log.Println("Ambiguous type", k, m)
				isambiguous[k] = true
				break
			}
		}
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "// Generated Code -- DO NOT EDIT.\n//go:generate go run msggen.go %s %s %s\n\n", flag.Arg(0), flag.Arg(1), flag.Arg(2))
	if err := tmpl.Execute(&buf, msgs); err != nil {
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
	"lower":       strings.ToLower,
	"upper":       strings.ToUpper,
	"title":       strings.Title,
	"notabs":      notabs,
	"isambiguous": isAmbiguous,
}

var wstospace = strings.NewReplacer("\t", " ", "\n", " ")

func notabs(s string) string { return wstospace.Replace(s) }

var isambiguous = map[string]bool{}

func isAmbiguous(classid string) bool { return isambiguous[classid] }
