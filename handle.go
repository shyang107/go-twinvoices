package invoices

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	util "github.com/shyang107/gout"
)

// Suffix contains file types
var Suffix = struct {
	CSV, JSN, JSON, YML, YAML, XML, XLSX string
}{
	CSV: ".csv", JSN: ".jsn", JSON: ".json", YML: ".yml", YAML: ".yaml",
	XML: ".xml", XLSX: ".xlsx",
}

var (
	inpIsBig5 = false
	cacheMu   sync.Mutex
)

// InvoiceMarshaller is marshal-operator of invoices
type InvoiceMarshaller interface {
	// MarshalInvoices(fn string, pvs []*Invoice) error
	MarshalInvoices(fn string, pvs *InvoiceCollection) error
}

// InvoiceUnmarshaller is unmarshal-operator of invoices
type InvoiceUnmarshaller interface {
	// UnmarshalInvoices(fn string) ([]*Invoice, error)
	UnmarshalInvoices(fn string) (*InvoiceCollection, error)
}

func logdebugmarshaller(marshaller string) {
	util.Glog.Debugf("➥  Connect to ⚓  [%s]", util.LogColorString("debug3", marshaller))
}

// ReadInvoices reads invoice-record from fn
func (c *InputFile) ReadInvoices() (*InvoiceCollection, error) {
	// cacheMu.Lock()
	// defer cacheMu.Unlock()

	util.DebugPrintCaller()

	var unmarshaller InvoiceUnmarshaller
	var fb = &FileBunker{Name: filepath.Base(c.Filename)}
	DB.Where(fb).First(fb)

	util.Glog.Infof("\n%s", fb.GetArgsTable("", 0))

	switch c.Suffix {
	case Suffix.CSV:
		logdebugmarshaller("CsvMarshaller")
		unmarshaller = CsvMarshaller{}
	case Suffix.JSN, Suffix.JSON:
		logdebugmarshaller("JSONMarshaller")
		unmarshaller = JSONMarshaller{}
	case Suffix.YML, Suffix.YAML:
		logdebugmarshaller("YAMLMarshaller")
		unmarshaller = YAMLMarshaller{}
	case Suffix.XML:
		logdebugmarshaller("XMLMarshaller")
		unmarshaller = XMLMarshaller{}
	case Suffix.XLSX:
		logdebugmarshaller("XlsxMarshaller")
		unmarshaller = XlsxMarshaller{}
	}
	if unmarshaller != nil {
		inpIsBig5 = c.IsBig5
		invs, err := unmarshaller.UnmarshalInvoices(os.ExpandEnv(c.Filename))
		return invs, err
	}
	return nil, fmt.Errorf("☠  not supprted file-type : %s (%s)",
		util.LogColorString("error", c.Suffix), util.LogColorString("error", c.Filename))
}

// WriteInvoices reads invoice-record from fn
func (o *OutputFile) WriteInvoices(invs *InvoiceCollection) error {
	// cacheMu.Lock()
	// defer cacheMu.Unlock()

	util.DebugPrintCaller()

	var marshaller InvoiceMarshaller

	switch o.Suffix {
	case Suffix.CSV:
		logdebugmarshaller("CsvMarshaller")
		marshaller = CsvMarshaller{}
	case Suffix.JSN, Suffix.JSON:
		logdebugmarshaller("JSONMarshaller")
		marshaller = JSONMarshaller{}
	case Suffix.YML, Suffix.YAML:
		logdebugmarshaller("YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case Suffix.XML:
		logdebugmarshaller("XMLMarshaller")
		marshaller = XMLMarshaller{}
	case Suffix.XLSX:
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
