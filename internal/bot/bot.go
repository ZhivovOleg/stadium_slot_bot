package bot

import (
	"context"
	"sync"

	"github.com/go-telegram/bot"
)

var stadiumBot *bot.Bot

func InitBot(isDebug bool) error {	
	opts := []bot.Option{
		bot.WithDefaultHandler(textHandler),
		bot.WithErrorsHandler(errorsHandler),
		bot.WithAllowedUpdates([]string{"/info", "/rent", "/training", "/race_info", "/race_reg", "/race_results"}),
	}

	if isDebug {
		opts = append(opts, bot.WithDebug())
	}

	b, botErr := bot.New("7080475284:AAEZfPgaeNcKzwtr8CadsAVZyxyqmqOnEJ0", opts...)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/info", bot.MatchTypeExact, commandInfoHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/rent", bot.MatchTypeExact, commandRentHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/training", bot.MatchTypeExact, commandTrainHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/race_info", bot.MatchTypeExact, commandRaceInfoHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/race_reg", bot.MatchTypeExact, commandRaceRegHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/race_results", bot.MatchTypeExact, commandRaceResultsHandler)
	stadiumBot = b
	return botErr
}

func StartBot(wg *sync.WaitGroup) {
	//ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer def(cancel, wg)
	stadiumBot.Start(context.Background())
}

func def(cancel context.CancelFunc, wg *sync.WaitGroup) {
	cancel()
	wg.Done()
}