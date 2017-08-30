package invoices

import (
	"fmt"

	util "github.com/shyang107/gout"
	yaml "gopkg.in/yaml.v2"
)

// YAMLInvoices is used for YAML file type
type YAMLInvoices struct {
	FileType    string             `yaml:"FILE_TYPE"`
	FileVersion int                `yaml:"FILE_VERSION"`
	Invoices    *InvoiceCollection `yaml:"INVOICES"`
}

// YAMLMarshaller collects the mathods marshalling or unmarshalling the csv-type data
type YAMLMarshaller struct{}

// MarshalInvoices marshalls the .yaml-type data of invoices
func (YAMLMarshaller) MarshalInvoices(fn string, invoices *InvoiceCollection) error {
	// Prun("  > Writing data to .jsn or .yaml file %q ...\n", fn)
	util.DebugPrintCaller()
	Glog.Infof("➥  Writing data to .yaml file [%s] ...", util.LogColorString("info", fn))
	y := YAMLInvoices{
		FileType:    fileType,
		FileVersion: fileVersion,
		Invoices:    invoices,
	}
	// b, err := jsoniter.MarshalIndent(&j, "", "    ")
	b, err := yaml.Marshal(&y)
	if err != nil {
		return err
	}
	util.WriteBytesToFile(fn, b)
	return nil
}

// UnmarshalInvoices unmarshalls the .yaml-type data of invoices
func (YAMLMarshaller) UnmarshalInvoices(fn string) (*InvoiceCollection, error) {
	// Prun("  > Reading data from .jsn or .yaml file %q ...\n", fn)
	util.DebugPrintCaller()
	Glog.Infof("➥  Reading data from .yaml file [%s] ...", util.LogColorString("info", fn))
	b, err := util.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	var y YAMLInvoices
	err = yaml.Unmarshal(b, &y)
	if err != nil {
		return nil, err
	}
	if y.FileType != fileType {
		return nil, fmt.Errorf("cannot read non-invoices .yaml file")
	}
	if y.FileVersion > fileVersion {
		return nil, fmt.Errorf("version %d is too new to read", y.FileVersion)
	}
	Glog.Infof("Invoices table:\n%s", y.Invoices.Table())
	y.Invoices.AddToDB()
	return y.Invoices, nil
}
