package utilities

import "log"

func CheckError(err error, str string) {
	if err != nil {
		log.Printf("%s:\t%s\n", str, err.Error())
	}
}

func PanicError(err error, str string) {
	if err != nil {
		log.Fatalf("%s:\t%s\n", str, err.Error())
	}
}
