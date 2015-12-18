package main

import (
	"github.com/ivaneye/codesum/monitor"
	"github.com/ivaneye/codesum/caculate"
	"github.com/ivaneye/codesum/conf"
	"github.com/ivaneye/codesum/filter"
	"github.com/ivaneye/codesum/iterator"
	"fmt"
	"os"
	"strconv"
	"os/exec"
	"path/filepath"
)

func main() {
	args := os.Args
	fmt.Println(args[1])
	if len(args) == 1 {
		fmt.Println("请输入统计目录!")
		return
	}
	conf := initConf(args)
	monitor := new(monitor.CmdMonitor)
	caculator := &caculate.Caculator{conf, monitor, 0}
	filte := &filter.FileNameFilter{conf.Suffix}
	iterat := iterator.Iterator{monitor, filte, caculator}
	iterat.Iterate(args[1])
	fmt.Println("sum=", caculator.Total)
}

func initConf(args []string) *conf.Manager {
	file,_ := exec.LookPath(args[0])
	path,_ := filepath.Abs(file)
	fmt.Println("file=",args[0])
	fmt.Println(path)
	if len(args) == 2 {
		//只有目录路径，使用默认配置
		return &conf.Manager{"//", "/\\*", "\\*/", "Action.java", false}
	}else {
		b, _ := strconv.ParseBool(args[6])
		return &conf.Manager{args[2], args[3], args[4], args[5], b}
	}
}
