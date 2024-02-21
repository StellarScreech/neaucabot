package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Replace "YOUR_TELEGRAM_BOT_TOKEN" with the API token provided by the BotFather.
	bot, err := tgbotapi.NewBotAPI("6733181321:AAF0U6SalFQtN5rQwb1Eb6sGeuOmxP3SRbM")
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	// Bot Settings
	prefix := "-"
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
	// User whitelist
	admins := [6]tgbotapi.User{
		{ID: 557161506}, // stellarscreech
		{ID: 842719267}, // partyislife
		//{ID: }, // ___
	}

	// Process incoming messages.
	for update := range updates {
		if update.Message == nil { // Ignore any non-Message updates.
			continue
		}

		// Check if the sender is whitelisted.
		isWhitelisted := false
		for _, admin := range admins {
			if update.Message.From.ID == admin.ID {
				isWhitelisted = true
				break
			}
		}

		// Print received message text and sender username.
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if isWhitelisted {
			m := update.Message
			//----------------Commands---------------//

			// wl - Check if the sender is whitelisted
			if strings.HasPrefix(m.Text, prefix+"wl") {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "true")

				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}

			// mute [user] [length] [reason]
			if strings.HasPrefix(m.Text, prefix+"mute") {
				// Split the message into parts: command, user, length, and reason
				parts := strings.Fields(m.Text)
				if len(parts) < 3 {
					// Invalid command format
					msg := tgbotapi.NewMessage(m.Chat.ID, "Invalid command format. Use: "+prefix+"mute [length] [reason] and reply to your target.")
					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("Error sending message: %v", err)
					}
					continue
				}

				userToMute := m.ReplyToMessage.From.UserName
				muteLength := parts[1]
				muteReason := strings.Join(parts[2:], " ")
				// Calculate mute duration based on muteLength
				if err != nil {
					msg := tgbotapi.NewMessage(m.Chat.ID, "Invalid duration format. Please specify a valid duration or 'forever'.")
					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("Error sending message: %v", err)
					}
					continue
				}
				// Convert
				muteDuration, _ := strconv.Atoi(muteLength)

				// Logic
				boolval := new(bool)
				*boolval = false

				restrictMemberCfg := tgbotapi.RestrictChatMemberConfig{
					ChatMemberConfig: tgbotapi.ChatMemberConfig{
						ChatID: update.Message.Chat.ID,
						UserID: m.ReplyToMessage.From.ID,
					},
					CanSendMessages: boolval,             // Prevent the user from sending messages
					UntilDate:       int64(muteDuration), // Mute duration
				}
				_, err = bot.RestrictChatMember(restrictMemberCfg)
				if err != nil {
					log.Printf("Error muting user: %v", err)
					continue
				}

				// Indicate that the user was muted.
				muteMessage := "User @" + userToMute + " has been muted for " + muteLength + " for: " + muteReason
				msg := tgbotapi.NewMessage(m.Chat.ID, muteMessage)
				_, err = bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}

			// unmute [user]
			if strings.HasPrefix(m.Text, prefix+"unmute") {
				parts := strings.Fields(m.Text)
				if len(parts) < 3 {
					// Invalid command format
					msg := tgbotapi.NewMessage(m.Chat.ID, "Invalid command format.")
					_, err := bot.Send(msg)
					if err != nil {
						log.Printf("Error sending message: %v", err)
					}
					continue
				}

				userToMute := m.ReplyToMessage.From.UserName

				// Logic
				boolval := new(bool)
				*boolval = true

				restrictMemberCfg := tgbotapi.RestrictChatMemberConfig{
					ChatMemberConfig: tgbotapi.ChatMemberConfig{
						ChatID: update.Message.Chat.ID,
						UserID: m.ReplyToMessage.From.ID,
					},
					CanSendMessages: boolval, // Make the member be able to send msgs.
				}

				_, err = bot.RestrictChatMember(restrictMemberCfg)
				if err != nil {
					log.Printf("Error unmuting user: %v", err)
					continue
				}

				// Indicate that the user was unmuted.
				muteMessage := "User @" + userToMute + " has been unmuted."
				msg := tgbotapi.NewMessage(m.Chat.ID, muteMessage)
				_, err = bot.Send(msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}

		}
	}
}
