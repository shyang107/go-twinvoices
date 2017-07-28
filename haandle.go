package invoices

import (
	"os"
	"path/filepath"

	"github.com/cpmech/gosl/chk"
)

var (
	inpIsBig5 = false
)

// InvoiceMarshaller is marshal-operator of invoices
type InvoiceMarshaller interface {
	MarshalInvoices(fn string, pvs []*Invoice) error
}

// InvoiceUnmarshaller is unmarshal-operator of invoices
type InvoiceUnmarshaller interface {
	UnmarshalInvoices(fn string) ([]*Invoice, error)
}

// ReadInvoices reads invoice-record from fn
func (c *InputFile) ReadInvoices() ([]*Invoice, error) {
	var unmarshaller InvoiceUnmarshaller
	startfunc(ffstart) //, "ReadInvoices")
	// pstat("file-type : %q\n", Opts.IfnSuffix)
	//
	var fb = FileBunker{Name: filepath.Base(c.Filename)}
	DB.Where(&fb).First(&fb)
	plog((&fb).GetArgsTable("", 0))
	//
	switch c.Suffix {
	case ".csv":
		pstat("%q\n", "CsvMarshaller")
		unmarshaller = CsvMarshaller{}
	case ".jsn", ".json":
		pstat("%q\n", "JSONMarshaller")
		unmarshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		pstat("%q\n", "YAMLMarshaller")
		unmarshaller = YAMLMarshaller{}
	case ".xml":
		pstat("%q\n", "XMLMarshaller")
		unmarshaller = XMLMarshaller{}
	case ".xlsx":
		pstat("%q\n", "XlsMarshaller")
		unmarshaller = XlsMarshaller{}
	}
	if unmarshaller != nil {
		inpIsBig5 = c.IsBig5
		invs, err := unmarshaller.UnmarshalInvoices(os.ExpandEnv(c.Filename))
		stopfunc(ffstop) //, "ReadInvoices")
		return invs, err
	}
	return nil, chk.Err("not supprted file-type : %s (%s)",
		c.Suffix, c.Filename)
}

// WriteInvoices reads invoice-record from fn
func (o *OutputFile) WriteInvoices(invs []*Invoice) error {
	var marshaller InvoiceMarshaller
	startfunc(ffstart) //, "ReadInvoices")
	switch o.Suffix {
	case ".csv":
		pstat("%q\n", "CsvMarshaller")
		marshaller = CsvMarshaller{}
	case ".jsn", ".json":
		pstat("%q\n", "JSONMarshaller")
		marshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		pstat("%q\n", "YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case ".xml":
		pstat("%q\n", "XMLMarshaller")
		marshaller = XMLMarshaller{}
	case ".xlsx":
		pstat("%q\n", "XlsMarshaller")
		marshaller = XlsMarshaller{}
	}
	if marshaller != nil {
		err := marshaller.MarshalInvoices(os.ExpandEnv(o.Filename), invs)
		stopfunc(ffstop) //, "ReadInvoices")
		return err
	}
	return chk.Err("not supprted file-type : %s (%s)", o.Suffix, o.Filename)
}
