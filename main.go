package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/g8rswimmer/go-twitter/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/tspython/teleMeme-bot/utils"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Create a new Twitter client
	client, err := twitter.Client(&twitter.ClientConfig{
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Telegram bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// Get the Telegram chat ID
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Read the list of profiles from a JSON file
	profiles, err := readProfiles("profiles.json")
	if err != nil {
		log.Fatal(err)
	}

	// Read the last checked time from a JSON file
	lastCheckedTime, err := utils.ReadLastCheckedTime("last_checked.json")
	if err != nil {
		log.Fatal(err)
	}

	// Set the frequency of checks
	checkFrequency := time.Minute * 30

	// Start an infinite loop to check for new posts at the specified frequency
	for {
		// Check for new posts on each profile
		for _, profile := range profiles {
			if err := utils.CheckForNewTweets(client, bot, chatID, profile.URL, lastCheckedTime); err != nil {
				log.Println(errors.Wrap(err, "error checking for new tweets"))
			}
		}

		// Update the last checked time
		lastCheckedTime = time.Now()
		if err := utils.WriteLastCheckedTime(lastCheckedTime, "last_checked.json"); err != nil {
			log.Println(errors.Wrap(err, "error writing last checked time"))
		}

		// Sleep
		// Sleep for the specified frequency before checking again
		time.Sleep(checkFrequency)
	}
}

// readProfiles reads a list of profiles from a JSON file.
func readProfiles(filename string) ([]utils.Profile, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of profiles
	var profiles []utils.Profile
	if err := json.Unmarshal(data, &profiles); err != nil {
		return nil, err
	}

	return profiles, nil
}
