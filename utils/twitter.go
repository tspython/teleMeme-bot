// Package utils contains utility functions for the Twitter Parser
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/g8rswimmer/go-twitter/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Profile represents an Twitter profile
type Profile struct {
	URL string `json:"url"`
}

// CheckForNewTweets checks for new tweets on the specified Twitter profile and sends a message to the specified
// Telegram chat if any new tweets are found
func CheckForNewTweets(client *twitter.Client, bot *tgbotapi.BotAPI, chatID int64, profileURL string, lastCheckedTime time.Time) error {
	// Parse the profile URL
	parsedURL, err := url.Parse(profileURL)
	if err != nil {
		return err
	}

	// Get the username from the profile URL
	username := parsedURL.Path[1:]

	// Get the user's tweets
	tweets, err := client.Timeline.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: username,
		SinceID:    lastCheckedTime.Unix(),
	})
	if err != nil {
		return err
	}

	// Check if there are any new tweets
	if len(tweets) == 0 {
		return nil
	}

	// Send a message to the Telegram chat for each new tweet
	for _, tweet := range tweets {
		// Update the last checked time to the time of the most recent tweet
		lastCheckedTime = tweet.CreatedAt

		// Send the tweet text to the Telegram chat
		msg := tgbotapi.NewMessage(chatID, tweet.Text)
		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

// ReadLastCheckedTime reads the last checked time from a JSON file
func ReadLastCheckedTime(filename string) (time.Time, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return time.Time{}, err
	}

	// Unmarshal the JSON data into a time.Time
	var lastCheckedTime time.Time
	if err := json.Unmarshal(data, &lastCheckedTime); err != nil {
		return time.Time{}, err
	}

	return lastCheckedTime, nil
}

// WriteLastCheckedTime writes the current time to a JSON file.
func WriteLastCheckedTime(t time.Time, filename string) error {
	// Marshal the time into JSON data
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	return nil
}
