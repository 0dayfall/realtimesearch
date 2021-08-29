package main

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

type Wordsform struct {
	words string
}

func main() {

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := tmpl.Execute(w, nil)
			if err != nil {
				logger.Fatal(err)
			}
			return
		}

		details := Wordsform{
			words: r.FormValue("words"),
		}

		// do something with details
		startServices(details.words)
		rules := returnSearchWords(details.words, "^")
		sentiment := returnSearchWords(details.words, "~")

		err := tmpl.Execute(w, struct {
			Success   bool
			Caret     []string
			Sentiment []string
		}{true, rules, sentiment})
		if err != nil {
			logger.Fatal(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal(err)
	}
}

func startServices(details string) {
	rules, _, _ := parseString(details)

}

func parseString(words string) ([]string, []string, []string) {
	rules := returnSearchWords(words, "^")
	sentiment := returnSearchWords(words, "~")
	return rules, sentiment, nil
	logger.Println(rules)
	logger.Println(sentiment)
}

func returnSearchWords(words string, token string) []string {
	wordList := strings.Split(words, " ")
	tokenString := returnSearchWords_2(wordList, token)
	tokenString = strings.TrimPrefix(tokenString, token)
	return strings.Split(tokenString, token)
}

func returnSearchWords_2(words []string, token string) string {
	var returnValue string
	if len(words) > 1 {
		returnValue = returnValue + returnSearchWords_2(words[1:], token)
	}
	if strings.HasPrefix(words[0], token) {
		return returnValue + words[0]
	}
	return returnValue
}

func registerRules(rules []string) {

	tweet.FilteredStream.AddRule(rules)
}
