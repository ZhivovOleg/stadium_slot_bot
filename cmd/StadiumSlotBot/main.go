package main

import (
	"StadiumSlotBot/internal/bot"
	"StadiumSlotBot/internal/dal"
	"StadiumSlotBot/internal/options"
	"StadiumSlotBot/internal/utils"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var Version string

//	@title						StadiumSlotBot API
//	@version					1.0
//	@description				Бот стадиона команды СЛОТ
//	@BasePath					/v1
// 	@externalDocs.description  	OpenAPI
// 	@externalDocs.url          	https://swagger.io/resources/open-api/
func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			arg = strings.TrimLeft(arg, "-")
			switch arg {
			case "v", "version", "Version" : fmt.Println(Version)
			}
		}
		return
	}

	isDebug := false
	if env, _ := os.LookupEnv("StadiumSlotBotEnv"); env == "dev" {
		isDebug = true		
	}

	utils.InitializeLogger(isDebug)
	defer utils.Logger.Sync()
	
	initSettingsErr := options.InitSettings(isDebug)

	if initSettingsErr != nil {
		utils.Logger.Error("Can't init settings: " + initSettingsErr.Error())
		panic("Can't init settings: " + initSettingsErr.Error())
	}

	utils.InitializeLogger(isDebug)
	defer utils.Logger.Sync()

	dbErr := dal.DB.Init(*options.StadiumSlotBotOptions.DBConnectionString)
	if dbErr != nil {
		utils.Logger.Error(dbErr.Error())
		panic(dbErr)
	}

	for {
		botErr := bot.InitBot(isDebug)
		if botErr != nil {
			utils.Logger.Error(botErr.Error())
			if strings.Contains(botErr.Error(), "getMe") {
				utils.Logger.Error("wait 3 seconds and retry")
				time.Sleep(time.Second*3)
				continue
			}
		}
		break
	}
	
	var wg sync.WaitGroup
	wg.Add(1)
	//go api.InitServer(*options.StadiumSlotBotOptions.Port, *options.StadiumSlotBotOptions.DBConnectionString, isDebug, &wg)
	go bot.StartBot(&wg)
	wg.Wait()
}