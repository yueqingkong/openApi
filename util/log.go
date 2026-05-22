package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func Log() {
	dir := "./LOG"
	res2, err := MkDir(dir) //创建文件夹
	if res2 == false {
		panic(err)
	}
	file, _ := os.OpenFile(dir+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //打开日志文件，不存在则创建
	defer file.Close()

	log.SetOutput(file)                                 //设置输出流
	log.SetPrefix("[Error]")                            //日志前缀
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime) //日志输出样式
	log.Println("Hi file")
}

func LogFormat(p ...interface{}) string {
	var s bytes.Buffer
	for _, v := range p {
		switch v.(type) {
		case string:
			s.WriteString(v.(string))
		case float32:
			s.WriteString(fmt.Sprintf("%.3f, ", v.(float32)))
		case int:
			s.WriteString(fmt.Sprintf("%d, ", v.(int)))
		default:
			s.WriteString(fmt.Sprintf("%v", v))
		}
	}
	return s.String()
}
