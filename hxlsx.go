package invoices

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shyang107/go-twinvoices/util"
	// "github.com/stanim/xlsxtra"

	"github.com/tealeg/xlsx"
)

const (
	black          = "FF000000"
	white          = "FFFFFFFF"
	red            = "FFFF0000"
	blue           = "FF0000FF"
	yellow         = "FFFFFF00"
	green          = "FF008000"
	pink           = "FFFF00FF"
	turquoise      = "FF00FFFF" // cyan
	darkRed        = "FF800000"
	darkBlue       = "FF000080"
	darkYellow     = "FF808000"
	darkPurple     = "FF660066"
	oceanBlue      = "FF0066CC"
	violet         = "FF800080"
	teal           = "FF008080"
	gray25         = "FFC0C0C0"
	gray40         = "FF969696"
	gray50         = "FF808080"
	gray80         = "FF333333"
	periwinkle     = "FF993366"
	ivory          = "FFFFFFCC"
	coral          = "FFFF8080"
	brightGreen    = "FF00FF00"
	lightGreen     = "FFCCFFCC"
	iceBlue        = "FFCCCCFF"
	lightBlue      = "FF3366FF"
	lightTurquoise = "FFCCFFFF" // light cyan
	lightYellow    = "FFFFFF99"
	//
	numfmtAccountant = `_($* #,##0.0_);_($* (#,##0.0);_($* "-"??_);_(@_)`
	numfmtDollar     = `"NT$"#,##0.0_);[red]"NT$"-#,##0.0`
	numfmt           = `#,##0.0 ;[red]-#,##0.0 `
)

// var (
// 	// letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
// 	// numbers = []rune("0123456789")
// 	// letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 	numbers = []rune("0123456789")
// 	letters = make(map[int]string)
// )

// func init() {
// 	for k, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
// 		letters[k] = string(v)
// 	}
// }

// XlsxMarshaller :
type XlsxMarshaller struct{}

// MarshalInvoices marshal the records of invoice using in .xlsx file
func (XlsxMarshaller) MarshalInvoices(fn string, pvs []*Invoice) error {
	// Prun("  > Writing data to .xlsx file %q ...\n", fn)
	util.DebugPrintCaller()
	glInfof("➥  Writing data to .xlsx file [%s] ...", util.LogColorString("info", fn))
	if pvs == nil || len(pvs) == 0 {
		return fmt.Errorf("pvs []*Invoice = nil or it's len = 0 ")
	}
	var vh, dh headType
	_, _, _, vh.head = util.GetFieldsInfo(Invoice{}, "cht", "Model")
	_, _, _, dh.head = util.GetFieldsInfo(Detail{}, "cht", "Model")
	vh.prepend("項次")
	dh.prepend("項次")

	fx := xlsx.NewFile()
	sht, _ := fx.AddSheet("消費發票")
	for i := 0; i < len(pvs); i++ {
		vh.addTo(sht.AddRow(), false)
		rowi := sht.AddRow()
		pvs[i].addTo(rowi, i+1)
		if len(pvs[i].Details) > 0 {
			dh.addTo(sht.AddRow(), true)
			for j := 0; j < len(pvs[i].Details); j++ {
				rowd := sht.AddRow()
				pvs[i].Details[j].addTo(rowd, j+1)
			}
		}
	}
	fx.Save(fn)
	return nil
}

type headType struct {
	head []string
}

func (ht *headType) prepend(value string) {
	ht.head = append(ht.head, "")
	copy(ht.head[1:], ht.head[0:])
	ht.head[0] = value
}

func (ht *headType) addTo(r *xlsx.Row, isDetail bool) {
	// style := getDefaultInvoiceCellStyle()
	if isDetail {
		r.AddCell()
		// style = getDefaultDetailCellStyle()
	}
	// cell := r.AddCell()
	// cell.SetString("項次")
	// cell.SetStyle(style)
	for i := 0; i < len(ht.head); i++ {
		cell := r.AddCell()
		cell.SetString(ht.head[i])
		// cell.SetStyle(style)
	}
}

func getDefaultDetailCellStyle() *xlsx.Style {
	s := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", lightTurquoise, "")
	s.Fill = fill
	s.ApplyFill = true
	// s.Alignment.ShrinkToFit = true
	// s.ApplyAlignment = true
	border := *xlsx.NewBorder("", "", "thin", "thin")
	s.Border = border
	s.ApplyBorder = true
	return s
}

func (d *Detail) addTo(r *xlsx.Row, id int) {
	style := getDefaultDetailCellStyle()
	//
	r.AddCell()
	cell := r.AddCell()
	cell.SetInt(id)
	// cell.SetStyle(style)

	val := reflect.ValueOf(*d)
	n := val.NumField() // typ.NumField()
	for i := 0; i < n; i++ {
		vi := val.Field(i).Interface()
		cell := &xlsx.Cell{}
		// cell := r.AddCell()
		cell.SetStyle(style)
		switch vi.(type) {
		case gorm.Model:
			continue
		case float64:
			// cell.SetFloat(vi.(float64))
			cell.SetFloatWithFormat(vi.(float64), numfmt)
		default:
			cell.SetString(vi.(string))
		}
		r.Cells = append(r.Cells, cell)
	}
}

func getDefaultInvoiceCellStyle() *xlsx.Style {
	s := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", lightGreen, "")
	s.Fill = fill
	s.ApplyFill = true
	// s.Alignment.ShrinkToFit = true
	// s.ApplyAlignment = true
	border := *xlsx.NewBorder("", "", "thin", "thin")
	s.Border = border
	s.ApplyBorder = true
	return s
}

func (v *Invoice) addTo(r *xlsx.Row, id int) {
	style := getDefaultInvoiceCellStyle()
	//
	cell := r.AddCell()
	cell.SetInt(id)
	// cell.SetStyle(style)

	val := reflect.ValueOf(*v)
	n := val.NumField()
	for i := 0; i < n; i++ {
		vvi := val.Field(i).Interface()
		cell := &xlsx.Cell{}
		// cell := r.AddCell()
		cell.SetStyle(style)
		switch vvi.(type) {
		case gorm.Model, []*Detail:
			continue
		case time.Time:
			cell.SetDate(vvi.(time.Time))
		case float64:
			// cell.SetFloatWithFormat(vvi.(float64), numfmtAccountant)
			cell.SetFloatWithFormat(vvi.(float64), numfmt)
		default:
			cell.SetString(vvi.(string))
		}
		r.Cells = append(r.Cells, cell)
	}
	return
}

// func getFieldsAndTags(obj interface{}, tag string) (fldnames, tagnames []string) {
// 	objval := reflect.ValueOf(obj)
// 	objtyp := objval.Type()
// 	for i := 0; i < objval.NumField(); i++ {
// 		fldval := objval.Field(i)
// 		fldtyp := objtyp.Field(i)
// 		switch fldval.Interface().(type) {
// 		case gorm.Model, []*Detail:
// 			continue
// 		default:
// 			fldnames = append(fldnames, fldtyp.Name)
// 			tagnames = append(tagnames, fldtyp.Tag.Get(tag))
// 		}
// 	}
// 	return fldnames, tagnames
// }

// func getFieldNameAndChtag(obj interface{}) (fldn, cfldn []string) {
// 	vv := reflect.ValueOf(obj)
// 	tv := vv.Type()
// 	for i := 0; i < vv.NumField(); i++ {
// 		field := tv.Field(i)
// 		// typename := field.Type.String()
// 		switch vv.Field(i).Interface().(type) {
// 		case gorm.Model, []*Detail:
// 			continue
// 		default:
// 			fldn = append(fldn, field.Name)
// 			cname := tv.Field(i).Tag.Get("cht")
// 			cfldn = append(cfldn, cname)
// 		}
// 	}
// 	return
// }

// UnmarshalInvoices unmarshal the records of invoice using in .xlsx file
func (XlsxMarshaller) UnmarshalInvoices(fn string) ([]*Invoice, error) {
	util.DebugPrintCaller()
	glInfof("➥  Reading data from .xlsx file %s ...", util.LogColorString("info", fn))
	glWarnf("☹  TODO: %q", util.CallerName(1))
	return nil, nil
}
