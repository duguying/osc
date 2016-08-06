package tweet

import (
	"net/url"
	"path/filepath"
	"regexp"

	"github.com/duguying/osc/utils"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
)

func Tweet(message string) {
	// read info
	var data interface{}
	var err error

	home := utils.GetHome()
	pathUserInfo := filepath.Join(home, ".osc", "userinfo")
	if com.FileExist(pathUserInfo) {
		json, _ := com.ReadFile(pathUserInfo)
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

	cookiePath := filepath.Join(home, ".osc", "oscid")
	http := utils.NewHTTPClient(cookiePath)
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

func Joke() {
	api := `http://www.tuling123.com/openapi/api?key=380abd77ba6541dd1dee43220c42776b&info=%E8%AE%B2%E4%B8%AA%E7%AC%91%E8%AF%9D`
	home := utils.GetHome()
	cookiePath := filepath.Join(home, ".osc", "oscid")
	http := utils.NewHTTPClient(cookiePath)
	msg, err := http.Get(api)
	if err != nil {
		log.Redln(err)
	}

	data, err := com.JsonDecode(msg)
	if err != nil {
		log.Redln(err)
	}

	json := data.(map[string]interface{})
	msg = json["text"].(string)

	reg := regexp.MustCompile(`<[\d\D]+>`)
	msg = reg.ReplaceAllString(msg, "")

	msg = com.SubString(msg, 0, 190)

	Tweet(msg)
}

func Weather(location string) {
	api := `http://www.tuling123.com/openapi/api?key=380abd77ba6541dd1dee43220c42776b&info=%E4%BB%8A%E5%A4%A9` + location + `%E5%A4%A9%E6%B0%94`
	home := utils.GetHome()
	cookiePath := filepath.Join(home, ".osc", "oscid")
	http := utils.NewHTTPClient(cookiePath)
	msg, err := http.Get(api)
	if err != nil {
		log.Redln(err)
	}

	data, err := com.JsonDecode(msg)
	if err != nil {
		log.Redln(err)
	}

	json := data.(map[string]interface{})
	msg = json["text"].(string)

	reg := regexp.MustCompile(`<[\d\D]+>`)
	msg = reg.ReplaceAllString(msg, "")

	msg = com.SubString(msg, 0, 190)

	Tweet(msg)
}
