package main

import (
	"fmt"
	"log"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/cmds"
	"github.com/shyang107/go-twinvoices/util"
)

func init() {
	// util.EnableLoggerOutToFile("debug")
}

func main() {
	start := time.Now()
	cmds.Execute()
	// outConfig("config.yaml")
	// outCases("ycases.yaml")
	// readCases("./cases.yaml")
	util.Glog.Println("\n", "run-time elapsed: ", time.Since(start).String())
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
