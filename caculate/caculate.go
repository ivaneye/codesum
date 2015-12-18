package caculate

import (
	"os"
	"regexp"
	"strconv"

	"github.com/ivaneye/codesum/conf"
	"github.com/ivaneye/codesum/monitor"
	"fmt"
)

// Caculate 为统计接口，负责实际的统计工作
type Caculate interface {
	Caculate(fileInfo os.FileInfo, cont []byte) int
}

//Caculator 为计算结构体
type Caculator struct {
	*conf.Manager
	Monitor monitor.Monitor
	Total   int
}

// Caculate 是Caculator的Caculate接口实现
func (caculator *Caculator) Caculate(fileInfo os.FileInfo, cont []byte) int {
	//过滤掉所有空行，注释。统计换行符
	//考虑重构
	regstr := caculator.SingleLineComment + `[.\s\r\n]*`
	cont = trim(regstr, cont, "")
	regstr = caculator.MultiLineCommentStart + `[^` + caculator.MultiLineCommentEnd + `]*` + caculator.MultiLineCommentEnd
	regstr = `/\*{1,2}[\s\S]*?\*/[\s\r\n]*`
	cont = trim(regstr, cont, "")
	regstr = `[\r\n]+\s*[\r\n]*`
	cont = trim(regstr, cont, "/")
	regstr = `/`
	r := regexp.MustCompile(regstr)
	lines := len(r.FindAll(cont, -1))
	if caculator.Detail {
		caculator.Monitor.Display("读取路径", fileInfo.Name(), "行数为:", strconv.Itoa(lines))
	}
	return lines
}

func trim(reg string, cont []byte, replace string) []byte {
	r := regexp.MustCompile(reg)
	return r.ReplaceAll(cont, []byte(replace))
}
