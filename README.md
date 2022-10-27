# golog
## v1.0
简单的对logrus和file-rotatelogs进行封装
相关配置：
```go
type Options struct {
Path string
DirType string
Level logrus.Level
MaxAge time.Duration
MaxSize int64 // 单位：M
RotationTime time.Duration
LastLogName string
}

var defaultOpt = Options{
Path: "storage/log",
DirType: "/%Y/%m/%d/golog.%H%M.log",
Level: DebugLevel,
MaxSize: 15 * 1024 * 1024,
MaxAge: 365 * 24 * time.Hour,
RotationTime: 24 * time.Hour,
LastLogName: "golog.log",
}
```
## v1.1规划-未开始
支持日志写入三方数据库：redis、MongoDB、MySQL、ELK等
支持三种模式：只写入文件、只写入第三方、同时写入文件和第三方