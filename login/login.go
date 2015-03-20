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
	response, err := http.Post("https://www.oschina.net/action/user/hash_login",url.Values{
		"email" : {username},
		"pwd" : {password},
		"save_login" : {"1"},
		})

	if err!=nil {
		log.Warnln("请检查网络")
		return
	}

	// log.Blueln(response)
	json, err := com.JsonDecode(response)
	if err==nil {
		msg, _ := json.(map[string]interface{})["msg"].(string)
		
		errorCode, ok := json.(map[string]interface{})["error"].(float64)
		if ok {
			log.Redf("error[%d] %s,%s\n", int(errorCode), msg, "请去网页版登录")
			return
		}

		failCount, ok := json.(map[string]interface{})["failCount"].(float64)
		if ok {
			log.Redln(msg)
			log.Warnln("你还有", 3 - int(failCount), "次尝试的机会")
		}else{
			log.Redln("Invalid Response")
		}
	}else{
		log.Greenln("登录成功")
	}
}

func storeSess() {
	
}

func getSess() {
	
}