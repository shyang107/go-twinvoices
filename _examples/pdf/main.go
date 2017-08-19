package main

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/shyang107/go-twinvoices/util"
)

func main() {
	test1()
}

func test1() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		util.Glog.Error(err.Error())
	}

}
