package utils

import (
	"log"
	"os"
)

func Logger(err error) {
	if err != nil {
		f, errFile := os.OpenFile("dot-hiring-go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if errFile != nil {
			log.Fatalf("error opening file: %v", errFile)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(err.Error())
	}
}
