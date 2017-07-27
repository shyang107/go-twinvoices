package main

import (
	"log"
	"time"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/cmd"
	"github.com/shyang107/go-twinvoices/util"
	yaml "gopkg.in/yaml.v2"
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
	// // outCases("ycases.yaml")
	// readCases("./ycases.yaml")
	util.Pf("run-time elapsed : %v\n", time.Since(start))
}
func readCases(fln string) {
	b, err := util.ReadFile(fln)
	if err != nil {
		log.Fatalln(err)
	}
	yaml.Unmarshal(b, &vp.Cases)
	for i, c := range vp.Cases {
		util.Pfgreen("%v", c.GetTable(util.Sf("Case #%d", i)))
	}
}

func outCases(fln string) {
	vp.Cases = []*vp.Case{&vp.DefaultCase, &vp.DefaultCase, &vp.DefaultCase}
	b, err := yaml.Marshal(vp.Cases)
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
