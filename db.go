package invoices

import (
	"log"
	"os"

	"github.com/cpmech/gosl/chk"
	"github.com/jinzhu/gorm"
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

// DBGetAllInvoices get the list from database
func DBGetAllInvoices() ([]*Invoice, error) {
	invs := []*Invoice{}
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
	pvs, err := DBGetAllInvoices()
	if err != nil {
		return err
	}
	return DBWriteInvoices(pvs, dumpFilename)
}

// DBWriteInvoices write all invoices to the file
func DBWriteInvoices(invs []*Invoice, fln string) error {
	fln = os.ExpandEnv(fln)
	// fn := ut.PathKey(fln) // + ".json"
	ext := ut.FnExt(fln)
	pstat("  >> Marshall data in %[2]q, and then write to %[1]q ...\n", fln, ext)
	var marshaller InvoiceMarshaller
	switch ext {
	case ".csv":
		pstat("%q\n", "CsvMarshaller")
		marshaller = CsvMarshaller{}
	case ".jsn", ".json":
		pstat("%q\n", "JSONMarshaller")
		marshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		pstat("%q\n", "YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case ".xml":
		pstat("%q\n", "XMLMarshaller")
		marshaller = XMLMarshaller{}
	case ".xlsx":
		pstat("%q\n", "XlsMarshaller")
		marshaller = XlsMarshaller{}
	}
	if marshaller != nil {
		err := marshaller.MarshalInvoices(fln, invs)
		return err
	}
	return chk.Err("not supprted %[2]q-type (%[1]q)", fln, ext)
}
