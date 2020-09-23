package log

// 日志等级
const (
	levelOff    int = 0 //关闭
	levelEmerg  int = 1
	levelCrit   int = 2
	levelError  int = 3
	levelWarn   int = 4
	levelNotice int = 5
	levelInfo   int = 6
	levelDebug  int = 7
	levelTrace  int = 8
	levelOn     int = 9 //9 全部打开
)

var (
	strEmerg  string = "emerg"
	strCrit   string = "crit"
	strError  string = "error"
	strWarn   string = "warn"
	strNotice string = "ntice"
	strInfo   string = "info"
	strDebug  string = "debug"
	strTrace  string = "trace"
)
