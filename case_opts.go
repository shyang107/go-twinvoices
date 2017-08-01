package invoices

import (
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/shyang107/go-twinvoices/util"
	// "github.com/shyang/invoices/inv/goini"
)

var (
	// Cases are the settings of all cases
	Cases []*Case

	// DefaultInput is default setting of InputFile
	DefaultInput = InputFile{
		Filename: os.ExpandEnv("./inp/09102989061.csv"),
		Suffix:   ".csv",
		IsBig5:   true,
	}

	// DefaultOutput is default setting of OutputFile
	DefaultOutput = OutputFile{
		Filename: os.ExpandEnv("./inp/09102989061.json"),
		Suffix:   ".json",
		IsOutput: true,
	}

	// DefaultPunch is default setting of PunchFile
	// DefaultPunch = PunchFile{Filename: "./inp/09102989061.log", Suffix: ".log", IsOutput: false}

	// DefaultCase is default setting of Case
	DefaultCase = Case{"", DefaultInput, []*OutputFile{&DefaultOutput}}
	// DefaultPunch}
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
	IsOutput bool // wrtie result to file
}

// PunchFile describes the information needed to punch
type PunchFile struct {
	Filename string
	Suffix   string
	IsOutput bool // wrtie result to file
}

// Case describes case to handle
type Case struct {
	Title   string
	Input   InputFile
	Outputs []*OutputFile
	// Punch  PunchFile
}

func (c Case) String() string {
	strdashk := strings.Repeat("-", 15)
	strdashv := strings.Repeat("-", 30)
	tab := util.ArgsTable(
		c.Title,
		"Input ------------", strdashk, strdashv,
		"file name", "Input.Filename", c.Input.Filename,
		"file type", "Input.Suffix", c.Input.Suffix,
		"Is Big-5 encoding?", "Input.IsBig5", c.Input.IsBig5,
		"Output -----------", strdashk, "[as the following ...]",
		// "file name", "Output.Filename", c.Output.Filename,
		// "file type", "Output.Suffix", c.Output.Suffix,
		// "do output?", "Output.IsWrite", c.Output,
		// "Punch ------------", strdashk, strdashv,
		// "file name", "Punch.Filename", c.Punch.Filename,
		// "file type", "Punch.Suffix", c.Punch.Suffix,
		// "do output?", "Punch.IsWrite", c.Punch.IsWrite,
	)
	title := "OUTPUT List"
	heads := []string{"No", "Filename", "IsOutput"}
	var data []interface{}
	for i := 0; i < len(c.Outputs); i++ {
		of := c.Outputs[i]
		data = append(data, i+1, of.Filename, of.IsOutput)
	}
	tab += util.ArgsTableN(title, 4, false, heads, data...)
	return tab
}

// GetTable returns the string of arguments table
func (c Case) GetTable(title string) string {
	strdashk := strings.Repeat("-", 15)
	strdashv := strings.Repeat("-", 30)
	tab := util.ArgsTable(
		title,
		"Input ------------", strdashk, strdashv,
		"file name", "Input.Filename", c.Input.Filename,
		"file type", "Input.Suffix", c.Input.Suffix,
		"Is Big-5 encoding?", "Input.IsBig5", c.Input.IsBig5,
		"Output -----------", strdashk, "[as the following ...]",
		// "file name", "Output.Filename", c.Output.Filename,
		// "file type", "Output.Suffix", c.Output.Suffix,
		// "do output?", "Output.IsWrite", c.Output.IsWrite,
		// "Punch ------------", strdashk, strdashv,
		// "file name", "Punch.Filename", c.Punch.Filename,
		// "file type", "Punch.Suffix", c.Punch.Suffix,
		// "do output?", "Punch.IsWrite", c.Punch.IsWrite,
	)
	otitle := "OUTPUT List"
	heads := []string{"No", "Filename", "IsOutput"}
	var data []interface{}
	for i := 0; i < len(c.Outputs); i++ {
		of := c.Outputs[i]
		data = append(data, i+1, of.Filename, of.IsOutput)
	}
	tab += util.ArgsTableN(otitle, 4, false, heads, data...)
	return tab
}

// NewCases return an new settings of cases
func NewCases() []*Case {
	return []*Case{&DefaultCase}
}

// ReadCaseConfigs reads the configuration
func (c *Config) ReadCaseConfigs() ([]*Case, error) {
	util.DebugPrintCaller()
	fln := os.ExpandEnv(c.CaseFilename)
	suffix := util.FnExt(fln)
	glInfof("➥  Reading options from [%[2]s] file [%[1]s] ...",
		util.LogColorString("info", fln), util.LogColorString("info", suffix))
	b, err := util.ReadFile(fln)
	if err != nil {
		return nil, err
	}
	var cases []*Case
	// err = jsoniter.Unmarshal(b, &cases)
	err = yaml.Unmarshal(b, &cases)
	if err != nil {
		return nil, err
	}
	return cases, nil
}
