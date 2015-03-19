package login

import (
	"net/url"
	"path/filepath"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"github.com/duguying/osc/utils"
)

func Login(username string, password string) {
	home := utils.GetHome()
	pathPwd := filepath.Join(home, ".osc", "password")

	password = utils.SHA1(password)
	com.WriteFile(pathPwd, password)

	pathUsr := filepath.Join(home, ".osc", "username")
	com.WriteFile(pathUsr, username)

	
	http := &utils.Http{}
	response := http.Post("https://www.oschina.net/action/user/hash_login",url.Values{
		"email" : {username},
		"pwd" : {password},
		"save_login" : {"1"},
		})

	response = http.Get("http://my.oschina.net/duguying/admin/portrait")
	log.Greenln(response)
}

func storeSess() {
	
}

func getSess() {
	
}