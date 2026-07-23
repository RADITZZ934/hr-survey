package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TelegramPayload struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func SendTelegramAlert(score float64, surveyTitle string, respondentName string) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		log.Println("ℹ️ Telegram credentials not configured. Skipping external alert.")
		return
	}

	// Run in background to avoid blocking the HTTP response thread
	go func() {
		message := fmt.Sprintf(
			"⚠️ *ALERT KEPUASAN KRITIS*\n\n"+
				"*Survei:* %s\n"+
				"*Responden:* %s\n"+
				"*Rata-rata Nilai:* %.1f / 5.0\n"+
				"*Waktu:* %s\n\n"+
				"_Silakan periksa dashboard HR untuk analisis lebih lanjut._",
			surveyTitle,
			respondentName,
			score,
			time.Now().Format("02 Jan 2006, 15:04:05 WIB"),
		)

		payload := TelegramPayload{
			ChatID:    chatID,
			Text:      message,
			ParseMode: "Markdown",
		}

		body, err := json.Marshal(payload)
		if err != nil {
			log.Printf("❌ Failed to marshal telegram payload: %v\n", err)
			return
		}

		url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("❌ Failed to create telegram request: %v\n", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("❌ Failed to send telegram notification: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("❌ Telegram API returned non-OK status: %s\n", resp.Status)
			return
		}

		log.Println("✅ Critical score Telegram notification sent successfully!")
	}()
}
