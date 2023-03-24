package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var tgadminid, _ = strconv.ParseInt(os.Getenv("TELEGRAM_BOT_ADMIN_ID"), 10, 64)
var tgdebug, _ = strconv.ParseBool(os.Getenv("TELEGRAM_BOT_DEBUG_MODE"))

func status() string {
	uptime := "System Uptime: " + systemUptime() + "\n"
	memoryusage := "Memory Usage: " + strconv.Itoa(Memoryusage()) + "% \n"
	cpuusage := "CPU Usage: " + strconv.Itoa(Cpuusage()) + "% \n"
	return uptime + memoryusage + cpuusage
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = tgdebug
	log.Printf("Registered On BOT: %s", bot.Self.UserName)
	log.Printf("Admin ID: %d", tgadminid)
	log.Printf("DEBUG MODE: %t\n\n", tgdebug)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.Message.Chat.ID != tgadminid {
			msgAccessDenied := "Access Denied!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgAccessDenied)
			if _, err = bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = `Commands:
			- /status
			- /up
			- /reboot
			`
		case "up":
			msg.Text = "UP and Running"
		case "status":
			msg.Text = status()
		case "reboot":
			rebootSystem()
		default:
			msg.Text = "command not found."
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
