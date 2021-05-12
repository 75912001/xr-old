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

var levelTag = []string{
	LevelOff:    "LevelOff",
	LevelEmerg:  "EME",
	LevelCrit:   "CRI",
	LevelError:  "ERR",
	LevelWarn:   "WAR",
	LevelNotice: "NOT",
	LevelInfo:   "INF",
	LevelDebug:  "GDB",
	LevelTrace:  "TRA",
	LevelOn:     "LevelOn",
}
