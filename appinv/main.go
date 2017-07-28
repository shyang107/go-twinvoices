package main

import (
	"fmt"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/cmd"
	// "github.com/shyang107/go-twinvoices/cmd"
	"github.com/shyang107/go-twinvoices/util"
	// yaml "gopkg.in/yaml.v2"
)

// inv "github.com/shyang107/go-twinvoices"

func init() {
	log.SetPrefix("LOG: ")
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	// log.Println("init started")
	// io.Verbose = true
	// util.Verbose = true
}

func main() {
	start := time.Now()
	cmd.Execute()
	// outConfig("yconfig.yaml")
	// outCases("ycases.yaml")
	// readCases("./cases.yaml")
	util.Pf("run-time elapsed : %v\n", time.Since(start))
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
