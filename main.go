package main

import (
	"github.com/ivaneye/codesum/monitor"
	"github.com/ivaneye/codesum/caculate"
	"github.com/ivaneye/codesum/conf"
	"github.com/ivaneye/codesum/filter"
	"github.com/ivaneye/codesum/iterator"
	"fmt"
	"flag"
)

func main() {
	lang := flag.String("lang", "java", "Language Type:java,go,c...")
	path := flag.String("path", "", "Path to Caculate")
	single := flag.String("single", "//", "Single Line Comment")
	prefix := flag.String("prefix", "/\\*", "Multi Line Comment Prefix")
	suffix := flag.String("suffix", "\\*/", "Multi Line Comment Suffix")
	detail := flag.Bool("detail", false, "Should show Detail?")
	flag.Parse()
	fmt.Println("lang=", *lang, ",path=", *path, ",single=", *single, ",prefix=", *prefix, ",suffix=", *suffix)
	if *path == "" {
		fmt.Println("Please input path to Caculate!")
		return
	}
	conf := initConf(*lang, *path, *single, *prefix, *suffix, *detail)
	monitor := new(monitor.CmdMonitor)
	caculator := &caculate.Caculator{conf, monitor, 0}
	filte := &filter.FileNameFilter{conf.Suffix}
	iterat := iterator.Iterator{monitor, filte, caculator}
	iterat.Iterate(*path)
	fmt.Println("sum=", caculator.Total)
}

func initConf(lang, path, single, prefix, suffix string, detail bool) *conf.Manager {
	return &conf.Manager{single, prefix, suffix, lang, detail}
}
