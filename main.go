package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	bot := GetBot("7689397415:AAE7i6UOFSm7PKO7uDDp1iQ6JInwFj4_O7M")

	publicIP, err := GetPublicIP()
	if err != nil {
		fmt.Printf("Errore: %v\n", err)
		return
	}

	IP := ("Il tuo indirizzo IP pubblico è: " + publicIP + "\n")

	chatID := (605277302)

	for {
		if CreateDBtxt(publicIP) {
			sendMessage(bot, int64(chatID), IP)
		}
	}

}

func GetBot(Token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot
}

// Fai richiesta ad api.ipify.org per ottenere l'indirizzo Ip
func GetPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text"

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("errore nella richiesta HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("errore nel leggere la risposta: %v", err)
	}

	return string(body), nil
}

// Manda un messaggio nella chat dove viene mandato un messaggio
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(message)
	if err != nil {
		log.Printf("Errore nell'invio del messaggio: %v", err)
		return
	}
	log.Println("Messaggio inviato con successo!")
}

func CreateDBtxt(publicIP string) bool {

	filename := "ActualIp.txt"

	content, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Errore nella lettura del file: %v", err)
	}

	if publicIP != string(content) {
		fmt.Println("Sono diversi\nip: " + publicIP + "\ncontent: " + string(content))
		err := os.WriteFile(filename, []byte(publicIP), 0644)
		if err != nil {
			log.Fatalf("Errore nella scrittura del file: %v", err)
		}
		return true
	}

	return false
}

func ReadFile(file *os.File) string {

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Errore durante la lettura del file:", err)
	}
	content := string(data)
	return content
}

func WriteFile(file *os.File, publicIP string) {

	_, err := file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Errore durante la posizione nel file: %v", err)
	}

	_, err = file.WriteString(publicIP)
	if err != nil {
		log.Fatalf("Errore nella scrittura del file: %v", err)
	}
}
