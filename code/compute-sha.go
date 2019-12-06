// Computes sha of a given input

package main

import (
	"flag"
	"os"
	"io/ioutil"
	"log"
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	)

var width = flag.Int("w", 256, "hash width (256 or 512)")

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if *width == 256 {
		fmt.Printf("%x\n", sha256.Sum256([]byte(b)))
	} else if *width == 384 {
		fmt.Printf("%x\n", sha512.Sum384([]byte(b)))
	} else if *width == 512 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(b)))
	}
}
