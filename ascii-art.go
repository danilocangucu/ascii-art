package asciioutput

import (
	"log"
	"os"
	"strings"
)

func Ascii(term string, style string) []string {
	input := strings.Split(term, "\r\n")
	var file string

	if style == "standard" {
		file = "asciioutput/standard.txt"
	} else if style == "shadow" {
		file = "asciioutput/shadow.txt"
	} else if style == "thinkertoy" {
		file = "asciioutput/thinkertoy.txt"
	}

	asciiCharactersLinesBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	asciiCharactersLines := strings.Split(string(asciiCharactersLinesBytes), "\n")

	var AsciiForTxt []string

	for _, termToBePrinted := range input {
		if termToBePrinted == "" {
			AsciiForTxt = append(AsciiForTxt, "")
			continue
		}
		for i := 0; i < 8; i++ {
			var asciiPrintingLine string
			for _, characterToBePrinted := range termToBePrinted {
				asciiPrintingLine += asciiCharactersLines[1+(int(characterToBePrinted)-' ')*9+i]
			}

			AsciiForTxt = append(AsciiForTxt, asciiPrintingLine)
		}

	}

	return AsciiForTxt
}
