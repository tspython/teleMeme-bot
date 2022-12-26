package main

import (
	"log"
	"os"
	"time"

	"./utils"
	"github.com/ahmdrz/goinsta/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read the list of Instagram profiles from a JSON file
	profiles, err := utils.ReadProfilesFromJSON("profiles.json")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new goinsta client
	insta := goinsta.New("your-client-id", "your-client-secret")

	// Create a new Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// Start a loop to check for new posts every 60 seconds
	for {
		for _, profile := range profiles {
			// Check for new posts on this profile
			err = utils.CheckForNewPosts(insta, bot, profile)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		// Sleep for 60 seconds before checking for new posts again
		time.Sleep(60 * time.Second)
	}
}
