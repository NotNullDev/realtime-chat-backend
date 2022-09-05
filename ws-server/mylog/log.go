package mylog

import "log"

const errorPrefix = "ERROR: "
const infoPrefix = "INFO: "

func Info(msg interface{}) {
	log.Printf("%s%v", infoPrefix, msg)
}

func Error(msg interface{}) {
	log.Printf("%s%v", errorPrefix, msg)
}
