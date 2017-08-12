package invoices

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shyang107/go-twinvoices/util"
)

var (
	// invoicesCache is used to reduce the count of created Invoice objects and
	// allows to reuse already created objects with required Attribute.
	invoicesCache   = make(map[string]*Invoice)
	invoicesCacheMu sync.Mutex // protects invoicesCache
)

// Invoice : 消費發票
// 表頭=M	發票狀態 發票號碼 發票日期 商店統編 商店店 載具名稱 載具號碼 總金額
// 範例：
// M 開立、作廢 ZZ00000050 20130111 97162640 新北市第1000號門市 手機條碼	/WYY+.,HG 97
type Invoice struct {
	// auto-populate columns: id, created_at, updated_at, deleted_at
	// gorm.Model
	// Or alternatively write:
	Model gorm.Model `json:"-" yaml:"-" gorm:"embedded"`
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
	Details []*Detail `cht:"明細清單" json:"DETAILS" yaml:"DETAILS" gorm:"ForeignKey:UINumber;AssociationForeignKey:UINumber"`
}

var invoiceFieldNames []string
var invoiceCtagNames []string
var invoiceIndex = make(map[string]int)

func init() {
	var err error
	invoiceFieldNames, _, _, invoiceCtagNames, err = util.GetFieldsInfo(&Invoice{}, "cht", "Model", "Details")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(invoiceFieldNames); i++ {
		invoiceIndex[invoiceFieldNames[i]] = i
	}

}

func (v Invoice) String() string {
	Sf, Ff := fmt.Sprintf, fmt.Fprintf
	var b bytes.Buffer
	val := reflect.ValueOf(v) //.Elem()
	fld := val.Type()
	var str string
	for i := 0; i < val.NumField(); i++ {
		f := fld.Field(i)
		v := val.Field(i)

		switch v.Interface().(type) {
		case gorm.Model, []*Detail:
			continue
		case time.Time:
			str = v.Interface().(time.Time).Format(ShortDateFormat)
		case float64:
			str = Sf("%.1f", v.Interface().(float64))
		default:
			switch f.Name {
			case invoiceFieldNames[invoiceIndex["UINumber"]]:
				str = v.Interface().(string)[0:2] + "-" + v.Interface().(string)[2:]
			default:
				str = v.Interface().(string)
			}
		}
		Ff(&b, " %s : %s |", invoiceCtagNames[invoiceIndex[f.Name]], str)
	}
	Ff(&b, "\n")
	lspaces := util.StrSpaces(4)
	for i, d := range v.Details {
		Ff(&b, "%s> %2d. %s", lspaces, i+1, d)
	}
	return b.String()
	// re, _ := regexp.Compile("^[\u4e00-\u9fa5]")
}

// TableName : set Invoice's table name to be `invoices`
func (Invoice) TableName() string {
	// custom table name, this is default
	return "invoices"
}

func getCachedInvoices(obj *Invoice) *Invoice {
	invoicesCacheMu.Lock()
	defer invoicesCacheMu.Unlock()
	invoice, ok := invoicesCache[obj.UINumber]
	if !ok {
		invoicesCache[obj.UINumber] = obj
	}
	return invoice
}

func setCachedInvoices(obj *Invoice) {
	invoicesCacheMu.Lock()
	defer invoicesCacheMu.Unlock()
	_, ok := invoicesCache[obj.UINumber]
	if !ok {
		invoicesCache[obj.UINumber] = obj
	}
	return
}

// GetArgsTable :
func (v *Invoice) GetArgsTable(title string) string {
	Sf := fmt.Sprintf
	if len(title) == 0 {
		title = "發票清單"
	}
	// heads := []string{"表頭", "發票狀態", "發票號碼", "發票日期",
	// "商店統編", "商店店名", "載具名稱", "載具號碼", "總金額", "明細清單"}
	lensp := 0
	table := util.ArgsTableN(title, lensp, false, invoiceCtagNames, v.Head, v.State,
		v.UINumber[0:2]+"-"+v.UINumber[2:], v.Date.Format(ShortDateFormat),
		v.SUN, v.SName, v.CName, v.CNumber,
		Sf("%.1f", v.Total), "[如下...]")
	lensp = 7
	table += GetDetailsTable(v.Details, lensp, false)
	return table
}

func (v *Invoice) mapToStringSlice(idx int) []string {
	return []string{
		fmt.Sprintf("%d", idx), v.Head, v.State, v.UINumber[0:2] + "-" + v.UINumber[2:],
		v.Date.Format(ShortDateFormat),
		v.SUN, v.SName, v.CName, v.CNumber, fmt.Sprintf("%.1f", v.Total),
	}
}

func (v *Invoice) toTableRowString(leading string, idx int, sizes []int, isleft bool) string {
	data := v.mapToStringSlice(idx)
	return sliceToString(leading, data, sizes, isleft)
}

func setMaxSizes(sizes *[]int, data *[]string) {
	for j, d := range *data {
		str := fmt.Sprintf("%v", d)
		_, _, nmix := util.CountChars(str)
		(*sizes)[j] = util.Imax((*sizes)[j], nmix)
	}
}

func sliceToString(leading string, data []string, sizes []int, isleft bool) string {
	str := ""
	for i, d := range data {
		sdf := util.GetColStr(d, sizes[i], isleft)
		if i == 0 {
			str += leading + fmt.Sprintf("%v", sdf)
			continue
		}
		str += fmt.Sprintf("  %v", sdf)
	}
	return str
}

// GetInvoicesTable returns the table string of the list of []*Invoice
func GetInvoicesTable(pinvs []*Invoice) string {
	// vheads := []string{"項次", "表頭", "發票狀態", "發票號碼", "發票日期",
	// 	"商店統編", "商店店名", "載具名稱", "載具號碼", "總金額"}
	// dheads := []string{"項次", "表頭", "發票號碼", "小計", "品項名稱"}

	vheads := append([]string{"項次"}, invoiceCtagNames...)
	dheads := append([]string{"項次"}, detailCtagNames...)
	vsizes, dsizes := util.NewSize(vheads), util.NewSize(dheads)

	for i, p := range pinvs {
		vdata := p.mapToStringSlice(i + 1)
		setMaxSizes(&vsizes, &vdata)
		for j, d := range p.Details {
			ddata := d.mapToStringSlice(j + 1)
			setMaxSizes(&dsizes, &ddata)
		}
	}

	var b bytes.Buffer
	bws := b.WriteString

	vnf := len(vheads)
	vn := util.Isum(vsizes...) + vnf + (vnf-1)*2 + 1
	title := "發票清單"
	_, _, vl := util.CountChars(title)
	vm := (vn - vl) / 2
	bws(util.StrSpaces(vm) + title + "\n")

	isleft := true
	vhtab := util.StrThickLine(vn)
	vhtab += sliceToString("", vheads, vsizes, isleft)
	vhtab += "\n" + util.StrThinLine(vn)

	lspaces := util.StrSpaces(7)
	dnf := len(dheads)
	dn := util.Isum(dsizes...) + dnf + (dnf-1)*2 + 1
	dhtab := lspaces + util.StrThickLine(dn)
	dhtab += sliceToString(lspaces, dheads, dsizes, isleft)
	dhtab += "\n" + lspaces + util.StrThinLine(dn)

	for i, p := range pinvs {
		bws(vhtab)
		bws(p.toTableRowString("", i+1, vsizes, isleft))
		bws("\n")

		// bws(GetDetailsTable(p.Details, 7, false))
		if len(p.Details) > 0 {
			bws(dhtab)
			for j, d := range p.Details {
				bws(d.toTableRowString(lspaces, j+1, dsizes, isleft))
				bws("\n")
			}
			bws(lspaces + util.StrThickLine(dn))
		}
	}

	return b.String()
}

// dumpCachedInvoicesTable returns the table string of "invoicesCache"
func dumpCachedInvoicesTable() string {
	invoicesCacheMu.Lock()
	defer invoicesCacheMu.Unlock()
	pinvs := make([]*Invoice, 0)
	for _, v := range invoicesCache {
		pinvs = append(pinvs, v)
	}
	return GetInvoicesTable(pinvs)
}

func printInvList(pvs []*Invoice) {
	var b bytes.Buffer
	fp := fmt.Fprintf
	for ip, pv := range pvs {
		// fp(&b, "%d : %s", ip+1, pv)
		fp(&b, "%s", pv.GetArgsTable(util.Sf("發票 %d", ip+1)))
		// for id, pd := range pv.Details {
		// 	fp(&b, "%s", pd.GetArgsTable(io.Sf("Invoices[%d] -- Details[%d]", ip, id), 7))
		// }
		// fp(&b, "\n")
	}
	//util.Pchk("%s", b.String())
	Glog.Debugf("%s", b.String())
}

//=========================================================
