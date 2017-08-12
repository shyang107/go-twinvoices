package invoices

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	iconv "github.com/djimenez/iconv-go"
	"github.com/shyang107/go-twinvoices/util"
)

const (
	csvSep = "|"
)

// CsvMarshaller collects the mathods marshalling or unmarshalling the .csv data
type CsvMarshaller struct{}

// MarshalInvoices marshalls the .csv data of invoices
func (CsvMarshaller) MarshalInvoices(fn string, pvs []*Invoice) error {
	// Prun("  > Writing data to .csv file %q ...\n", fn)
	util.DebugPrintCaller()
	glInfof("➥  Writing data to .csv file [%s] ...", util.LogColorString("info", fn))
	var b bytes.Buffer
	fmt.Fprintln(&b, fileType)
	fmt.Fprintln(&b, io.Sf("%v", fileVersion))
	for _, pv := range pvs {
		fmt.Fprintln(&b, pv.toCSVString())
		for _, d := range pv.Details {
			fmt.Fprintln(&b, d.toCSVString())
		}
	}
	io.WriteFile(fn, &b)
	return nil
}

func (v *Invoice) toCSVString() string {
	csv := []string{
		v.Head,
		v.State,
		v.UINumber,
		v.Date.Format(dateFormat),
		v.SUN,
		v.SName,
		v.CName,
		v.CNumber,
		fmt.Sprintf("%v", v.Total),
	}
	return strings.Join(csv, csvSep)
}

func (d *Detail) toCSVString() string {
	csv := []string{
		d.Head,
		d.UINumber,
		fmt.Sprintf("%v", d.Subtotal),
		d.Name,
	}
	return strings.Join(csv, csvSep)
}

// UnmarshalInvoices unmarshalls the .csv data of invoices
func (CsvMarshaller) UnmarshalInvoices(fn string) ([]*Invoice, error) {
	// Pstat("  > Reading data from .csv file %q ...\n", fn)
	util.DebugPrintCaller()
	glInfof("➥  Reading data from .csv file [%s] ...", util.LogColorString("info", fn))
	f, err := io.OpenFileR(fn)
	if err != nil {
		return nil, err
	}
	var pinvs []*Invoice
	var pdets []*Detail
	err = io.ReadLinesFile(f, func(idx int, line string) (stop bool) {
		// plog("line = %v\n", line)
		if inpIsBig5 {
			line = big5ToUtf8(line)
		} else {
			switch idx {
			case 0:
				ft := strings.Trim(line, " ")
				if ft != fileType {
					util.Panic("☠  type of .csv file is not matched (%q)", fileType)
				}
			case 1:
				fv := io.Atoi(strings.Trim(line, " "))
				if fv != fileVersion {
					util.Panic("☠  version (%v) of .csv file is not matched (%v)", fv, fileVersion)
				}
			}
		}
		recs := strings.Split(line, csvSep)
		head := recs[0]
		switch head {
		case "M": // invoice
			pinv := unmarshalCSVInvoice(recs)
			pinvs = append(pinvs, pinv)
		case "D": // deltail of invoice
			pdet := unmarshalCSVDetail(recs)
			pdets = append(pdets, pdet)
		}
		return
	})
	if err != nil {
		return nil, err
	}
	combineInvoice(pinvs, pdets)
	// fmt.Println("♲  Invoices list:")
	// for i, v := range pinvs {
	// 	fmt.Printf("%d: %v", i, v)
	// }
	glInfof("♲  Invoices list:\n%s", GetInvoicesTable(pinvs))
	dbInsertFrom(pinvs)
	return pinvs, nil
}

func combineInvoice(pvs []*Invoice, pds []*Detail) {
	glInfof("♲  combining invoices ...")
	for _, d := range pds {
		no := d.UINumber
		for _, p := range pvs {
			if p.UINumber == no {
				// d.Invoice = p
				p.Details = append(p.Details, d)
				break
			}
		}
	}
}

func unmarshalCSVDetail(recs []string) *Detail {
	// pchk("%sDetail : %#v\n", io.StrSpaces(4), recs)
	det := Detail{
		Head:     recs[0],
		UINumber: recs[1],
		Subtotal: util.Atof(recs[2]),
		Name:     recs[3],
	}
	return &det
}

func unmarshalCSVInvoice(recs []string) *Invoice {
	// pchk("%sInvoice : %#v\n", io.StrSpaces(4), recs)
	date, err := time.Parse(dateFormat, recs[3])
	location, _ := time.LoadLocation("Local")
	if err != nil {
		panic(chk.Err("%v : %s", err, recs[3]))
	}
	inv := Invoice{
		Head:     recs[0],
		State:    recs[1],
		UINumber: recs[2],
		Date:     date.In(location),
		SUN:      recs[4],
		SName:    recs[5],
		CName:    recs[6],
		CNumber:  recs[7],
		Total:    io.Atof(recs[8]),
	}
	return &inv
}

func big5ToUtf8(str string) string {
	res, err := iconv.ConvertString(str, "big5", "utf-8")
	util.CheckErr(err)
	return res
}
