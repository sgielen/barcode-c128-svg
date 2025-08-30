package barcode

import (
	"strconv"
)

type barArray struct {
	Code  string
	MaxW  int
	MaxH  int
	BCode []bCode
}

type bCode struct {
	T bool
	W int
	H int
	P int
}

// GetBarcodeSVG generates a SVG xml for given code
func GetBarcodeSVG(code, variant string, w, h int, color string, showCode bool, inline bool) (string, error) {
	var svg string
	barcodeArray, err := barcodeC128(code, variant)
	if err != nil {
		return "", err
	}

	if !inline {
		svg = "<?xml version=\"1.0\" standalone=\"no\" ?>\n"
		svg = svg + "<!DOCTYPE svg PUBLIC \"-//W3C//DTD SVG 1.1//EN\" \"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd\">\n"
	}

	svg = svg + "<svg width=\"" + strconv.Itoa(barcodeArray.MaxW*w) + "\" height=\"" + strconv.Itoa(h) + "\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\">\n"
	x := 0
	bw := 0
	bh := 0
	for _, value := range barcodeArray.BCode {
		bw = value.W * w
		bh = value.H * h / barcodeArray.MaxH

		if showCode {
			bh = bh - 12
		}

		if value.T {
			y := value.P * h / barcodeArray.MaxH
			svg = svg + "\t<rect x=\"" + strconv.Itoa(x) + "\" y=\"" + strconv.Itoa(y) + "\" width=\"" + strconv.Itoa(bw) + "\" height=\"" + strconv.Itoa(bh) + "\" fill=\"" + color + "\" stroke=\"none\" />\n"
		}
		x = (x + bw)
	}
	if showCode {
		xCode := (barcodeArray.MaxW * w) / 2
		codeX := strconv.FormatInt(int64(xCode), 10)
		svg = svg + "\t<text x=\"" + codeX + "\" text-anchor=\"middle\" y=\"" + strconv.FormatInt(int64((bh+12)), 10) + "\" id=\"code\" fill=\"" + color + "\" font-size=\"12px\">" + barcodeArray.Code + "</text>\n"
	}
	svg = svg + "\n</svg>\n"
	return svg, nil
}
