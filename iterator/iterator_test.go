package iterator

import (
	"fmt"
	"testing"

	"github.com/ivaneye/codesum/caculate"
	"github.com/ivaneye/codesum/conf"
	"github.com/ivaneye/codesum/filter"
	"github.com/ivaneye/codesum/monitor"
)

func TestIterator(t *testing.T) {
	monitor := new(monitor.CmdMonitor)
	conf := &conf.Manager{"//", "/\\*", "\\*/", "java", false}
	caculator := &caculate.Caculator{conf, monitor, 0}
	filte := &filter.FileNameFilter{conf.Suffix}
	iterat := Iterator{monitor, filte, caculator}
	iterat.Iterate("E:\\code\\fpay-gateway")
	fmt.Println("sum=", caculator.Total)
}
