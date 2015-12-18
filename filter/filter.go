package filter

import "strings"

// Filter 为过滤接口，过滤不符合条件的文件
type Filter interface {
	Filte(name string) bool
}

//FileNameFilter 是依据文件名进行过滤的过滤器
type FileNameFilter struct {
	Suffix string
}

//Filte 是FileNameFilter的Filter接口实现方法
func (filter *FileNameFilter) Filte(name string) bool {
	return strings.HasSuffix(name, filter.Suffix)
}
