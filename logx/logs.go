package logx

import "fmt"

func Info(msg string) map[string]string {
	return map[string]string{"message": fmt.Sprintf(msg)}
}
