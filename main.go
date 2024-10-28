package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benoitkugler/pdf/formfill"
	"github.com/benoitkugler/pdf/reader"
)

func main() {

	doc, _, err := reader.ParsePDFFile("title-application.pdf", reader.Options{})
	if err != nil {
		log.Fatalf("reading input: %s", err)
	}

	for _, field := range doc.Catalog.AcroForm.Fields {
		fmt.Println(field.FullFieldName())
	}

	var data = []formfill.FDFField{
		{T: "year", Values: formfill.Values{V: formfill.FDFText("2008")}},
		{T: "make", Values: formfill.Values{V: formfill.FDFText("Mazda")}},
		{T: "bodystyle", Values: formfill.Values{V: formfill.FDFText("Sedan")}},
	}

	_ = formfill.FillForm(&doc, formfill.FDFDict{Fields: data}, false)

	out, err := os.Create("filled_title_application.pdf")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if err = doc.Write(out, nil); err != nil {
		fmt.Println("Error:", err)
	}
}
