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
	util.Glog.Infof("\n%s", (&fb).GetArgsTable("", 0))
	//
	switch c.Suffix {
	case ".csv":
		// Pinfo("==> connect to %q\n", "CsvMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "CsvMarshaller")
		unmarshaller = CsvMarshaller{}
	case ".jsn", ".json":
		// Pinfo("==> connect to %q\n", "JSONMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "JSONMarshaller")
		unmarshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		// Pinfo("==> connect to %q\n", "YAMLMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "YAMLMarshaller")
		unmarshaller = YAMLMarshaller{}
	case ".xml":
		// Pinfo("==> connect to %q\n", "XMLMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "XMLMarshaller")
		unmarshaller = XMLMarshaller{}
		// case ".xlsx":
		// 	Pinfo("==> connect to %q\n", "XlsMarshaller")
		// 	util.Glog.Debugf("==> connect to %q", "XlsMarshaller")
		// 	unmarshaller = XlsMarshaller{}
	}
	if unmarshaller != nil {
		inpIsBig5 = c.IsBig5
		invs, err := unmarshaller.UnmarshalInvoices(os.ExpandEnv(c.Filename))
		// Stopfunc(Ffstop) //, "ReadInvoices")
		// util.Glog.Debug(util.StrThickLine(60))
		return invs, err
	}
	return nil, fmt.Errorf("not supprted file-type : %s (%s)",
		c.Suffix, c.Filename)
}

// WriteInvoices reads invoice-record from fn
func (o *OutputFile) WriteInvoices(invs []*Invoice) error {
	var marshaller InvoiceMarshaller
	// Startfunc(Ffstart) //, "ReadInvoices")
	util.DebugPrintCaller()
	switch o.Suffix {
	case ".csv":
		// Pinfo("==> connect to %q\n", "CsvMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "CsvMarshaller")
		marshaller = CsvMarshaller{}
	case ".jsn", ".json":
		// Pinfo("==> connect to %q\n", "JSONMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "JSONMarshaller")
		marshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		// Pinfo("==> connect to %q\n", "YAMLMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case ".xml":
		// Pinfo("==> connect to %q\n", "XMLMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "XMLMarshaller")
		marshaller = XMLMarshaller{}
	case ".xlsx":
		// Pinfo("==> connect to %q\n", "XlsMarshaller")
		util.Glog.Debugf("➠  connect to ⚓  %q", "XlsMarshaller")
		marshaller = XlsMarshaller{}
	}
	if marshaller != nil {
		err := marshaller.MarshalInvoices(os.ExpandEnv(o.Filename), invs)
		// Stopfunc(Ffstop) //, "ReadInvoices")
		// util.Glog.Debug(util.StrThickLine(60))
		return err
	}
	return fmt.Errorf("not supprted file-type : %s (%s)", o.Suffix, o.Filename)
}
