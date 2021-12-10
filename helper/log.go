package helper

import (
	"fmt"
	"time"
)

// Log 记录日志
func Log(format string, args ...interface{}) {
	time.Sleep(time.Millisecond * 100)
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", time.Now().Format("2006-01-02 15:04:05.000"), msg)
}
