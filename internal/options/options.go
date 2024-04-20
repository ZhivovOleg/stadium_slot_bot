package options

import (
	"fmt"
	"os"
)

type Options struct {
	Port *string				`json:"stadiumSlotBotPort"`
	DBConnectionString *string	`json:"stadiumSlotBotDbConnectionString"`
	BotToken *string			`json:"botToken"`
	AdminChatId *string			`json:"adminChatId"`
}

var StadiumSlotBotOptions *Options = new(Options)

// InitSettings - initialize setting from ENV or appSettings.json
func InitSettings(isDebug bool) error {
	port, portExistsOk := os.LookupEnv("StadiumSlotBotPort")
	dbConnectionString, dbConnectionStringOk := os.LookupEnv("StadiumSlotBotDbConnectionString")	
	adminChatId, adminChatIdOk := os.LookupEnv("StadiumSlotBotAdminId")

	if !portExistsOk || !dbConnectionStringOk || !adminChatIdOk {
		err := fmt.Errorf("не удалось получить настройки приложения")
		return err
	}

	StadiumSlotBotOptions.AdminChatId = &adminChatId
	StadiumSlotBotOptions.Port = &port
	StadiumSlotBotOptions.DBConnectionString = &dbConnectionString
	
	return nil
}