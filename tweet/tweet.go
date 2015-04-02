package tweet

import (
	"github.com/duguying/osc/utils"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"net/url"
	"path/filepath"
)

func Tweet(message string) {
	// read info
	var data interface{}
	var err error

	home := utils.GetHome()
	pathUserInfo := filepath.Join(home, ".osc", "userinfo")
	if com.FileExist(pathUserInfo) {
		json := com.ReadFile(pathUserInfo)
		data, err = com.JsonDecode(json)
		if err != nil {
			log.Redln("[Error]", "Parse userinfo file failed")
			return
		}
	} else {
		log.Warnln("login first")
		return
	}

	jsonData, ok := data.(map[string]interface{})
	if !ok {
		log.Redln("[Error]", "illeage data")
		return
	}

	userId, ok := jsonData["user"].(string)
	if !ok {
		log.Redln("[Error]", "get user id failed")
		return
	}

	userCode, ok := jsonData["user_code"].(string)
	if !ok {
		log.Redln("[Error]", "get user code failed")
		return
	}

	http := &utils.Http{}
	response, err := http.Post("https://www.oschina.net/action/tweet/pub", url.Values{
		"user":      {userId},
		"user_code": {userCode},
		"msg":       {message},
	})

	if err != nil {
		log.Warnln("[Error]", err)
	}

	tweetResult, err := com.JsonDecode(response)
	if err != nil {
		log.Redln("发送失败")
		return
	} else {
		_, ok := tweetResult.(map[string]interface{})["error"].(float64)
		if ok {
			log.Warnln(tweetResult.(map[string]interface{})["msg"])
			return
		}

		id, ok := tweetResult.(map[string]interface{})["log"].(float64)
		if ok {
			log.Greenf("开源中国第[%d]条动态发送成功\n", int64(id))
			log.Blueln(message)
		}
	}

}
