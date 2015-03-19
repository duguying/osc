package utils

import (
	"github.com/gogather/com/log"
)

func GetHome() string {
	home, err := Home()
	if err!=nil {
		log.Fatalln("Can NOT find user path!")
	}
	return home
}
