package invoices

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	ut "github.com/shyang107/go-twinvoices/util"
	// use for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is database
var DB *gorm.DB

// Initialdb initialize database
func Initialdb() error {
	ut.Verbose = true
	if ut.IsFileExist(Cfg.DBfilename) {
		pstat("  > Removing file %q ...\n", Cfg.DBfilename)
		err := os.Remove(Cfg.DBfilename)
		if err != nil {
			// panic(err)
			return err
		}
	}
	pstat("  > Creating file %q ...\n", Cfg.DBfilename)
	db, err := gorm.Open("sqlite3", os.ExpandEnv(Cfg.DBfilename))
	if err != nil {
		// ut.Panic("failed to connect database")
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
	//初始化并保持连接
	var err error
	DB, err = gorm.Open("sqlite3", Cfg.DBfilename)
	//    DB.LogMode(true)//打印sql语句
	if err != nil {
		log.Fatalf("database connect is err: %s", err.Error())
	} else {
		// log.Print("connect database is success")
		ut.Pfyel("* connect database is success\n")
	}
	err = DB.DB().Ping()
	if err != nil {
		DB.DB().Close()
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}
	DB.Model(&Invoice{}).Related(&Detail{}, "uin")
}

// DBDumpInvoices get the list from database
func DBDumpInvoices() ([]Invoice, error) {
	invs := []Invoice{}
	DB.Find(&invs)
	for i := range invs {
		// DB.Model(invs[i]).Related(&invs[i].UINumber)
		DB.Model(&invs[i]).Association("details").Find(&invs[i].Details)
	}
	return invs, nil
}

// DBInsertFrom creats records from []*Invoice into database
func DBInsertFrom(pvs []*Invoice) {
	for _, v := range pvs {
		// io.Pforan("# %v", *v)
		// DB.FirstOrCreate(v)
		DB.Where(Invoice{UINumber: v.UINumber}).FirstOrCreate(v)
	}
}

// DBDumpData dumps all data from db
func DBDumpData(dumpFilename string) error {
	pstat("  > Dumping data from database %q ...\n", Cfg.DBfilename)
	pvs, err := DBDumpInvoices()
	if err != nil {
		return err
	}
	// for i, p := range pvs {
	// 	io.Pfgreen2("Rec. %d : %v\n", i+1, p)
	// }
	fn := ut.PathKey(dumpFilename) + ".json"
	pstat("  >> Marshall data in JSON-type, and then write to %q ...\n", fn)
	b, err := jsoniter.Marshal(&pvs)
	if err != nil {
		return err
	}
	ut.WriteBytesToFile(fn, b)
	printSepline(60)
	return nil
}
