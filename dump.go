// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/daedaleanai/ublox"
)

func main() {
	d := ublox.NewDecoder(os.Stdin)
	for {
		msg, err := d.Decode()
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("%v\n", msg)
	}
}
