package invoices

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/shyang107/go-twinvoices/util"
)

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

var detailFieldNames []string
var detailCtagNames []string
var detailIndeces = make(map[string]int)

func init() {
	detailFieldNames, _, _, detailCtagNames = util.GetFieldsInfo(Detail{}, "cht", "Model")
	for i := 0; i < len(detailFieldNames); i++ {
		detailIndeces[detailFieldNames[i]] = i
	}
}

func (d Detail) String() string {
	Sf, Ff := fmt.Sprintf, fmt.Fprintf
	var b bytes.Buffer
	val := reflect.ValueOf(d) //.Elem()
	fld := val.Type()
	var str string
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		f := fld.Field(i)
		switch f.Name {
		case detailFieldNames[detailIndeces["Model"]]:
			continue
		case detailFieldNames[detailIndeces["Subtotal"]]:
			str = Sf("%.1f", v.Interface().(float64))
		case detailFieldNames[detailIndeces["UINumber"]]:
			str = v.Interface().(string)[0:2] + "-" + v.Interface().(string)[2:]
		default:
			str = v.Interface().(string)
		}
		Ff(&b, " %s : %s |", detailCtagNames[detailIndeces[f.Name]], str)
	}
	Ff(&b, "\n")
	return b.String()
}

// GetArgsTable :
func (d *Detail) GetArgsTable(title string, lensp int) string {
	Sf := fmt.Sprintf
	if len(title) == 0 {
		title = "明細清單"
	}
	// dheads := []string{"表頭", "發票號碼", "小計", "品項名稱"}
	if lensp < 0 {
		lensp = 0
	}
	table := util.ArgsTableN(title, lensp, false, detailCtagNames, d.Head,
		d.UINumber[0:2]+"-", d.UINumber[2:], Sf("%.1f", d.Subtotal), d.Name)
	return table
}

// TableName : set Detail's table name to be `details`
func (Detail) TableName() string {
	// custom table name, this is default
	return "details"
}

func getCachedInvoicesFrom(obj *Detail) (*Invoice, error) {
	invoicesCacheMu.Lock()
	defer invoicesCacheMu.Unlock()
	invoice, ok := invoicesCache[obj.UINumber]
	if !ok {
		return nil, fmt.Errorf("the corresponding invoice does not exist (UINumber: %s)", obj.UINumber)
	}
	return invoice, nil
}

func setCachedInvoicesFrom(obj *Detail) error {
	invoicesCacheMu.Lock()
	defer invoicesCacheMu.Unlock()
	invoice, ok := invoicesCache[obj.UINumber]
	if !ok {
		return fmt.Errorf("the corresponding invoice does not exist (UINumber: %s)", obj.UINumber)
	}
	invoice.Details = append(invoice.Details, obj)
	return nil
}

// GetDetailsTable returns the table string of the list of []*Detail
func GetDetailsTable(pds []*Detail, lensp int) string {
	Sf := fmt.Sprintf
	title := "明細清單"
	dheads := []string{"項次"} //, "表頭", "發票號碼", "小計", "品項名稱"}
	dheads = append(dheads, detailCtagNames...)
	if lensp < 0 {
		lensp = 0
	}
	var data []interface{}
	for i, d := range pds {
		data = append(data, i+1, d.Head,
			d.UINumber[0:2]+"-"+d.UINumber[2:], Sf("%.1f", d.Subtotal), d.Name)
	}
	table := util.ArgsTableN(title, lensp, false, dheads, data...)
	return table
}
