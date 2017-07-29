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
	ut.DebugPrintCaller()
	// ut.Verbose = true
	if ut.IsFileExist(Cfg.DBfilename) {
		// Pstat("  > Removing file %q ...\n", Cfg.DBfilename)
		ut.Glog.Infof("> Removing file %q ...\n", Cfg.DBfilename)
		err := os.Remove(Cfg.DBfilename)
		if err != nil {
			// panic(err)
			return err
		}
	}
	// Pstat("  > Creating file %q ...\n", Cfg.DBfilename)
	ut.Glog.Infof("  > Creating file %q ...\n", Cfg.DBfilename)
	db, err := gorm.Open("sqlite3", os.ExpandEnv(Cfg.DBfilename))
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
	// ut.Glog.Debugf("* %q called by %q", ut.CallerName(1), ut.CallerName(2))
	ut.DebugPrintCaller()
	//初始化并保持连接
	var err error
	DB, err = gorm.Open("sqlite3", Cfg.DBfilename)
	//    DB.LogMode(true)//打印sql语句
	if err != nil {
		log.Fatalf("database connect is err: %s", err.Error())
	} else {
		// log.Print("connect database is success")
		// Pinfo("* connect database is success\n")
		ut.Glog.Debugf("> connect database is success")
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
	ut.DebugPrintCaller()
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
	ut.Glog.Infof(">> updating database ...")
	ut.DebugPrintCaller()
	for _, v := range pvs {
		// io.Pforan("# %v", *v)
		// DB.FirstOrCreate(v)
		DB.Where(Invoice{UINumber: v.UINumber}).FirstOrCreate(v)
	}
}

// DBDumpData dumps all data from db
func DBDumpData(dumpFilename string) error {
	ut.DebugPrintCaller()
	// Prun("  > Dumping data from %q ...\n", Cfg.DBfilename)
	ut.Glog.Infof(">> Dumping data from %q ...", Cfg.DBfilename)
	pvs, err := DBGetAllInvoices()
	if err != nil {
		return err
	}
	return DBWriteInvoices(pvs, dumpFilename)
}

// DBWriteInvoices write all invoices to the file
func DBWriteInvoices(invs []*Invoice, fln string) error {
	ut.DebugPrintCaller()
	fln = os.ExpandEnv(fln)
	// fn := PathKey(fln) // + ".json"
	ext := ut.FnExt(fln)
	// Prun("  >> Prepare %[2]q data, and then write to %[1]q ...\n", fln, ext)
	ut.Glog.Infof(">> Prepare %[2]q data, and then write to %[1]q ...", fln, ext)
	var marshaller InvoiceMarshaller
	switch ext {
	case ".csv":
		// Pinfo("==> connect to %q\n", "CsvMarshaller")
		ut.Glog.Debugf("==> connect to %q", "CsvMarshaller")
		marshaller = CsvMarshaller{}
	case ".jsn", ".json":
		// Pinfo("==> connect to %q\n", "JSONMarshaller")
		ut.Glog.Debugf("==> connect to %q", "JSONMarshaller")
		marshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		// Pinfo("==> connect to %q\n", "YAMLMarshaller")
		ut.Glog.Debugf("==> connect to %q", "YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case ".xml":
		// Pinfo("==> connect to %q\n", "XMLMarshaller")
		ut.Glog.Debugf("==> connect to %q", "XMLMarshaller")
		marshaller = XMLMarshaller{}
	case ".xlsx":
		// Pinfo("==> connect to %q\n", "XlsMarshaller")
		ut.Glog.Debugf("==> connect to %q", "XlsMarshaller")
		marshaller = XlsMarshaller{}
	}
	if marshaller != nil {
		err := marshaller.MarshalInvoices(fln, invs)
		return err
	}
	return chk.Err("not supprted %[2]q-type (%[1]q)", fln, ext)
}
