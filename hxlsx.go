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
	Glog.Infof("➥  Writing data to .xlsx file [%s] ...", util.LogColorString("info", fn))
	if pvs == nil || len(pvs) == 0 {
		return fmt.Errorf("pvs []*Invoice = nil or it's len = 0 ")
	}
	var vh = headType{head: []string{"項次"}}
	var dh = headType{head: []string{"項次"}}
	vh.head = append(vh.head, invoiceCtagNames...)
	dh.head = append(dh.head, detailCtagNames...)

	fx := xlsx.NewFile()
	sht, _ := fx.AddSheet("消費發票")

	total := 0.0
	for i := 0; i < len(pvs); i++ {
		vh.addTo(sht.AddRow(), false)
		rowi := sht.AddRow()
		pvs[i].addTo(rowi, i+1)
		total += pvs[i].Total
		if len(pvs[i].Details) > 0 {
			dh.addTo(sht.AddRow(), true)
			for j := 0; j < len(pvs[i].Details); j++ {
				rowd := sht.AddRow()
				pvs[i].Details[j].addTo(rowd, j+1)
			}
		}
	}

	msg := fmt.Sprintf("發票 %d 至 %d 累計金額：", 1, len(pvs))
	sht.AddRow()
	addSum(sht.AddRow(), msg, 3, total)

	fx.Save(fn)
	return nil
}

func addSum(r *xlsx.Row, msg string, mergedcells int, value float64) {
	cell := r.AddCell()
	cell.Merge(mergedcells, 0)
	cell.SetValue(msg)

	s := xlsx.NewStyle()
	s.Alignment.Horizontal = "right"
	s.ApplyAlignment = true
	cell.SetStyle(s)

	for i := 0; i < mergedcells; i++ {
		r.AddCell()
	}
	cell = r.AddCell()
	cell.SetFloatWithFormat(value, numfmt)

	s = xlsx.NewStyle()
	// s.Fill = xlsx.Fill{PatternType: "solid", BgColor: darkBlue, FgColor: ivory}
	s.Fill = *xlsx.NewFill("solid", ivory, "")
	s.Font.Color = blue
	s.Font.Bold = true
	s.ApplyFill = true
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	s.Border = border
	s.ApplyBorder = true
	cell.SetStyle(s)
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
		v := val.Field(i)
		cell := &xlsx.Cell{}
		// cell := r.AddCell()
		cell.SetStyle(style)
		switch v.Interface().(type) {
		case gorm.Model, []*Detail:
			continue
		case time.Time:
			cell.SetDate(v.Interface().(time.Time))
		case float64:
			// cell.SetFloatWithFormat(v.(float64), numfmtAccountant)
			cell.SetFloatWithFormat(v.Interface().(float64), numfmt)
		default:
			cell.SetString(v.Interface().(string))
		}
		r.Cells = append(r.Cells, cell)
	}
	return
}

// UnmarshalInvoices unmarshal the records of invoice using in .xlsx file
func (XlsxMarshaller) UnmarshalInvoices(fn string) ([]*Invoice, error) {
	util.DebugPrintCaller()
	Glog.Infof("➥  Reading data from .xlsx file %s ...", util.LogColorString("info", fn))
	Glog.Warnf("☹  TODO: %q", util.CallerName(1))
	return nil, nil
}
