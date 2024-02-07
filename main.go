package neaucabot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	// Replace "YOUR_TELEGRAM_BOT_TOKEN" with the API token provided by the BotFather.
	bot, err := tgbotapi.NewBotAPI("6733181321:AAF0U6SalFQtN5rQwb1Eb6sGeuOmxP3SRbM")
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}
	// Set the bot to use debug mode (verbose logging).
	bot.Debug = true
	log.Printf("Authorized as @%s", bot.Self.UserName)
	// Set up updates configuration.
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// Get updates from the Telegram API.
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}
	// Process incoming messages.
	for update := range updates {
		if update.Message == nil { // Ignore any non-Message updates.
			continue
		}
		// Print received message text and sender username.
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		// Respond to the user.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! I am your Telegram bot.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}
