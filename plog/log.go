package plog

import (
	"log"
	"os"
)

//SetLog设置log存储文件的位置
func SetLog(path string, prefix string) {
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix(prefix)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

}
