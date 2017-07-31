package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kataras/golog"

	yaml "gopkg.in/yaml.v2"

	"github.com/fatih/color"
	"github.com/kataras/pio"
	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/sirupsen/logrus"
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
	// 	util.Glog.Errorf("Failed to log to file, using default stderr")
	// }
}

func main() {
	start := time.Now()
	// cmds.Execute()
	// outConfig("config.yaml")
	// outCases("ycases.yaml")
	// readCases("./cases.yaml")
	testColors()
	testColors2()
	testColors3()
	// testlog()
	// for i := uint(0); i < 9; i++ {
	// 	fmt.Printf("1 << %d : %v\n", i, 1<<i)
	// 	fmt.Printf("(1 + %d) << 8 : %v\n", i, (1+i)<<8)
	// }
	println("\n\n", "run-time elapsed: ", time.Since(start).String())
}

func testlog() {
	util.Glog.SetLevel("debug")
	var log golog.ExternalLogger
	log = util.Glog
	log.Info("info")
	log.Error("Error")
	log.Warn("Warn")
	log.Debug("Debug")
}

func testColors3() {
	util.Pl()
	util.Pf("*** io.Pf... ***\n")
	util.Pfcyan("util.Pfcyan\n")
	util.Pfcyan2("util.Pfcyan2\n")
	util.Pfyel("util.Pfyel\n")
	util.Pfdyel("util.Pfdyel\n")
	util.Pfdyel2("util.Pfdyel2\n")
	util.Pfred("util.Pfred\n")
	util.Pfgreen("util.Pfgreen\n")
	util.Pfgreen2("util.Pfgreen2\n")
	util.Pfdgreen("util.Pfdgreen\n")
	util.Pfblue("util.Pfblue\n")
	util.Pfblue2("util.Pfblue2\n")
	util.Pfmag("util.Pfmag\n")
	util.Pfpink("util.Pfpink\n")
	util.Pfpurple("util.Pfpurple\n")
	util.Pfgrey("util.Pfgrey\n")
	util.Pfgrey2("util.Pfgrey2\n")
	util.Pforan("util.Pforan\n")
	util.Pf("\n*** io.Pf...: high intensity ***\n")
	util.PfCyan("util.PfCyan\n")
	util.PfYel("util.PfYel\n")
	util.PfGreen("util.PfGreen\n")
	util.PfBlue("util.PfBlue\n")
	util.PfMag("util.PfMag\n")
	util.PfWhite("util.PfWhite\n")
}

func testColors2() {
	p := pio.NewTextPrinter("color", os.Stdout)
	p.Println("*** pio ***")
	p.Println(pio.Blue("pio.Blue: this is a blue text"))
	p.Println(pio.Gray("pio.Gray: this is a gray text"))
	p.Println(pio.Red("pio.Red: this is a red text"))
	p.Println(pio.Purple("pio.Purple: this is a purple text"))
	p.Println(pio.Yellow("pio.Yellow: this is a yellow text"))
	p.Println(pio.Green("pio.Green: this is a green text"))
}

func testColors() {
	color.Black("color.Black:%v ", "This is testing the colorful text!")
	color.HiBlack("color.HiBlack:%v ", "This is testing the colorful text!")
	color.Blue("color.Blue:%v ", "This is testing the colorful text!")
	color.HiBlue("color.HiBlue:%v ", "This is testing the colorful text!")
	color.Cyan("color.Cyan:%v ", "This is testing the colorful text!")
	color.HiCyan("color.HiCyan:%v ", "This is testing the colorful text!")
	color.Green("color.Green:%v ", "This is testing the colorful text!")
	color.HiGreen("color.HiGreen:%v ", "This is testing the colorful text!")
	color.Magenta("color.Magenta:%v ", "This is testing the colorful text!")
	color.HiMagenta("color.HiMagenta:%v ", "This is testing the colorful text!")
	color.Red("color.Red:%v ", "This is testing the colorful text!")
	color.HiRed("color.HiRed:%v ", "This is testing the colorful text!")
	color.White("color.White:%v ", "This is testing the colorful text!")
	color.HiWhite("color.HiWhite:%v ", "This is testing the colorful text!")
	color.Yellow("color.Yellow:%v ", "This is testing the colorful text!")
	color.HiYellow("color.HiYellow:%v ", "This is testing the colorful text!")
	cl := color.New(color.BgGreen, color.FgBlack)
	cl.Print("color.New(color.BgRed,color.FgYellow) ", "This is testing the colorful text!", "\n")
	util.Pforan("%v %v\n", "Pforan", "This is testing the colorful text!")
}

func readCases(fln string) {
	// cfg := vp.GetDefualtConfig()
	cfg := &vp.Config{CaseFilename: os.ExpandEnv(fln)}
	cases, err := cfg.ReadCaseConfigs()
	if err != nil {
		util.Glog.Error(err.Error())
	}
	util.Verbose = true
	for _, c := range cases {
		util.PfGreen("PfGreen:\n%v", *c)
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
