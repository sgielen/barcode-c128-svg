package main

import (
	"os"

	barcode "github.com/sgielen/barcode-c128-svg"
)

func main() {
	data, err := barcode.GetBarcodeSVG("ABCDEFG", "", 1, 50, "black", true, false)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("out.svg", []byte(data), 0644); err != nil {
		panic(err)
	}
}
