package invoices

import (
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/shyang107/go-twinvoices/util"
)

const (
	// Version is the version of this app
	Version     = "0.0.4"
	fileType    = "INVOICES" // using in text file
	magicNumber = 0x125D     // using in binary file
	fileVersion = 100        // using in all filetype
	dateFormat  = "20060102" // allways using the date
	// ShortDateFormat is short date layout
	ShortDateFormat = "2006-01-02 MST" // allways using the date
	// LongDateFormat is long date layout
	LongDateFormat = "2006-01-02 15:04:05 -07:00 MST" // allways using the date
	// CfgFile is default config file
	CfgFile = "./config.yaml" // the path of config-file
	//
)

var (
	// Cfg is configure
	Cfg *Config
	//
	Glog     = util.Glog
	glInfo   = Glog.Info
	glInfof  = Glog.Infof
	glWarn   = Glog.Warn
	glWarnf  = Glog.Warnf
	glError  = Glog.Error
	glErrorf = Glog.Errorf
	glDebug  = Glog.Debug
	glDebugf = Glog.Debugf
)

func init() {
	// Level defaults to "info",but you can change it:
	Glog.SetLevel("disable")
	// Glog.SetLevel("debug")

	// util.PfBlue("config.init called\n")
	Cfg = GetDefualtConfig()

	// util
	util.Verbose = Cfg.Verbose
	util.ColorsOn = Cfg.ColorsOn
}

// Config decribes configuration
type Config struct {
	DBfilename string `default:"./data/invoices.db" yaml:"DB_Filename"`
	// IsInitializeDB = true to remove DBPath and create new database named DBPath
	IsInitialDB bool `default:"false" yaml:"Is_Initial_DB"`
	// Verbose activates display of messages on console
	Verbose bool `default:"false" yaml:"Verbose"`
	// ColorsOn activates use of colors on console
	ColorsOn bool `default:"true" yaml:"ColorsOn"`
	// IsDump = true, dumped all records from DBPath
	IsDump bool `default:"false" yaml:"IsDump"`
	// DumpFilename is the path dumped all records from DBPath
	DumpFilename string `default:"./all_invoices.json" yaml:"Dump_Filename"`
	// CaseFilename is the case settings
	CaseFilename string `default:"./cases.yaml" yaml:"Case_Filename"`
}

func (c Config) String() string {
	tab := util.ArgsTable(
		"CONFIG",
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

// NewConfig return a new Config veriable
func NewConfig() (cfg *Config, err error) {
	if err := cfg.ReadConfigs(); err != nil {
		// util.Panic("%v\n", err)
		// fmt.Printf("%v\n", err)
		cfg = GetDefualtConfig()
		util.Verbose = cfg.Verbose
		return cfg, err
	}
	return cfg, err
}

// GetDefualtConfig creates a new Env variable
func GetDefualtConfig() *Config {
	return &Config{
		DBfilename:   os.ExpandEnv("./data/invoices.db"),
		IsInitialDB:  false,
		Verbose:      true,
		ColorsOn:     true,
		IsDump:       false,
		DumpFilename: os.ExpandEnv("./all_invoices.json"),
		CaseFilename: os.ExpandEnv("./cases.yaml"),
	}
}

// ReadConfigs read the configs
func (c *Config) ReadConfigs() error {
	util.DebugPrintCaller()
	// Prun("  > Reading configuration from  %q ...\n", CfgFile)
	glInfof("➥  Reading configuration from  %q ...", CfgFile)
	//
	if util.IsFileExist(CfgFile) {
		b, err := util.ReadFile(CfgFile)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(b, Cfg)
		if err != nil {
			return err
		}
	} else {
		b, err := yaml.Marshal(Cfg)
		if err != nil {
			return err
		}
		util.WriteBytesToFile(CfgFile, b)
	}
	util.Verbose = c.Verbose
	util.ColorsOn = c.ColorsOn
	// plog("Default config:\n%v\n", Cfg)
	// plog("Default configuration:\n")
	return nil
}
