package invoices

import (
	"fmt"

	"github.com/json-iterator/go"
	"github.com/shyang107/go-twinvoices/util"
)

// JSONInvoices is used for JSON file type
type JSONInvoices struct {
	FileType    string             `json:"FILE_TYPE"`
	FileVersion int                `json:"FILE_VERSION"`
	Invoices    *InvoiceCollection `json:"INVOICES"`
}

// JSONMarshaller collects the mathods marshalling or unmarshalling the csv-type data
type JSONMarshaller struct{}

// MarshalInvoices marshalls the .json-type data of invoices
func (JSONMarshaller) MarshalInvoices(fn string, invoices *InvoiceCollection) error {
	// Prun("  > Writing data to .jsn or .json file %q ...\n", fn)
	util.DebugPrintCaller()
	Glog.Infof("➥  Writing data to .json file [%s] ...", util.LogColorString("info", fn))
	j := JSONInvoices{
		FileType:    fileType,
		FileVersion: fileVersion,
		Invoices:    invoices,
	}
	// b, err := jsoniter.MarshalIndent(&j, "", "    ")
	b, err := jsoniter.Marshal(&j)
	if err != nil {
		return err
	}
	util.WriteBytesToFile(fn, b)
	return nil
}

// UnmarshalInvoices unmarshalls the .json-type data of invoices
func (JSONMarshaller) UnmarshalInvoices(fn string) (*InvoiceCollection, error) {
	// Prun("  > Reading data from .jsn or .json file %q ...\n", fn)
	util.DebugPrintCaller()
	Glog.Infof("➥  Reading data from .json file [%s] ...", util.LogColorString("info", fn))
	b, err := util.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	var j JSONInvoices
	err = jsoniter.Unmarshal(b, &j)
	if err != nil {
		return nil, err
	}
	if j.FileType != fileType {
		return nil, fmt.Errorf("☠  cannot read non-invoices json file")
	}
	if j.FileVersion > fileVersion {
		return nil, fmt.Errorf("☠  version %d is too new to read", j.FileVersion)
	}
	// Plog(GetInvoicesTable(j.Invoices))
	// Prun("    updating database ...\n")
	Glog.Infof("Invoices list ---\n%s", j.Invoices.GetInvoicesTable())
	j.Invoices.AddToDB()
	return j.Invoices, nil
}
