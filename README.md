# barcode-c128-svg

This is a fork of https://github.com/juliankoehn/barcode but with
everything removed that is not SVG generation of a Code 128 barcode. Also,
error generation has improved (the function returns an error instead of
printing to stderr).

Usage:

```
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
```

## Call:

* `code`: {string} Your Code
* `variant`: {string} one of Supported Barcodes ("", "A", "B", "C")
* `w`: {int} barcode with * w multiplier
* `h`: {int} height of the barcode in px
* `color`: {string} color as CSS compatible string value
* `showCode`: {bool} display code under BARCODE
* `inline`: {bool} removes XML/SVG headers from output

### Returns
SVG as `string`

```go
GetBarcodeSVG(code, variant, w, h, color, showCode, inline)
```
