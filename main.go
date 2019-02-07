package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"
)

type Config struct {
	Token  string `json:"token"`
	ChatID int    `json:"chat_id"`
}

type QueryResuls struct {
	Data []struct {
		ProductID string `json:"productID"`
	} `json:"data"`
}

type BookData struct {
	Title   string `json:"title"`
	Summary string `json:"oneLiner"`
}

const (
	packtpubURL = "https://www.packtpub.com/packt/offers/free-learning"
	queryURL    = "https://services.packtpub.com/free-learning-v1/offers?dateFrom=%s&dateTo=%s"
	bookInfoURL = "https://static.packt-cdn.com/products/%s/summary"
)

var (
	telegramURL    = ""
	convoID        = -1
	executablePath = getExecutablePath()
)

func main() {
	config := parseConfig()
	telegramURL = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.Token)
	convoID = config.ChatID

	query := fmt.Sprintf(queryURL, getDateFrom(), getDateTo())
	queryData := QueryResuls{}
	getJSON(query, &queryData)

	if len(queryData.Data) == 0 {
		str := fmt.Sprintf("Invalid data returned from query: %s", query)
		sendError(str)
		return
	}

	bookQuery := fmt.Sprintf(bookInfoURL, queryData.Data[0].ProductID)
	bookData := BookData{}
	getJSON(bookQuery, &bookData)

	message := fmt.Sprintf("PacktPub book today: <a href=\"%s\">%s</a>\n\n%s", packtpubURL, bookData.Title, bookData.Summary)
	sendToTelegram(message, true)
}

func safeJoin(path1, path2 string) string {
	return filepath.FromSlash(path.Join(path1, path2))
}

func getExecutablePath() string {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

func parseConfig() Config {
	configPath := safeJoin(executablePath, "config.json")
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(err)
	}

	return config
}

func getJSON(url string, model interface{}) {
	res, err := http.Get(url)

	if err != nil {
		processError(err)
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(model)

	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	errString := fmt.Sprintf("ERROR: %s", err)
	sendError(errString)
	panic(errString)
}

func dateFormatted(date time.Time) string {
	year := date.Year()
	month := int(date.Month())
	day := date.Day()

	return fmt.Sprintf("%d-%02d-%02dT00:00:00.000Z", year, month, day)
}

func getDateFrom() string {
	return dateFormatted(time.Now())
}

func getDateTo() string {
	return dateFormatted(time.Now().AddDate(0, 0, 1))
}

func sendError(message string) {
	sendToTelegram(message, false)
}

func sendToTelegram(message string, isHTML bool) {
	query := url.Values{}
	query.Add("chat_id", fmt.Sprint(convoID))
	query.Add("text", message)

	if isHTML {
		query.Add("parse_mode", "HTML")
		query.Add("disable_web_page_preview", "True")
	}

	req, _ := http.NewRequest("GET", telegramURL, nil)
	req.URL.RawQuery = query.Encode()

	http.Get(req.URL.String())
}
