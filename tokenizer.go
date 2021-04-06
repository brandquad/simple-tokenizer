package simple_tokenizer

import (
	"errors"
	"github.com/kljensen/snowball"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

//var StopWordsList = []string{"http", "https", "brandquad", "atlassian", "спасиб", "secur", "channel", "xlsx", "docx", "imag", "attach", "jira"}

func unique(income []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range income {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func checkStopWords(token string, stopWordsList []string) bool {
	return isValueInList(token, stopWordsList)
}

func TokenizeWord(word string) (string, error) {

	word = strings.TrimSpace(word)
	word = strings.ToLower(word)

	// Check Number
	_, err := strconv.ParseFloat(string(word[0]), 64)
	if err == nil {
		return "", errors.New("is numeric")
	}

	if utf8.RuneCountInString(word) < 3 {
		return "", errors.New("too short")
	}

	var (
		stemmed  string
		language = "english"
	)

	// Check Cyrillic
	if unicode.Is(unicode.Cyrillic, []rune(word)[0]) == true {
		language = "russian"
	}

	stemmed, _ = snowball.Stem(word, language, true)

	return stemmed, nil
}

func Tokenize(text string, stopWordsList []string) (ret []string) {

	var words = strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for _, word := range words {

		stemmed, err := TokenizeWord(word)
		if err != nil {
			continue
		}

		if checkStopWords(stemmed, stopWordsList) == false {
			ret = append(ret, stemmed)
		}
	}
	return unique(ret)
}
