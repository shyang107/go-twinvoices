package invoices

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shyang107/go-twinvoices/util"
)

var (
	// invoiceFieldNames []string
	invoiceFieldNames = []string{"Head", "State", "UINumber", "Date",
		"SUN", "SName", "CName", "CNumber", "Total"}
	// invoiceCtagNames  []string
	invoiceCtagNames = []string{"表頭", "發票狀態", "發票號碼", "發票日期",
		"商店統編", "商店店名", "載具名稱", "載具號碼", "總金額"}
	invoiceCtagMap = make(map[string]string)
	// invoiceCtagMap = map[string]string{
	// 	"Head": "表頭", "State": "發票狀態", "UINumber": "發票號碼", "Date": "發票日期",
	// 	"SUN": "商店統編", "SName": "商店店名",
	// 	"CName": "載具名稱", "CNumber": "載具號碼", "Total": "總金額",
	// }
	invoiceIndex = make(map[string]int)
	// invoiceIndex = map[string]int{
	// 	"Head": 0, "State": 1, "UINumber": 2, "Date": 3,
	// 	"SUN": 4, "SName": 5, "CName": 6, "CNumber": 7, "Total": 8,
	// }
)

func init() {
	// var err error
	// if invoiceFieldNames, err = util.Names(&Invoice{}, "Model", "Details"); err != nil {
	// 	util.Panic("retrive field names failed (%q)!", "invoiceFieldNames")
	// }
	// if invoiceCtagNames, err = util.NamesFromTag(&Invoice{}, "cht", "Model", "Details"); err != nil {
	// 	util.Panic("retrive field-tag names failed [(%q).Tag(%q)]!", "invoiceCtagNames", "cht")
	// }
	for i := 0; i < len(invoiceFieldNames); i++ {
		invoiceIndex[invoiceFieldNames[i]] = i
		invoiceCtagMap[invoiceFieldNames[i]] = invoiceCtagNames[i]
	}

}

// Invoice : 消費發票
// 表頭=M	發票狀態 發票號碼 發票日期 商店統編 商店店 載具名稱 載具號碼 總金額
// 範例：
// M 開立、作廢 ZZ00000050 20130111 97162640 新北市第1000號門市 手機條碼	/WYY+.,HG 97
type Invoice struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	// gorm.Model
	// Or alternatively write:
	Model gorm.Model `json:"-" yaml:"-" gorm:"embedded" xlsx:"-"`
	// ID    int    `json:"-" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Head  string `cht:"表頭" json:"HEAD" yaml:"HEAD" sql:"DEFAULT:'M'"`
	State string `cht:"發票狀態" json:"STATE" yaml:"STATE"`
	// Uniform-Invoice Number or  tax ID number
	UINumber string    `cht:"發票號碼" json:"UNIFORM_INVOICE_NUMBER" yaml:"UNIFORM_INVOICE_NUMBER" sql:"size:10;unique;index" gorm:"column:uin"`
	Date     time.Time `cht:"發票日期" json:"PURCHASE_DATE" yaml:"PURCHASE_DATE" sql:"index"`
	// Date    string     `cht:"發票日期" json:"date" sql:"index"`
	SUN     string  `cht:"商店統編" json:"STORE_UNIFORM_NUMBER" yaml:"STORE_UNIFORM_NUMBER"`
	SName   string  `cht:"商店店名" json:"STORE_NAME" yaml:"STORE_NAME"`
	CName   string  `cht:"載具名稱" json:"CARRIER_NAME" yaml:"CARRIER_NAME"`
	CNumber string  `cht:"載具號碼" json:"CARRIER_NUMBER" yaml:"CARRIER_NUMBER"`
	Total   float64 `cht:"總金額" json:"TOTAL_AMOUNT" yaml:"TOTAL_AMOUNT"`
	// one-to-many relationship
	Details []*Detail `cht:"明細清單" json:"DETAILS" yaml:"DETAILS" gorm:"ForeignKey:UINumber;AssociationForeignKey:UINumber" xlsx:"-"`
}

func (v Invoice) String() string {
	Sf, Ff := fmt.Sprintf, fmt.Fprintf
	var b bytes.Buffer
	val := reflect.ValueOf(v) //.Elem()
	fld := val.Type()
	var str, line string
	for i := 0; i < val.NumField(); i++ {
		f := fld.Field(i)
		v := val.Field(i).Interface()
		tag := f.Tag

		switch v.(type) {
		case gorm.Model, []*Detail:
			continue
		case time.Time:
			str = v.(time.Time).Format(ShortDateFormat)
		case float64:
			str = Sf("%.1f", v.(float64))
		default:
			switch f.Name {
			case "UINumber":
				str = v.(string)[0:2] + "-" + v.(string)[2:]
			default:
				str = v.(string)
			}
		}
		line += Sf(" %s : %s |", tag.Get("cht"), str)
	}
	line = strings.Trim(line, " ")
	Ff(&b, "%v\n", line)

	lspaces := util.StrSpaces(4)
	for i, d := range v.Details {
		Ff(&b, "%s> %2d. %s", lspaces, i+1, d.String())
	}
	return b.String()
	// re, _ := regexp.Compile("^[\u4e00-\u9fa5]")
}

// TableName : set Invoice's table name to be `invoices`
func (Invoice) TableName() string {
	// custom table name, this is default
	return "invoices"
}

// Table :
func (v *Invoice) Table(title string) string {
	if len(title) == 0 {
		title = "發票清單"
	}
	// heads := []string{"表頭", "發票狀態", "發票號碼", "發票日期",
	// "商店統編", "商店店名", "載具名稱", "載具號碼", "總金額", "明細清單"}
	lensp := 0
	// table := util.ArgsTableN(title, lensp, true, invoiceCtagNames, v.Head, v.State,
	// 	v.UINumber[0:2]+"-"+v.UINumber[2:], v.Date.Format(ShortDateFormat),
	// 	v.SUN, v.SName, v.CName, v.CNumber,
	// 	 fmt.Sprintf("%.1f", v.Total), "[如下...]")
	data := v.interfaceSlice(-1)
	table := util.ArgsTableN(title, lensp, true, invoiceCtagNames, data...)
	lensp = 7
	table += GetDetailsTable(v.Details, lensp, false)
	return table
}

var vcb util.ValuesCallback = func(f reflect.StructField,
	v interface{}) (value interface{}, isIgnored bool) {

	switch v.(type) {
	case gorm.Model, []*Detail:
		value, isIgnored = nil, true
	case time.Time:
		a := v.(time.Time)
		value, isIgnored = interface{}(a.Format(ShortDateFormat)), false
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

func (v *Invoice) interfaceSlice(idx int) []interface{} {
	// if idx < 0 {
	// 	return []interface{}{
	// 		v.Head, v.State, v.UINumber[0:2] + "-" + v.UINumber[2:],
	// 		v.Date.Format(ShortDateFormat),
	// 		v.SUN, v.SName, v.CName, v.CNumber, fmt.Sprintf("%.1f", v.Total),
	// 	}
	// }
	// return []interface{}{
	// 	fmt.Sprint(idx), v.Head, v.State, v.UINumber[0:2] + "-" + v.UINumber[2:],
	// 	v.Date.Format(ShortDateFormat),
	// 	v.SUN, v.SName, v.CName, v.CNumber, fmt.Sprintf("%.1f", v.Total),
	// }

	out, err := util.ValuesWithFunc(v, vcb)
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

func (v *Invoice) stringSlice(idx int) []string {
	// if idx < 0 {
	// 	return []string{
	// 		v.Head, v.State, v.UINumber[0:2] + "-" + v.UINumber[2:],
	// 		v.Date.Format(ShortDateFormat),
	// 		v.SUN, v.SName, v.CName, v.CNumber, fmt.Sprintf("%.1f", v.Total),
	// 	}
	// }
	// return []string{
	// 	fmt.Sprintf("%d", idx), v.Head, v.State, v.UINumber[0:2] + "-" + v.UINumber[2:],
	// 	v.Date.Format(ShortDateFormat),
	// 	v.SUN, v.SName, v.CName, v.CNumber, fmt.Sprintf("%.1f", v.Total),
	// }

	out, err := util.StrValuesWithFunc(v, vcb)
	if err != nil {
		util.Panic("retrive field-values of `*v` failed!")
	}

	if idx < 0 {
		return out
	}
	res := []string{fmt.Sprintf("%d", idx)}
	res = append(res, out...)

	return res
}

func (v *Invoice) rowString(leading string, idx int, sizes []int, isleft bool) string {
	data := v.stringSlice(idx)
	return getInvoiceTableRowString(&data, leading, idx, sizes, isleft)
}

func getInvoiceTableRowString(data *[]string,
	leading string, idx int, sizes []int, isleft bool) string {
	// Total
	l := len(*data)
	(*data)[l-1] = util.AlignToRight((*data)[l-1], sizes[l-1])

	return sliceToString(leading, data, sizes, isleft)
}

func sliceToString(leading string, data *[]string, sizes []int, isleft bool) string {
	str := ""
	for i, d := range *data {
		sdf := util.GetColStr(d, sizes[i], isleft)
		if i == 0 {
			str += leading + sdf // fmt.Sprintf("%v", sdf)
			continue
		}
		str += sdf //fmt.Sprintf(" %v", sdf)
	}
	return str
}

//=========================================================

// InvoiceCollection is the collection of "*Invoice"
type InvoiceCollection []*Invoice

func (v *InvoiceCollection) String() string {
	var b bytes.Buffer
	for i, p := range *v {
		fmt.Fprintf(&b, "Invoice #%d: %v", i, p.String())
	}
	return b.String()
}

// stringSlice returns all field-values to string slice and get max sizes into vsizes and dsizes
// 	vsizes, dsizes []int : original maximum sizes w.r.t. fields
// 	vdata [][]string : vdata[idx of invoices][idx of fields]
// 	rvsizes []int: maiximum sizes w.r.t. fields of invoices
// 	ddata [][][]string: ddata[idx of invoices][idx of details][idx of fields of detail]
// 	rdsizes []int: maiximum sizes w.r.t. fields of details
func (v *InvoiceCollection) stringSlice(vsizes, dsizes []int) (
	vdata [][]string, rvsizes []int,
	ddata [][][]string, rdsizes []int) {
	rvsizes, rdsizes = vsizes, dsizes
	vdata = make([][]string, len(*v))
	ddata = make([][][]string, len(*v))
	for i, p := range *v {
		vdata[i] = p.stringSlice(i + 1)
		for k, f := range vdata[i] {
			_, _, n := util.CountChars(f)
			rvsizes[k] = util.Imax(rvsizes[k], n)
		}

		ddata[i] = make([][]string, len(p.Details))
		for j, d := range p.Details {
			ddata[i][j] = d.stringSlice(j + 1)
			for k, f := range ddata[i][j] {
				_, _, n := util.CountChars(f)
				rdsizes[k] = util.Imax(rdsizes[k], n)
			}
		}
	}
	return vdata, rvsizes, ddata, rdsizes
}

// Table returns the table string of the list of []*Invoice
func (v *InvoiceCollection) Table() string {
	// vheads := []string{"項次", "表頭", "發票狀態", "發票號碼", "發票日期",
	// 	"商店統編", "商店店名", "載具名稱", "載具號碼", "總金額"}
	// dheads := []string{"項次", "表頭", "發票號碼", "小計", "品項名稱"}

	// nv := len(pinvs)
	vheads := append([]string{"項次"}, invoiceCtagNames...)
	dheads := append([]string{"項次"}, detailCtagNames...)
	vsizes, dsizes := util.GetSizes(vheads), util.GetSizes(dheads)
	vnf, dnf := len(vheads), len(dheads)

	vdata, vsizes, ddata, dsizes := v.stringSlice(vsizes, dsizes)

	var b bytes.Buffer
	bws := b.WriteString

	vn := util.Isum(vsizes...) + vnf*2 // + vnf + (vnf - 1) + 1
	title := "發票清單"
	_, _, vl := util.CountChars(title)
	vm := (vn - vl) / 2
	bws(util.StrSpaces(vm) + title + "\n")

	isleft := true
	vheads[vnf-1] = util.AlignToRight(vheads[vnf-1], vsizes[vnf-1]) // Title
	vhtab := util.StrThickLine(vn)
	vhtab += sliceToString("", &vheads, vsizes, isleft)
	vhtab += "\n" + util.StrThinLine(vn)

	lspaces := util.StrSpaces(6)
	dn := util.Isum(dsizes...) + dnf*2                              //+ dnf + (dnf - 1) + 1
	dheads[dnf-2] = util.AlignToRight(dheads[dnf-2], dsizes[dnf-2]) // SubTitle
	dhtab := lspaces + util.StrThickLine(dn)
	dhtab += sliceToString(lspaces, &dheads, dsizes, isleft)
	dhtab += "\n" + lspaces + util.StrThinLine(dn)

	for i := 0; i < len(vdata); i++ {
		bws(vhtab)
		bws(getInvoiceTableRowString(&vdata[i], "", i+1, vsizes, isleft))
		bws("\n")

		bws(dhtab)
		for j := 0; j < len(ddata[i]); j++ {
			bws(getDeailTableRowString(&ddata[i][j], lspaces, j+1, dsizes, isleft))
			bws("\n")
		}
		bws(lspaces + util.StrThickLine(dn))
	}

	return b.String()

}

// GetInvoicesTable returns the table string of the list of []*Invoice
func GetInvoicesTable(pinvs []*Invoice) string {
	var cl InvoiceCollection = pinvs
	return cl.Table()
}

// List returns the brief list string of invoices
func (v *InvoiceCollection) List() string {
	var b bytes.Buffer
	for ip, pv := range *v {
		fmt.Fprintf(&b, "%s", pv.Table(util.Sf("發票 %d", ip+1)))
	}
	return b.String()
}

// GetInvoicesList returns the brief list string of invoices `pvs`
func GetInvoicesList(pvs []*Invoice) string {
	var cl InvoiceCollection = pvs
	return cl.List()
}

// Add adds `p` into `v`
func (v *InvoiceCollection) Add(p *Invoice) {
	*v = append(*v, p)
}

// AddToDB creats records from []*Invoice into database
func (v *InvoiceCollection) AddToDB() {
	dbInsertFrom(([]*Invoice)(*v))
}

// Combine combine DetailCollection into `v`
func (v *InvoiceCollection) Combine(ds *DetailCollection) {
	util.DebugPrintCaller()
	Glog.Infof("♲  combining invoices ...")

	for _, d := range *ds {
		no := d.UINumber
		for _, p := range *v {
			if p.UINumber == no {
				// d.Invoice = p
				p.Details = append(p.Details, d)
				break
			}
		}
	}
}

//=========================================================

// Outformatter is a interface to out-format of object
type Outformatter interface {
	Table() string
	List() string
	String() string
}

func printout(t Outformatter, format string, level string) {
	var p func(...interface{})
	var cs func(...interface{}) string
	var ok bool
	p, ok = util.LogMapFuncs[strings.ToLower(level)]
	if !ok {
		p = util.Glog.Print
	}
	cs, ok = util.LogColorStringFuncs[strings.ToLower(level)]
	if !ok {
		cs = fmt.Sprint
	}

	switch format {
	case "table":
		p(cs("\n", t.Table()))
	case "list":
		p(cs("\n", t.List()))
	default: // "string":
		p(cs("\n", t.String()))
	}
}

// PrintTable print out the table format of `t`
func PrintTable(t Outformatter) {
	printout(t, "table", "print")
}

// PrintList print out the list format of `t`
func PrintList(t Outformatter) {
	printout(t, "list", "print")
}

// PrintString print out the simple format of `t`
func PrintString(t Outformatter) {
	printout(t, "string", "print")
}
