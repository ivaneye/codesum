package monitor

import "fmt"

//Monitor 为监视器，显示信息
type Monitor interface {
	Display(cont ...string)
}

// CmdMonitor 是命令行显示
type CmdMonitor struct{}

// Display 是CmdMonitor的Monitor接口实现
func (monitor *CmdMonitor) Display(cont ...string) {
	fmt.Println(cont)
}
