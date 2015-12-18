package conf

//Manager 为配置，加载配置文件
type Manager struct {
	SingleLineComment     string
	MultiLineCommentStart string
	MultiLineCommentEnd   string
	Suffix                string
	Detail                bool
}
