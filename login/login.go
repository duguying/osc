package login

import (
	"path/filepath"
	"github.com/gogather/com"
	"github.com/duguying/osc/utils"
)

func Login(username string, password string) {
	home := utils.GetHome()
	path := filepath.Join(home, ".osc", "password")

	password = com.Md5(password)
	com.WriteFile(path, password)
}

func storeSess() {
	
}

func getSess() {
	
}