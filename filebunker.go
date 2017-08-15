package invoices

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shyang107/go-twinvoices/util"
)

// FileBunker use to backup original file of invoices
type FileBunker struct {
	Model    gorm.Model `json:"-" yaml:"-" gorm:"embedded"`
	Name     string     `cht:"檔案名稱" json:"NAME" yaml:"NAME"`
	Size     int        `cht:"檔案大小" json:"SIZE" yaml:"SIZE"`
	ModAt    time.Time  `cht:"修改時間" json:"MODTIME_AT" yaml:"MODTIME_AT" sql:"index"` // modification time
	Encoding string     `cht:"編碼" json:"ENCODING" yaml:"ENCODING"`
	Checksum string     `cht:"檢查碼" json:"CHECKSUM" json:"CHECKSUM"` //sha256
	Contents []byte     `cht:"內容" json:"-" yaml:"-"`
}

var (
	fileBunkerFieldNames []string
	fileBunkerCtagNames  []string
	fileBunkerIndeces    = make(map[string]int)
)

func init() {
	var err error
	if fileBunkerFieldNames, err = util.Names(&FileBunker{}, "Model"); err != nil {
		util.Panic("retrive field names failed (%q)!", "fileBunkerFieldNames")
	}
	if fileBunkerCtagNames, err = util.NamesFromTag(&FileBunker{}, "cht", "Model"); err != nil {
		util.Panic("retrive field-tag names failed [(%q).Tag(%q)]!", "fileBunkerCtagNames", "cht")
	}
	for i := 0; i < len(fileBunkerFieldNames); i++ {
		fileBunkerIndeces[fileBunkerFieldNames[i]] = i
	}
}

func (f FileBunker) String() string {
	Sf := fmt.Sprintf
	location, _ := time.LoadLocation("Local")
	val := reflect.ValueOf(f) //.Elem()
	fld := val.Type()
	var str string
	var cols = make([]string, 0)
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		f := fld.Field(i)

		switch f.Name {
		case fileBunkerFieldNames[fileBunkerIndeces["Model"]]:
			continue
		case fileBunkerFieldNames[fileBunkerIndeces["ModAt"]]:
			str = Sf("%v", v.Interface().(time.Time).In(location))
		case fileBunkerFieldNames[fileBunkerIndeces["Size"]]:
			str = util.BytesSizeToString(v.Interface().(int))
		case fileBunkerFieldNames[fileBunkerIndeces["Contents"]]:
			str = "[略...]"
		default:
			// str = v.Interface().(string)
			str = util.Sf("%v", v.Interface().(string))
		}
		cols = append(cols, Sf("%s:%s", fileBunkerCtagNames[fileBunkerIndeces[f.Name]], str))
	}
	return strings.Join(cols, csvSep)
}

// TableName : set Detail's table name to be `details`
func (FileBunker) TableName() string {
	// custom table name, this is default
	return "filebunker"
}

// GetArgsTable :
func (f *FileBunker) GetArgsTable(title string, lensp int) string {
	util.DebugPrintCaller()
	// Sf := fmt.Sprintf
	location, _ := time.LoadLocation("Local")
	if len(title) == 0 {
		title = "原始發票檔案清單"
	}
	// var heads = []string{"項次"}
	// _, _, _, heads := util.GetFieldsInfo(FileBunker{}, "cht", "Model")
	if lensp < 0 {
		lensp = 0
	}
	// heads = append(heads, tmp...)
	strSize := util.BytesSizeToString(f.Size)
	table := util.ArgsTableN(title, lensp, true, fileBunkerCtagNames,
		f.Name, strSize, f.ModAt.In(location), f.Encoding, f.Checksum, "[略...]")
	return table
}

// GetFileBunkersTable returns the table string of the list of []*Detail
func GetFileBunkersTable(pfbs []*FileBunker, lensp int) string {
	util.DebugPrintCaller()
	Sf := fmt.Sprintf
	location, _ := time.LoadLocation("Local")
	title := "原始發票檔案清單"
	heads := []string{"項次"}
	// _, _, _, tmp := util.GetFieldsInfo(FileBunker{}, "cht", "Model")
	heads = append(heads, fileBunkerCtagNames...)
	if lensp < 0 {
		lensp = 0
	}
	var data []interface{}
	for i, f := range pfbs {
		strSize := util.BytesSizeToString(f.Size)
		data = append(data, i+1,
			f.Name, strSize, Sf("%v", f.ModAt.In(location)), f.Encoding, f.Checksum, "[略...]")
	}
	table := util.ArgsTableN(title, lensp, false, heads, data...)
	return table
}

// UpdateFileBunker updates DB
func (c *Case) UpdateFileBunker() error {
	util.DebugPrintCaller()
	fi, err := os.Stat(c.Input.Filename)
	if err != nil {
		return err
	}
	if strings.ToLower(c.Input.Suffix) == Suffix.CSV && c.Input.IsBig5 {
		b, err := util.ReadFile(c.Input.Filename)
		if err != nil {
			return err
		}
		sum := fmt.Sprintf("%x", sha256.Sum256(b))
		fn := filepath.Base(c.Input.Filename)
		fb := FileBunker{
			Name:     fn,
			Size:     int(fi.Size()),
			ModAt:    fi.ModTime(),
			Encoding: "big-5",
			Checksum: sum,
			Contents: b,
		}
		DB.Where(&fb).FirstOrCreate(&fb)
		// fbs = append(fbs, &fb)
	}
	// plog((&fb).GetArgsTable("", 0))
	return nil
}

//=========================================================

// FileBunkerCollection is the collection of "*FileBunker"
type FileBunkerCollection []*FileBunker

func (p FileBunkerCollection) String() string {
	return p.GetFileBunkersTable()
}

// GetFileBunkersTable returns the table string of the list of []*FileBunker
func (p *FileBunkerCollection) GetFileBunkersTable() string {
	return GetFileBunkersTable(([]*FileBunker)(*p), 0)
}

// GetAllOriginalDataFromDB gets all original records from database
func GetAllOriginalDataFromDB() *FileBunkerCollection {
	util.DebugPrintCaller()

	fbs := []*FileBunker{}
	DB.Find(&fbs)
	var fbc FileBunkerCollection = fbs
	return &fbc
}

// GetAllFromDB gets all original records from database
func (p *FileBunkerCollection) GetAllFromDB() {
	p = GetAllOriginalDataFromDB()
}
