package invoices

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/shyang107/go-twinvoices/util"
	// use for sqlite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB is database
var DB *gorm.DB

// Initialdb initialize database
func Initialdb() error {
	util.DebugPrintCaller()
	// util.Verbose = true
	if util.IsFileExist(Cfg.DBfilename) {
		// Pstat("  > Removing file %q ...\n", Cfg.DBfilename)
		glInfof("☞  Removing file %q ...\n", Cfg.DBfilename)
		err := os.Remove(Cfg.DBfilename)
		if err != nil {
			// panic(err)
			return err
		}
	}
	// Pstat("  > Creating file %q ...\n", Cfg.DBfilename)
	glInfof("♲  Creating file %q ...\n", Cfg.DBfilename)
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
		glDebugf("♲  Connect database is success")
	}
	err = DB.DB().Ping()
	if err != nil {
		DB.DB().Close()
		log.Fatalf("☠  Error on opening database connection: %s", err.Error())
	}
	DB.Model(&Invoice{}).Related(&Detail{}, "uin")
}

// DBGetAllInvoices get the list from database
func DBGetAllInvoices() ([]*Invoice, error) {
	util.DebugPrintCaller()
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
	glInfof("♲  Updating database ...")
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
	glInfof("♲  Dumping data from %q ...", Cfg.DBfilename)
	pvs, err := DBGetAllInvoices()
	if err != nil {
		return err
	}
	return DBWriteInvoices(pvs, dumpFilename)
}

// DBWriteInvoices write all invoices to the file
func DBWriteInvoices(invs []*Invoice, fln string) error {
	util.DebugPrintCaller()
	fln = os.ExpandEnv(fln)
	// fn := PathKey(fln) // + ".json"
	ext := util.FnExt(fln)
	// Prun("  ➾  Prepare %[2]q data, and then write to %[1]q ...\n", fln, ext)
	glInfof("➾  Prepare %[2]q data, and then write to %[1]q ...", fln, ext)
	var marshaller InvoiceMarshaller
	switch ext {
	case ".csv":
		// Pinfo("==> connect to %q\n", "CsvMarshaller")
		glDebugf("➥  Connect to ⚓  %q", "CsvMarshaller")
		marshaller = CsvMarshaller{}
	case ".jsn", ".json":
		// Pinfo("==> connect to %q\n", "JSONMarshaller")
		glDebugf("➥  Connect to ⚓  %q", "JSONMarshaller")
		marshaller = JSONMarshaller{}
	case ".yml", ".yaml":
		// Pinfo("==> connect to %q\n", "YAMLMarshaller")
		glDebugf("➥  Connect to ⚓  %q", "YAMLMarshaller")
		marshaller = YAMLMarshaller{}
	case ".xml":
		// Pinfo("==> connect to %q\n", "XMLMarshaller")
		glDebugf("➥  Connect to ⚓  %q", "XMLMarshaller")
		marshaller = XMLMarshaller{}
	case ".xlsx":
		// Pinfo("==> connect to %q\n", "XlsMarshaller")
		glDebugf("➥  Connect to ⚓  %q", "XlsMarshaller")
		marshaller = XlsMarshaller{}
	}
	if marshaller != nil {
		err := marshaller.MarshalInvoices(fln, invs)
		return err
	}
	return fmt.Errorf("☠  Not supprted %[2]q-type (%[1]q)", fln, ext)
}
