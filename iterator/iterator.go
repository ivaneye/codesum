package iterator

import (
	"io/ioutil"

	"github.com/ivaneye/codesum/caculate"
	"github.com/ivaneye/codesum/filter"
	"github.com/ivaneye/codesum/monitor"
)

// Iteratable 为迭代器接口，根据路径与Filter遍历文件
type Iteratable interface {
	Iterate(path string)
}

//Iterator 是接口实现
type Iterator struct {
	Monitor   monitor.Monitor
	Filter    filter.Filter
	Caculator *caculate.Caculator
}

//Iterate 用于迭代传递路径下的所有文件
func (iterator *Iterator) Iterate(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		iterator.Monitor.Display("读取路径", path, "失败!", err.Error())
		return
	}
	for _, f := range files {
		if f.IsDir() {
			iterator.Iterate(path + "/" + f.Name())
		} else {
			if iterator.Filter.Filte(f.Name()) {
				b, er := ioutil.ReadFile(path + "/" + f.Name())
				if er != nil {
					iterator.Monitor.Display("读取文件", path, "失败!", err.Error())
					return
				}
				iterator.Caculator.Total += iterator.Caculator.Caculate(f, b)
			}
		}
	}
}
