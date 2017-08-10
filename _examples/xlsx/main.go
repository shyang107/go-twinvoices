package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/jinzhu/gorm"
	inv "github.com/shyang107/go-twinvoices"
	"github.com/xuri/excelize"
)

var (
	// letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// numbers = []rune("0123456789")
	letters = make(map[int]string)
)

func init() {
	for k, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letters[k] = string(v)
	}
}

func main() {
	// test1()
	test2()
}

func test2() {
	_, vheads := getFieldsAndTags(inv.Invoice{}, "cht")
	_, dheads := getFieldsAndTags(inv.Detail{}, "cht")
	fmt.Printf("%v\n", vheads)
	fmt.Printf("%v\n", dheads)

	xl := excelize.NewFile()
	// xl.NewSheet(2, "消費發票")
	for idx, val := range vheads {
		ax := letters[idx] + "1"
		fmt.Print(ax, " ")
		xl.SetCellValue("Sheet1", ax, val)
	}
	fmt.Println()
	for idx, val := range dheads {
		ax := letters[idx+1] + "2"
		fmt.Print(ax, " ")
		xl.SetCellValue("Sheet1", ax, val)
	}
	fmt.Println()
	xl.SetSheetName("Sheet1", "消費發票")
	// xl.SetActiveSheet(1)
	// Save xlsx file by the given path.
	err := xl.SaveAs("./test.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getFieldsAndTags(obj interface{}, tag string) (fldnames, tagnames []string) {
	objval := reflect.ValueOf(obj)
	objtyp := objval.Type()
	for i := 0; i < objval.NumField(); i++ {
		fldval := objval.Field(i)
		fldtyp := objtyp.Field(i)
		switch fldval.Interface().(type) {
		case gorm.Model, []*inv.Detail:
			continue
		default:
			fldnames = append(fldnames, fldtyp.Name)
			tagnames = append(tagnames, fldtyp.Tag.Get(tag))
		}
	}
	return fldnames, tagnames
}
func test1() {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	fmt.Printf("letters = %#v\n", letters)
	fmt.Printf("letters = %q\n", letters)
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("numbers = %q\n", numbers)
	for i := 65; i < 91; i++ {
		fmt.Printf("<%d,%q> ", i, i)
	}
	fmt.Println()
	var lmap = make(map[int]string)
	// var i rune
	for i := 0; i < 26; i++ {
		// lmap[i] = strconv.QuoteRune(rune(i + 65))
		lmap[i] = fmt.Sprintf("%q", i+65)
		// letters[i] = string(i + 65)
	}
	fmt.Printf("%v\n", lmap)
}
