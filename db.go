package invoices

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	util "github.com/shyang107/gout"
	// use for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is database
var DB *gorm.DB

// Initialdb initialize database
func Initialdb() error {
	util.DebugPrintCaller()
	// util.Verbose = true
	// if util.IsFileExist(Cfg.DBfilename) {
	// 	// Pstat("  > Removing file %q ...\n", Cfg.DBfilename)
	// 	Glog.Infof("☞  Removing file %q ...\n", Cfg.DBfilename)
	// 	err := os.Remove(Cfg.DBfilename)
	// 	if err != nil {
	// 		// panic(err)
	// 		return err
	// 	}
	// }
	fl := os.ExpandEnv(Cfg.DBfilename)
	_, err := os.Stat(fl)
	Glog.Debug(err)
	if !os.IsNotExist(err) {
		Glog.Infof("♲  Removing file %q ...\n", fl)
		err := os.RemoveAll(fl)
		if err != nil {
			// panic(err)
			util.DebugPrintCaller()
			Glog.Error(err)
			return err
		}
	}
	// Pstat("  > Creating file %q ...\n", Cfg.DBfilename)
	Glog.Infof("♲  Creating file %q ...\n", fl)
	db, err := gorm.Open("sqlite3", fl)
	if err != nil {
		// Panic("failed to connect database")
		return err
	}
	defer db.Close()
	//
	// Migrate the schema
	db.AutoMigrate(&Invoice{}, &Detail{}, &FileBunker{})
	db.Model(&Invoice{}).Related(&Detail{}, "uin")
	// db.Model(&Invoice{}).AddUniqueIndex("idx_invoices_number", "uin")
	// db.Model(&Invoice{}).AddForeignKey("uin", "details(id)", "RESTRICT", "RESTRICT")
	return nil
}

// Connectdb connect the database
func Connectdb() {
	// glDebugf("* %q called by %q", util.CallerName(1), util.CallerName(2))
	util.DebugPrintCaller()
	//初始化并保持连接
	var err error
	DB, err = gorm.Open("sqlite3", Cfg.DBfilename)
	//    DB.LogMode(true)//打印sql语句
	if err != nil {
		log.Fatalf("☠  Database connect is err: %s", err.Error())
	} else {
		// log.Print("connect database is success")
		// Pinfo("* connect database is success\n")
		Glog.Debugf("♲  Connect database is success")
	}
	err = DB.DB().Ping()
	if err != nil {
		DB.DB().Close()
		log.Fatalf("☠  Error on opening database connection: %s", err.Error())
	}
	DB.Model(&Invoice{}).Related(&Detail{}, "uin")
}

// DBGetAllInvoices get the list from database
func DBGetAllInvoices() (*InvoiceCollection, error) {
	util.DebugPrintCaller()
	invs := []*Invoice{}
	DB.Find(&invs)
	for _, p := range invs {
		DB.Model(&p).Association("details").Find(&p.Details)
	}
	var res InvoiceCollection = invs
	return &res, nil
}

// DBGetAllInvoices get the list from database
func (v *InvoiceCollection) DBGetAllInvoices() error {
	util.DebugPrintCaller()
	invs := []*Invoice{}
	DB.Find(&invs)
	for _, p := range invs {
		DB.Model(&p).Association("details").Find(&p.Details)
	}
	var res InvoiceCollection = invs
	v = &res
	return nil
}

// dbInsertFrom creats records from []*Invoice into database
func dbInsertFrom(pvs []*Invoice) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	Glog.Infof("♲  Updating database ...")
	util.DebugPrintCaller()
	for _, v := range pvs {
		// io.Pforan("# %v", *v)
		// DB.FirstOrCreate(v)
		DB.Where(Invoice{UINumber: v.UINumber}).FirstOrCreate(v)
	}
}

// DBDumpData dumps all data from db
func DBDumpData(dumpFilename string) error {
	util.DebugPrintCaller()
	// Prun("  > Dumping data from %q ...\n", Cfg.DBfilename)
	Glog.Infof("♲  Dumping data from %q ...", Cfg.DBfilename)
	pvs, err := DBGetAllInvoices()
	if err != nil {
		return err
	}

	// return dbWriteInvoices(pvs, dumpFilename)
	out := &OutputFile{
		Filename: os.ExpandEnv(dumpFilename),
		Suffix:   util.FnExt(dumpFilename),
		IsOutput: true,
	}

	return out.WriteInvoices(pvs)
}
