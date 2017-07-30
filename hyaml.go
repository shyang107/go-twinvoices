package invoices

import (
	"fmt"

	"github.com/shyang107/go-twinvoices/util"
	yaml "gopkg.in/yaml.v2"
)

// YAMLInvoices is used for YAML file type
type YAMLInvoices struct {
	FileType    string     `yaml:"FILE_TYPE"`
	FileVersion int        `yaml:"FILE_VERSION"`
	Invoices    []*Invoice `yaml:"INVOICES"`
}

// YAMLMarshaller collects the mathods marshalling or unmarshalling the csv-type data
type YAMLMarshaller struct{}

// MarshalInvoices marshalls the .yaml-type data of invoices
func (YAMLMarshaller) MarshalInvoices(fn string, invoices []*Invoice) error {
	// Prun("  > Writing data to .jsn or .yaml file %q ...\n", fn)
	util.DebugPrintCaller()
	glInfof("➥  Writing data to .yaml file %q ...", fn)
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
func (YAMLMarshaller) UnmarshalInvoices(fn string) ([]*Invoice, error) {
	// Prun("  > Reading data from .jsn or .yaml file %q ...\n", fn)
	util.DebugPrintCaller()
	glInfof("➥  Reading data from .yaml file %q ...", fn)
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
	// Plog(GetInvoicesTable(y.Invoices))
	// Prun("    updating database ...\n")
	glInfof("Invoices list ---\n%s", GetInvoicesTable(y.Invoices))
	DBInsertFrom(y.Invoices)
	return y.Invoices, nil
}
