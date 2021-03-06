package invoices

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	util "github.com/shyang107/gout"
)

const (
	idxDetHead = iota
	idxDetUINumber
	idxDetSubtotal
	idxDetName
)

var (
	// detailFieldNames []string
	detailFieldNames = []string{"Head", "UINumber", "Subtotal", "Name"}
	// detailCtagNames  []string
	detailCtagNames = []string{"表頭", "發票號碼", "小計", "品項名稱"}
	detailIndex     = make(map[string]int)
)

func init() {
	// var err error
	// if detailFieldNames, err = util.Names(&Detail{}, "Model"); err != nil {
	// 	util.Panic("retrive field names failed (%q)!", "detailFieldNames")
	// }
	// if detailCtagNames, err = util.NamesFromTag(&Detail{}, "cht", "Model"); err != nil {
	// 	util.Panic("retrive field-tag names failed [(%q).Tag(%q)]!", "detailCtagNames", "cht")
	// }
	for i := 0; i < len(detailFieldNames); i++ {
		detailIndex[detailFieldNames[i]] = i
	}
}

// Detail : 消費發票明細
// 明細=D	發票號碼 小計 品項名稱
// 範例：
// D ZZ00000050 42.00 拿鐵熱咖啡(中)
// D ZZ00000050 55.00 拿鐵冰咖啡(大)
type Detail struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	// gorm.Model
	// Or alternatively write:
	Model gorm.Model `json:"-" yaml:"-" gorm:"embedded"`
	// ID       int     `json:"-" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Head     string  `cht:"表頭" json:"HEAD" yaml:"HEAD" sql:"DEFAULT:'D'"`
	UINumber string  `cht:"發票號碼" json:"UNIFORM_INVOICE_NUMBER" yaml:"UNIFORM_INVOICE_NUMBER" sql:"size:10;index" gorm:"column:uin"`
	Subtotal float64 `cht:"小計" json:"SUBTOTAL_AMOUNT" yaml:"SUBTOTAL_AMOUNT"`
	Name     string  `cht:"品項名稱" json:"NAME" yaml:"NAME"`
	// Invoice  *Invoice   `json:"-"`
}

func (d Detail) String() string {
	Sf, Ff := fmt.Sprintf, fmt.Fprintf
	var b bytes.Buffer
	val := reflect.ValueOf(d) //.Elem()
	fld := val.Type()
	var str string
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i).Interface()
		f := fld.Field(i)
		tag := f.Tag

		switch v.(type) {
		case gorm.Model:
			continue
		case float64:
			str = Sf("%.1f", v.(float64))
		default:
			switch f.Name {
			case detailFieldNames[idxDetUINumber]: // "UINumber":
				str = v.(string)[0:2] + "-" + v.(string)[2:]
			default:
				str = v.(string)
			}
		}
		Ff(&b, " %s : %s |", tag.Get("cht"), str)
	}
	Ff(&b, "\n")
	return b.String()
}

// TableName : set Detail's table name to be `details`
func (Detail) TableName() string {
	// custom table name, this is default
	return "details"
}

// Table :
func (d *Detail) Table(title string, lensp int) string {
	if len(title) == 0 {
		title = "明細清單"
	}
	// dheads := []string{"表頭", "發票號碼", "小計", "品項名稱"}
	if lensp < 0 {
		lensp = 0
	}
	// table := util.ArgsTableN(title, lensp, false, detailCtagNames, d.Head,
	// d.UINumber[0:2]+"-", d.UINumber[2:], Sf("%.1f", d.Subtotal), d.Name)
	slice := d.interfaceSlice(-1)
	table := util.ArgsTableN(title, lensp, false, detailCtagNames, slice...)
	return table
}

var dcb util.ValuesCallback = func(f reflect.StructField,
	v interface{}) (value interface{}, isIgnored bool) {
	switch v.(type) {
	case gorm.Model:
		value, isIgnored = nil, true
	case float64:
		a := v.(float64)
		value, isIgnored = interface{}(fmt.Sprintf("%.1f", a)), false
	default:
		switch f.Name {
		case "UINumber":
			a := v.(string)
			value, isIgnored = interface{}(a[0:2]+"-"+a[2:]), false
		default:
			value, isIgnored = v, false
		}
	}
	return value, isIgnored
}

func (d *Detail) interfaceSlice(idx int) []interface{} {
	out, err := util.ValuesWithFunc(d, dcb, "Model")
	if err != nil {
		util.Panic("retrive value of `*v` struct failed!")
	}

	if idx < 0 {
		return out
	}
	res := []interface{}{fmt.Sprintf("%d", idx)}
	res = append(res, out...)

	return res
}

func (d *Detail) stringSlice(idx int) []string {
	out, err := util.StrValuesWithFunc(d, dcb)
	if err != nil {
		util.Panic("retrive value of `*v` struct failed!")
	}

	if idx < 0 {
		return out
	}
	res := []string{fmt.Sprintf("%d", idx)}
	res = append(res, out...)

	return res
}

func (d *Detail) rowString(leading string, idx int, sizes []int, isleft bool) string {
	data := d.stringSlice(idx)
	return getDeailTableRowString(&data, leading, idx, sizes, isleft)
}

func getDeailTableRowString(data *[]string,
	leading string, idx int, sizes []int, isleft bool) string {
	// SubTotal
	l := len(*data)
	(*data)[l-2] = util.AlignToRight((*data)[l-2], sizes[l-2])

	return sliceToString(leading, data, sizes, isleft)
}

//=========================================================

// DetailCollection is the collection of "*Detail"
type DetailCollection []*Detail

func (d DetailCollection) String() string {
	var lines string
	for i, p := range d {
		lines += fmt.Sprintf("#%d: %s", i, p.String())
	}
	return lines
}

// Table returns the table string of the list of []*Detail
func (d *DetailCollection) Table() string {
	pds := ([]*Detail)(*d)
	lensp := 6
	return GetDetailsTable(pds, lensp, "", true)
}

// GetDetailsTable returns the table string of the list of []*Detail
func GetDetailsTable(pds []*Detail, lensp int, title string, isTitle bool) string {
	if isTitle {
		if len(title) == 0 {
			title = "明細清單"
		}
	} else {
		title = ""
	}

	dheads := []string{"項次"} //, "表頭", "發票號碼", "小計", "品項名稱"}
	dheads = append(dheads, detailCtagNames...)
	if lensp < 0 {
		lensp = 0
	}
	var data []interface{}
	for i, p := range pds {
		data = append(data, p.interfaceSlice(i+1)...)
	}
	table := util.ArgsTableN(title, lensp, true, dheads, data...)
	return table
}

// Add adds `p` into `v`
func (d *DetailCollection) Add(p *Detail) {
	*d = append(*d, p)
}
