package game

import "time"

var (
	messageSetTime  time.Time
	MessageDuration int = 3
	messageContent  string
)

func InputMessage(
	message string,
) {
	if message == "" {
		return
	}
	messageContent = message
	messageSetTime = time.Now()
}

func MessageResponse() (string, bool, int) {
	if messageContent == "" {
		return "", false, 0
	}

	elapsedMs := time.Since(messageSetTime).Milliseconds()
	remainingMs := (int64(MessageDuration) * 1000) - elapsedMs

	if remainingMs <= 0 {
		messageContent = ""
		return "", false, 0
	}

	return messageContent, true, int(remainingMs)
}
