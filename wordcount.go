package wordcount

import (
	"bytes"
	"io"
	"log"
	"regexp"
	"sort"
	"strings"
)

// MostFrequent returns a map of the most frequently used words excluding common stopwords
func MostFrequent(r io.Reader, n int) Words {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	// Regex to strip all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z ]+| +")
	if err != nil {
		log.Fatal(err)
	}
	processedContent := removeStopWords(reg.ReplaceAll(buf.Bytes(), []byte(" ")))

	wordCount := collectTopWords(countTopWords(processedContent), n)

	return wordCount

}

// MostFrequentAll returns a map of the most frequently used words including all stopwords.
func MostFrequentAll(r io.Reader, n int) Words {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	// Regex to strip all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z ]+| +")
	if err != nil {
		log.Fatal(err)
	}
	processedContent := reg.ReplaceAll(buf.Bytes(), []byte(" "))

	wordCount := collectTopWords(countTopWords(processedContent), n)

	return wordCount
}

func countTopWords(c []byte) map[string]int {
	words := bytes.Fields(c)
	m := make(map[string]int)

	for _, word := range words {
		m[string(word)] += 1
	}

	return m
}

func removeStopWords(content []byte) []byte {
	var buffer bytes.Buffer
	contentStr := strings.Fields(strings.ToLower(string(content)))

	for _, word := range contentStr {
		if _, ok := stopwords[word]; ok {
			buffer.WriteString("")
		} else {
			buffer.WriteString(word + " ")
		}
	}

	//fmt.Println(buffer.String())
	return buffer.Bytes()
}

func collectTopWords(wordMap map[string]int, top int) Words {
	words := Words{}
	for word, count := range wordMap {
		words = append(words, Word{Word: word, Count: count})
	}
	sort.Sort(words)

	if top >= len(words) || top == 0 {
		return words
	}
	return words[:top]
}
