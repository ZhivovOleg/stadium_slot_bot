package bot

import (
	"StadiumSlotBot/internal/dal"
	"StadiumSlotBot/internal/options"
	"StadiumSlotBot/internal/utils"
	"context"
	"errors"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Отправка сообщений об ошибке администратору
func sendErrorMessageToAdmin(message string) {
	_, err := stadiumBot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: options.StadiumSlotBotOptions.AdminChatId,
		Text:   message,
	})
	if err != nil {
		utils.Logger.Error(err.Error())
	}
}

func sendBaseErrorAnswerToUser(chatId int64) {
	_, err := stadiumBot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID: chatId,
		Text:   "У нас проблемы на сервере, придется вам обратиться к боту попозже",
	})
	if err != nil {
		utils.Logger.Error(err.Error())
	}
}

// Базовый обработчик ошибок бота
func errorsHandler(err error) {
	utils.Logger.Error(err.Error())
	sendErrorMessageToAdmin(err.Error())
}

// Обработчик неуказанного текста/команд.
func textHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   base_answer,
	})
	if err != nil {
		utils.Logger.Error(err.Error())
		sendErrorMessageToAdmin(err.Error())
	}
}

// Получить информацию по стадиону.
func commandInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := dal.InitQuery(dal.DB, &ctx)
	q = q.From("stadium_slot_info").Columns([]string{"message_text"}).Where("message_name = 'stadium_coordinates' OR message_name = 'stadium_navi'").Select()

	if q.Error != nil {
		utils.Logger.Error(q.Error.Error())
		sendErrorMessageToAdmin(q.Error.Error())
		sendBaseErrorAnswerToUser(update.Message.Chat.ID)
	}

	switch q.Result.(type) {
		case []string: {
			for _, row := range q.Result.([]string) {
				if len(row) > 1 {
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID:		update.Message.Chat.ID,
						Text:		row,
						ParseMode: 	models.ParseModeHTML,
					})
					if err != nil {
						utils.Logger.Error(err.Error())
						sendErrorMessageToAdmin(err.Error())
					}
				}
			}
		}
		default: {
			err := errors.New("ошибка при получении данных из БД")
			utils.Logger.Error(err.Error())
			sendErrorMessageToAdmin(err.Error())
			sendBaseErrorAnswerToUser(update.Message.Chat.ID)
		}
	}
}

// Регистрация на соревнование
func commandRaceRegHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:		update.Message.Chat.ID,
		Text:		race_reg,
		ParseMode: 	models.ParseModeHTML,
	})
	if err != nil {
		utils.Logger.Error(err.Error())
		sendErrorMessageToAdmin(err.Error())
	}
}

// Информация по прокату
func commandRentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := dal.InitQuery(dal.DB, &ctx)
	q = q.From("stadium_slot_info").Columns([]string{"message_text"}).Where("message_name = 'stadium_rent'").Select()

	if q.Error != nil {
		utils.Logger.Error(q.Error.Error())
		sendErrorMessageToAdmin(q.Error.Error())
		sendBaseErrorAnswerToUser(update.Message.Chat.ID)
	}

	switch q.Result.(type) {
		case []string: {
			for _, row := range q.Result.([]string) {
				if len(row) > 1 {
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID:		update.Message.Chat.ID,
						Text:		row,
						ParseMode: 	models.ParseModeHTML,
					})
					if err != nil {
						utils.Logger.Error(err.Error())
						sendErrorMessageToAdmin(err.Error())
					}
				}
			}
		}
		default: {
			err := errors.New("ошибка при получении данных из БД")
			utils.Logger.Error(err.Error())
			sendErrorMessageToAdmin(err.Error())
			sendBaseErrorAnswerToUser(update.Message.Chat.ID)
		}
	}
}

// Информация о тренировках
func commandTrainHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:		update.Message.Chat.ID,
		Text:		training_info,
		ParseMode: 	models.ParseModeHTML,
		LinkPreviewOptions: &models.LinkPreviewOptions{IsDisabled: bot.True()},
	})
	if err != nil {
		utils.Logger.Error(err.Error())
		sendErrorMessageToAdmin(err.Error())
	}
}

// Результаты соревнований
func commandRaceResultsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:		update.Message.Chat.ID,
		Text:		race_results_info,
		ParseMode: 	models.ParseModeHTML,
	})
	if err != nil {
		utils.Logger.Error(err.Error())
		sendErrorMessageToAdmin(err.Error())
	}
}

// Информация о предстоящих гонках на стадионе
func commandRaceInfoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	q := dal.InitQuery(dal.DB, &ctx)
	q = q.From("stadium_slot_races").Columns([]string{"race_date", "race_name", "race_info"}).Where("race_date::date < NOW()::date").Select()

	if q.Error != nil {
		utils.Logger.Error(q.Error.Error())
		sendErrorMessageToAdmin(q.Error.Error())
		sendBaseErrorAnswerToUser(update.Message.Chat.ID)
	}

	switch q.Result.(type) {
		case []string: {
			for _, row := range q.Result.([]string) {
				if len(row) > 1 {
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID:		update.Message.Chat.ID,
						Text:		row,
						ParseMode: 	models.ParseModeHTML,
					})
					if err != nil {
						utils.Logger.Error(err.Error())
						sendErrorMessageToAdmin(err.Error())
					}
				}
			}
		}
		default: {
			err := errors.New("ошибка при получении данных из БД")
			utils.Logger.Error(err.Error())
			sendErrorMessageToAdmin(err.Error())
			sendBaseErrorAnswerToUser(update.Message.Chat.ID)
		}
	}
}