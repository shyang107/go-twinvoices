package invoices

import (
	"strings"

	"github.com/cpmech/gosl/io"
	jsoniter "github.com/json-iterator/go"
	"github.com/shyang107/go-twinvoices/util"
	// "github.com/shyang/invoices/inv/goini"
)

var (
	// Opt is option for this application
	Opt *Option
	// Opts are options for this application
	Opts []*Option

	// Cases are the settings of all cases
	Cases []*Case

	// DefaultOption sets a list of safe recommended option. Feel free to modify these to suit your needs.
	DefaultOption = Option{
		InpFn:       "./inp/09102989061.csv",
		IfnSuffix:   ".csv",
		IsNative:    true,
		IfnEncoding: "Big5",
		OutFn:       "./out/09102989061.json",
		OfnSuffix:   ".json",
		PunchFn:     "./out/punch.out",
	}

	// DefaultInput is default setting of InputFile
	DefaultInput = InputFile{Filename: "./inp/09102989061.csv", Suffix: ".csv", IsBig5: true}
	// DefaultOutput is default setting of OutputFile
	DefaultOutput = OutputFile{Filename: "./inp/09102989061.json", Suffix: ".json", IsWrite: true}
	// DefaultPunch is default setting of PunchFile
	DefaultPunch = PunchFile{Filename: "./inp/09102989061.log", Suffix: ".log", IsWrite: false}
	// DefaultCase is default setting of Case
	DefaultCase = Case{"", DefaultInput, DefaultOutput, DefaultPunch}
)

// InputFile describes input data
type InputFile struct {
	Filename string
	Suffix   string
	IsBig5   bool
}

// OutputFile describes the information needed to output file
type OutputFile struct {
	Filename string
	Suffix   string
	IsWrite  bool // wrtie result to file
}

// PunchFile describes the information needed to punch
type PunchFile struct {
	Filename string
	Suffix   string
	IsWrite  bool // wrtie result to file
}

// Case describes case to handle
type Case struct {
	Title  string
	Input  InputFile
	Output OutputFile
	Punch  PunchFile
}

func (c Case) String() string {
	strdashk := strings.Repeat("-", 15)
	strdashv := strings.Repeat("-", 30)
	return util.ArgsTable(
		"Case",
		"Input ------------", strdashk, strdashv,
		"file name", "Input.Filename", c.Input.Filename,
		"file type", "Input.Suffix", c.Input.Suffix,
		"Is Big-5 encoding?", "Input.IsBig5", c.Input.IsBig5,
		"Output -----------", strdashk, strdashv,
		"file name", "Output.Filename", c.Output.Filename,
		"file type", "Output.Suffix", c.Output.Suffix,
		"do output?", "Output.IsWrite", c.Output.IsWrite,
		"Punch ------------", strdashk, strdashv,
		"file name", "Punch.Filename", c.Punch.Filename,
		"file type", "Punch.Suffix", c.Punch.Suffix,
		"do output?", "Punch.IsWrite", c.Punch.IsWrite,
	)
}

// GetTable returns the string of arguments table
func (c Case) GetTable(title string) string {
	strdashk := strings.Repeat("-", 15)
	strdashv := strings.Repeat("-", 30)
	return util.ArgsTable(
		title,
		"Input ------------", strdashk, strdashv,
		"file name", "Input.Filename", c.Input.Filename,
		"file type", "Input.Suffix", c.Input.Suffix,
		"Is Big-5 encoding?", "Input.IsBig5", c.Input.IsBig5,
		"Output -----------", strdashk, strdashv,
		"file name", "Output.Filename", c.Output.Filename,
		"file type", "Output.Suffix", c.Output.Suffix,
		"do output?", "Output.IsWrite", c.Output.IsWrite,
		"Punch ------------", strdashk, strdashv,
		"file name", "Punch.Filename", c.Punch.Filename,
		"file type", "Punch.Suffix", c.Punch.Suffix,
		"do output?", "Punch.IsWrite", c.Punch.IsWrite,
	)
}

// NewCases return an new settings of cases
func NewCases() []*Case {
	return []*Case{&DefaultCase}
}

///////////////////////////////////////////////////////////////////////

// Option setup the application
type Option struct {
	// [input]
	InpFn       string `json:"input_filename"`
	IfnSuffix   string `json:"input_filename_extention"`
	IsNative    bool   `json:"is_native"`
	IfnEncoding string `json:"encoding_name_of_text"`
	// [output]
	OutFn     string `json:"output_filename"`
	OfnSuffix string `json:"output_filename_extention"`
	IsOutput  bool   `json:"is_output"`
	// [punch]
	PunchFn string `json:"punch_filename"`
}

func (o Option) String() string {
	strdashk := strings.Repeat("-", 15)
	strdashv := strings.Repeat("-", 30)
	tab := util.ArgsTable(
		"Option",
		"Input:", strdashk, strdashv,
		"input file name", "InpFn", o.InpFn,
		"input file type", "IfnSuffix", o.IfnSuffix,
		"is official invoices file from government?", "IsNative", o.IsNative,
		"encoding of input file", "INFencoding", o.IfnEncoding,
		"Output:", strdashk, strdashv,
		"output file name (if you want)", "OutFn", o.OutFn,
		"output file type", "OfnSuffix", o.OfnSuffix,
		"do output?", "IsOutput", o.IsOutput,
		"Punch:", strdashk, strdashv,
		"punch file name (not use for now)", "PunchFn", o.PunchFn,
	)
	return tab
}

// ReadOptions reads the configuration
func (Case) ReadOptions(cpath string) ([]*Option, error) {
	// startfunc(fostart)
	pstat("  > Reading options from .jsn or .json file %q ...\n", cpath)
	// if !isOpened(cpath) {
	// 	panic(chk.Err("config-file %q can not open", cpath))
	// }
	if util.IsFileExist(cpath) {
		util.Panic("config-file %q does not exist!", cpath)
	}
	//
	b, err := io.ReadFile(cpath)
	if err != nil {
		return nil, err
	}
	var opts []*Option
	err = jsoniter.Unmarshal(b, &opts)
	if err != nil {
		return nil, err
	}
	// stopfunc(fostop)
	return opts, nil
}

// NewOptions return an new Options of cases
func NewOptions() *Case {
	Opts = []*Option{&DefaultOption}
	return new(Case)
}
