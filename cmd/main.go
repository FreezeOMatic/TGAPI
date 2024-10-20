package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"tgapi/internal/botcore"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	bot, err := botcore.InitWithContext(ctx, os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	err = bot.GetUpdateChan()
	if err != nil {
		log.Fatal(err)
	}
	defer bot.CloseUpdateChan()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		defer close(sigChan)

		select {
		case <-sigChan:
			cancel()
			return
		}
	}()
c:
	for {
		select {
		case update := <-bot.Updates():
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				reply, ok := bot.ProcessMessage(update)
				if !ok {
					continue
				}

				log.Printf("[%s] %v", update.Message.From.UserName, reply)

			}
		case <-ctx.Done():
			log.Println("stop signal received")
			break c
		}
	}
}
