package initial

import (
  "os"
  "github.com/gogather/com/log"
  "github.com/duguying/osc/login"
  "github.com/duguying/osc/tweet"
  "fmt"
)

const (
	VERSION = "0.0.0"
)

func Run() {
	length := len(os.Args)

	username := ""
	password := ""
	message := ""
	
	if length < 2 {
		showHelp()
	}else{
		if os.Args[1] == "login" {
			if length < 4 {
				log.Dangerln("Invalid command, please use")
				log.Warnln(  "    osc login username password")
			}else{
				username = os.Args[2]
				password = os.Args[3]

				login.Login(username, password)
			}
		}else if os.Args[1] == "tweet" {
			if length < 3 {
				log.Dangerln("Invalid command, please use")
				log.Warnln(  "    osc tweet message")
			}else{
				message = os.Args[2]

				tweet.Tweet(message)
			}
		}else if os.Args[1] == "help" {
			showHelp()
		}else{
			log.Dangerln("Invalid command, please use")
				log.Warnln(  "    osc help")
		}
	}
}

func showHelp() {
	log.Warnln(  "oschina command line tool")
	fmt.Println( "Usage:")
	fmt.Println( "    login user password")
	fmt.Println( "    tweet message")
	log.Blueln(  "version",VERSION)
}

func readConfig() {

}
