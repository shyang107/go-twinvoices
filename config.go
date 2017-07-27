package invoices

import (
	"os"

	"github.com/shyang107/go-twinvoices/util"
)

const (
	appversion  = "0.0.4"
	fileType    = "INVOICES" // using in text file
	magicNumber = 0x125D     // using in binary file
	fileVesion  = 100        // using in all filetype
	dateFormat  = "20060102" // allways using the date
	// ShortDateFormat is short date layout
	ShortDateFormat = "2006-01-02 MST" // allways using the date
	// LongDateFormat is long date layout
	LongDateFormat = "2006-01-02 15:04:05 -07:00 MST" // allways using the date
	// CfgFile is default config file
	CfgFile = ".invoices.yaml" // the path of config-file
)

var (
	// Cfg is configure
	Cfg *Config
	// DefaultConfig is default environment settings
	DefaultConfig = GetDefualtConfig()
)

// Config decribes configuration
type Config struct {
	DBfilename string `default:"./data/invoices.db" yaml:"DBfilename"`
	// IsInitializeDB = true to remove DBPath and create new database named DBPath
	IsInitialDB bool `default:"false" yaml:"IsInitialDB"`
	// Verbose activates display of messages on console
	Verbose bool `default:"false" yaml:"Verbose"`
	// ColorsOn activates use of colors on console
	ColorsOn bool `default:"true" yaml:"ColorsOn"`
	// IsDump = true, dumped all records from DBPath
	IsDump bool `default:"false" yaml:"IsDump"`
	// DumpFilename is the path dumped all records from DBPath
	DumpFilename string `default:"./all_invoices.json" yaml:"DumpFilename"`
	// CaseFilename is the case settings
	CaseFilename string `default:"./cases.yaml" yaml:"CaseFilename"`
}

func (c Config) String() string {
	tab := util.ArgsTable(
		"Configuration",
		"Filename of database", "DBfilename", c.DBfilename,
		"Does initalize database?", "IsInitialDB", c.IsInitialDB,
		"Activate display of messages on console", "Verbose", c.Verbose,
		"Activate use of colors on console", "ColorsOn", c.ColorsOn,
		"Does dump all records from database?", "IsDump", c.IsDump,
		"Filename saves all records from database", "DumpFilename", c.DumpFilename,
		"Filename of case", "CaseFilename", c.CaseFilename,
	)
	return tab
}

// GetDefualtConfig creates a new Env variable
func GetDefualtConfig() *Config {
	return &Config{
		DBfilename:   os.ExpandEnv("./data/invoices.db"),
		IsInitialDB:  false,
		Verbose:      false,
		IsDump:       false,
		DumpFilename: os.ExpandEnv("./all_invoices.json"),
		CaseFilename: os.ExpandEnv("./cases.yaml"),
	}
}
