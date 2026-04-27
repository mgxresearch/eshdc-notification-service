package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ZeptoMailPayload struct {
	From struct {
		Address string `json:"address"`
	} `json:"from"`
	To []struct {
		EmailAddress struct {
			Address string `json:"address"`
			Name    string `json:"name"`
		} `json:"email_address"`
	} `json:"to"`
	Subject string `json:"subject"`
	HTMLBody string `json:"htmlbody"`
}

func SendEmail(to, name, subject, body string) (string, error) {
	apiKey := os.Getenv("ZEPTOMAIL_API_KEY")
	senderEmail := os.Getenv("ZEPTOMAIL_SENDER_ADDRESS")
	
	if apiKey == "" {
		return "", fmt.Errorf("ZEPTOMAIL_API_KEY is not set")
	}

	payload := ZeptoMailPayload{
		Subject: subject,
		HTMLBody: body,
	}
	payload.From.Address = senderEmail
	payload.To = append(payload.To, struct {
		EmailAddress struct {
			Address string `json:"address"`
			Name    string `json:"name"`
		} `json:"email_address"`
	}{
		EmailAddress: struct {
			Address string `json:"address"`
			Name    string `json:"name"`
		}{
			Address: to,
			Name:    name,
		},
	})

	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.zeptomail.com/v1.1/email", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Zoho-enczapikey "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to send email: status code %d", resp.StatusCode)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	// ZeptoMail returns a request_id in the response
	if reqID, ok := result["request_id"].(string); ok {
		return reqID, nil
	}

	return "success", nil
}
