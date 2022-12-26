package utils

import (
	"github.com/ahmdrz/goinsta/v2"
)

func CheckForNewPosts(insta *goinsta.Instagram, bot *tgbotapi.BotAPI, profile InstagramProfile) error {
	// Get the user's profile
	user, err := insta.Profiles.ByURL(profile.URL)
	if err != nil {
		return err
	}

	// Get the user's media
	media, err := user.RecentMedia()
	if err != nil {
		return err
	}

	// Get the first media item (which will be the latest post)
	latestPost := media.Items[0]

	// Send a message to the Telegram chat if the post is new
	if latestPost.CreatedTime > profile.LastCheckedTime {
		msg := tgbotapi.NewMessage(your-chat-id, "New post from "+profile.Username+":")
		if latestPost.IsVideo() {
			msg.Text += "\n" + latestPost.VideoURL
		} else {
			msg.Text += "\n" + latestPost.Images.StandardResolution.URL
		}
		_, err = bot.Send(msg)
		if err != nil {
			return err
		}

		// Update the LastCheckedTime for this profile
		profile.LastCheckedTime = latestPost.CreatedTime
	}

	return nil
}
