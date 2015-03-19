package login

import (
	"github.com/gogather/com"
)

func Login(username string, password string) {
	username = com.Md5(username)
}

func storeSess() {
	
}

func getSess() {
	
}