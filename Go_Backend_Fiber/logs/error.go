package logs

import "log"

func Error(s string,err error) {
	if err != nil {
		log.Println("❌🔥",s,err)
	}
}