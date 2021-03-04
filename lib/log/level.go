package log

// 日志等级
const (
	LevelOff    int = 0 //关闭
	LevelEmerg  int = 1
	LevelCrit   int = 2
	LevelError  int = 3
	LevelWarn   int = 4
	LevelNotice int = 5
	LevelInfo   int = 6
	LevelDebug  int = 7
	LevelTrace  int = 8
	LevelOn     int = 9 //9 全部打开
)

var (
	strEmerg  string = "EME"
	strCrit   string = "CRI"
	strError  string = "ERR"
	strWarn   string = "WAR"
	strNotice string = "NOT"
	strInfo   string = "INF"
	strDebug  string = "GDB"
	strTrace  string = "TRA"
)
