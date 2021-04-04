package kernel

import "log"

func ErrorHandler(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
