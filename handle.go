package invoices

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shyang107/go-twinvoices/util"
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

func logdebugmarshaller(marshaller string) {
	util.Glog.Debugf("➥  Connect to ⚓  [%s]", util.LogColorString("debug3", marshaller))
}

// ReadInvoices reads invoice-record from fn
func (c *InputFile) ReadInvoices() ([]*Invoice, error) {
	var unmarshaller InvoiceUnmarshaller
	util.DebugPrintCaller()
	// Pfstart(Format[Ffstart], CallerName(2))
	// pstat("file-type : %q\n", Opts.IfnSuffix)
	//
	var fb = FileBunker{Name: filepath.Base(c.Filename)}
	DB.Where(&fb).First(&fb)
	// Plog((&fb).GetArgsTable("", 0))
	glInfof("\n%s", (&fb).GetArgsTable("", 0))
	//
	switch c.Suffix {
	case ".csv":
		logdebugmarshaller("CsvMarshaller")
		unmarshaller = CsvMarshaller{}
	case ".jsn", ".json":
		logdebugmarshaller("JSONMarshaller")
		unmarshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		logdebugmarshaller("YAMLMarshaller")
		unmarshaller = YAMLMarshaller{}
	case ".xml":
		logdebugmarshaller("XMLMarshaller")
		unmarshaller = XMLMarshaller{}
	case ".xlsx":
		logdebugmarshaller("XlsxMarshaller")
		unmarshaller = XlsxMarshaller{}
	}
	if unmarshaller != nil {
		inpIsBig5 = c.IsBig5
		invs, err := unmarshaller.UnmarshalInvoices(os.ExpandEnv(c.Filename))
		// Stopfunc(Ffstop) //, "ReadInvoices")
		// glDebug(util.StrThickLine(60))
		return invs, err
	}
	return nil, fmt.Errorf("☠  not supprted file-type : %s (%s)",
		util.LogColorString("error", c.Suffix), util.LogColorString("error", c.Filename))
}

// WriteInvoices reads invoice-record from fn
func (o *OutputFile) WriteInvoices(invs []*Invoice) error {
	var marshaller InvoiceMarshaller
	// Startfunc(Ffstart) //, "ReadInvoices")
	util.DebugPrintCaller()
	switch o.Suffix {
	case ".csv":
		logdebugmarshaller("CsvMarshaller")
		marshaller = CsvMarshaller{}
	case ".jsn", ".json":
		logdebugmarshaller("JSONMarshaller")
		marshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		logdebugmarshaller("YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case ".xml":
		logdebugmarshaller("XMLMarshaller")
		marshaller = XMLMarshaller{}
	case ".xlsx":
		logdebugmarshaller("XlsxMarshaller")
		marshaller = XlsxMarshaller{}
	}
	if marshaller != nil {
		err := marshaller.MarshalInvoices(os.ExpandEnv(o.Filename), invs)
		return err
	}
	return fmt.Errorf("☠  not supprted file-type : %s (%s)",
		util.LogColorString("error", o.Suffix), util.LogColorString("error", o.Filename))
}
