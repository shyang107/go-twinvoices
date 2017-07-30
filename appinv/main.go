package main

import (
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/cmd"
	"github.com/sirupsen/logrus"
	// "github.com/shyang107/go-twinvoices/cmd"
	"github.com/shyang107/go-twinvoices/util"
	// yaml "gopkg.in/yaml.v2"
	// "github.com/kataras/golog"
)

// inv "github.com/shyang107/go-twinvoices"
// var glog = golog.New()
var log = logrus.New()

func init() {
	// myloger := log.New(os.Stdout, "", 0)
	// myloger.SetPrefix("LOG: ")
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	// myloger.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	// log.Println("init started")
	// io.Verbose = true
	// util.Verbose = true
	// util.Glog.InstallStd(myloger)

	// simulate a logrus preparation:
	// logrus.SetLevel(logrus.InfoLevel)
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetFormatter(&logrus.TextFormatter{})
	// logrus.SetFormatter(&logrus.TextFormatter{})
	// pass logrus.StandardLogger() to print logs using using the default,
	// package-level logrus' instance of Logger:
	// util.Glog.Install(logrus.StandardLogger())

	///////////////////////////////////////////////////////////////////////////
	// file, err := os.OpenFile(util.TodayFilename(), os.O_CREATE|os.O_WRONLY, 0666)
	// file, err := util.NewLogFile()
	// // defer file.Close()
	// if err == nil {
	// 	util.Glog.AddOutput(file)
	// } else {
	// 	util.Glog.Errf("Failed to log to file, using default stderr")
	// }
}

func main() {
	start := time.Now()
	cmd.Execute()
	// outConfig("config.yaml")
	// outCases("ycases.yaml")
	// readCases("./cases.yaml")
	util.Glog.Println("\nrun-time elapsed: ", time.Since(start))
}

func readCases(fln string) {
	b, err := util.ReadFile(fln)
	if err != nil {
		log.Fatalln(err)
	}
	yaml.Unmarshal(b, &vp.Cases)
	for _, c := range vp.Cases {
		util.Pfgreen("%v", *c)
	}
}

func outCases(fln string) {
	sf := fmt.Sprintf
	var cases []*vp.Case
	for i := 0; i < 3; i++ {
		c := &vp.Case{
			Title: sf("CASE #%[1]d: [09102989061-%[1]d]", i+1),
			Input: vp.InputFile{
				Filename: sf("./inp/09102989061-%d.csv", i+1),
				Suffix:   ".csv",
				IsBig5:   true,
			},
			Outputs: []*vp.OutputFile{
				&vp.OutputFile{
					Filename: sf("./out/09102989061-%du.csv", i+1),
					Suffix:   ".csv",
					IsOutput: true,
				},
				&vp.OutputFile{
					Filename: sf("./out/09102989061-%d.json", i+1),
					Suffix:   ".json",
					IsOutput: true,
				},
				&vp.OutputFile{
					Filename: sf("./out/09102989061-%d.xml", i+1),
					Suffix:   ".xml",
					IsOutput: true,
				},
				&vp.OutputFile{
					Filename: sf("./out/09102989061-%d.xlsx", i+1),
					Suffix:   ".xlsx",
					IsOutput: true,
				},
				&vp.OutputFile{
					Filename: sf("./out/09102989061-%d.yaml", i+1),
					Suffix:   ".yaml",
					IsOutput: false,
				},
			},
		}
		cases = append(cases, c)
	}
	// vp.Cases = []*vp.Case{&vp.DefaultCase, &vp.DefaultCase, &vp.DefaultCase}
	b, err := yaml.Marshal(cases)
	if err != nil {
		log.Fatalln(err)
	}
	util.WriteBytesToFile(fln, b)
}

func outConfig(fln string) {
	b, err := yaml.Marshal(vp.GetDefualtConfig())
	if err != nil {
		log.Fatalln(err)
	}
	util.WriteBytesToFile(fln, b)
}
