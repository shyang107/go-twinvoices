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
var invoiceIndeces = make(map[string]int)

func init() {
	invoiceFieldNames, _, _, invoiceCtagNames = util.GetFieldsInfo(Invoice{}, "cht", "Model")
	for i := 0; i < len(invoiceFieldNames); i++ {
		invoiceIndeces[invoiceFieldNames[i]] = i
	}
}

func (pv Invoice) String() string {
	Sf, Ff := fmt.Sprintf, fmt.Fprintf
	var b bytes.Buffer
	val := reflect.ValueOf(pv) //.Elem()
	fld := val.Type()
	var str string
	for i := 0; i < val.NumField(); i++ {
		f := fld.Field(i)
		v := val.Field(i)
		switch f.Name {
		case invoiceFieldNames[invoiceIndeces["Model"]],
			invoiceFieldNames[invoiceIndeces["Details"]]:
			continue
		case invoiceFieldNames[invoiceIndeces["Date"]]:
			str = v.Interface().(time.Time).Format(ShortDateFormat)
		case invoiceFieldNames[invoiceIndeces["Total"]]:
			str = Sf("%.1f", v.Interface().(float64))
		case invoiceFieldNames[invoiceIndeces["UINumber"]]:
			str = v.Interface().(string)[0:2] + "-" + v.Interface().(string)[2:]
		default:
			str = v.Interface().(string)
		}
		Ff(&b, " %s : %s |", invoiceCtagNames[invoiceIndeces[f.Name]], str)
	}
	Ff(&b, "\n")
	lspaces := util.StrSpaces(4)
	for i, d := range pv.Details {
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
func (pv *Invoice) GetArgsTable(title string) string {
	Sf := fmt.Sprintf
	if len(title) == 0 {
		title = "發票清單"
	}
	// heads := []string{"表頭", "發票狀態", "發票號碼", "發票日期",
	// "商店統編", "商店店名", "載具名稱", "載具號碼", "總金額", "明細清單"}
	lensp := 0
	table := util.ArgsTableN(title, lensp, false, invoiceCtagNames, pv.Head, pv.State,
		pv.UINumber[0:2]+"-"+pv.UINumber[2:], pv.Date.Format(ShortDateFormat),
		pv.SUN, pv.SName, pv.CName, pv.CNumber,
		Sf("%.1f", pv.Total), "[如下...]")
	lensp = 7
	table += GetDetailsTable(pv.Details, lensp)
	return table
}

type invoiceSlcie struct {
	data    []string
	details []detailSlcie
}
type detailSlcie struct {
	data []string
}

// GetInvoicesTable returns the table string of the list of []*Invoice
func GetInvoicesTable(pinvs []*Invoice) string {
	Sf, StrSpaces, StrThickLine, StrThinLine := util.Sf, util.StrSpaces, util.StrThickLine, util.StrThinLine
	// vheads := []string{"項次", "表頭", "發票狀態", "發票號碼", "發票日期",
	// 	"商店統編", "商店店名", "載具名稱", "載具號碼", "總金額"}
	// dheads := []string{"項次", "表頭", "發票號碼", "小計", "品項名稱"}
	vnf := len(invoiceCtagNames)
	dnf := len(detailCtagNames)
	vsizes := make([]int, vnf)
	dsizes := make([]int, vnf)
	for i := 0; i < vnf; i++ {
		_, _, vsizes[i] = util.CountChars(invoiceCtagNames[i])
	}
	for i := 0; i < dnf; i++ {
		_, _, dsizes[i] = util.CountChars(invoiceCtagNames[i])
	}
	//
	invs := make([]invoiceSlcie, len(pinvs))
	for i := 0; i < len(pinvs); i++ {
		p := pinvs[i]
		invs[i].data = []string{
			Sf("%d", i+1), p.Head, p.State, p.UINumber[0:2] + "-" + p.UINumber[2:],
			p.Date.Format(ShortDateFormat),
			p.SUN, p.SName, p.CName, p.CNumber, Sf("%.1f", p.Total),
		}
		for j := 0; j < vnf; j++ {
			str := Sf("%v", invs[i].data[j])
			_, _, nmix := util.CountChars(str)
			vsizes[j] = util.Imax(vsizes[j], nmix)
		}
		for j := 0; j < len(p.Details); j++ {
			d := p.Details[j]
			detail := detailSlcie{
				data: []string{
					Sf("%d", j+1), d.Head, d.UINumber[0:2] + "-" + d.UINumber[2:],
					Sf("%.1f", d.Subtotal), d.Name,
				},
			}
			invs[i].details = append(invs[i].details, detail)
			for k := 0; k < dnf; k++ {
				str := Sf("%v", detail.data[k])
				_, _, nmix := util.CountChars(str)
				dsizes[k] = util.Imax(dsizes[k], nmix)
			}
		}
	}
	vn := util.Isum(vsizes...) + vnf + (vnf-1)*2 + 1
	title := "發票清單"
	_, _, vl := util.CountChars(title)
	vm := (vn - vl) / 2
	isleft := true
	//
	var b bytes.Buffer
	bws := b.WriteString
	//
	bws(StrSpaces(vm) + title + "\n")
	//
	vhtab := StrThickLine(vn)
	svfields := make([]string, vnf)
	for i := 0; i < vnf; i++ {
		svfields[i] = util.GetColStr(invoiceCtagNames[i], vsizes[i], isleft)
		switch i {
		case 0:
			vhtab += Sf("%v", svfields[i])
		default:
			vhtab += Sf("  %v", svfields[i])
		}
	}
	vhtab += "\n" + StrThinLine(vn)
	lspaces := util.StrSpaces(7)
	dn := util.Isum(dsizes...) + dnf + (dnf-1)*2 + 1
	dhtab := lspaces + StrThickLine(dn)
	sdfields := make([]string, dnf)
	for i := 0; i < dnf; i++ {
		sdfields[i] = util.GetColStr(detailCtagNames[i], dsizes[i], isleft)
		switch i {
		case 0:
			dhtab += lspaces + Sf("%v", sdfields[i])
		default:
			dhtab += Sf("  %v", sdfields[i])
		}
	}
	dhtab += "\n" + lspaces + StrThinLine(dn)
	//
	for i := 0; i < len(invs); i++ {
		v := &invs[i].data
		bws(vhtab)
		// pchk("%v : %v \n", vnf, v)
		for j := 0; j < vnf; j++ {
			svfields[j] = util.GetColStr((*v)[j], vsizes[j], isleft)
			switch j {
			case 0:
				bws(Sf("%v", svfields[j]))
			default:
				bws(Sf("  %v", svfields[j]))
			}
		}
		bws("\n")
		//
		details := invs[i].details
		ndetails := len(details)
		if ndetails > 0 {
			bws(dhtab)
			for k := 0; k < len(details); k++ {
				d := &details[k].data
				for j := 0; j < dnf; j++ {
					sdfields[j] = util.GetColStr((*d)[j], dsizes[j], isleft)
					switch j {
					case 0:
						bws(lspaces + Sf("%v", sdfields[j]))
					default:
						bws(Sf("  %v", sdfields[j]))
					}
				}
				bws("\n")
			}
			bws(lspaces + StrThickLine(dn))
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
	glDebugf("%s", b.String())
}

//=========================================================
