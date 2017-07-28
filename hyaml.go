package invoices

import (
	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	yaml "gopkg.in/yaml.v2"
)

// YAMLInvoices is used for YAML file type
type YAMLInvoices struct {
	FileType   string
	FileVesion int
	Invoices   []*Invoice
}

// YAMLMarshaller collects the mathods marshalling or unmarshalling the csv-type data
type YAMLMarshaller struct{}

// MarshalInvoices marshalls the csv-type data of invoices
func (YAMLMarshaller) MarshalInvoices(fn string, invoices []*Invoice) error {
	prun("  > Writing data to .jsn or .yaml file %q ...\n", fn)
	y := YAMLInvoices{
		FileType:   fileType,
		FileVesion: fileVesion,
		Invoices:   invoices,
	}
	// b, err := jsoniter.MarshalIndent(&j, "", "    ")
	b, err := yaml.Marshal(&y)
	if err != nil {
		return err
	}
	io.WriteBytesToFile(fn, b)
	return nil
}

// UnmarshalInvoices unmarshalls the csv-type data of invoices
func (YAMLMarshaller) UnmarshalInvoices(fn string) ([]*Invoice, error) {
	prun("  > Reading data from .jsn or .yaml file %q ...\n", fn)
	b, err := io.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	var y YAMLInvoices
	err = yaml.Unmarshal(b, &y)
	if err != nil {
		return nil, err
	}
	if y.FileType != fileType {
		return nil, chk.Err("cannot read non-invoices .yaml file")
	}
	if y.FileVesion > fileVesion {
		return nil, chk.Err("version %d is too new to read", y.FileVesion)
	}
	plog(GetInvoicesTable(y.Invoices))
	prun("    updating database ...\n")
	DBInsertFrom(y.Invoices)
	return y.Invoices, nil
}
