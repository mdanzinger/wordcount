package wordcount

import (
	"bytes"
	"io"
	"log"
	"regexp"
	"sort"
	"strings"
)

type Words []Word

type Word struct {
	Word  string
	Count int
}

func (w Words) Len() int {
	return len(w)
}
func (w Words) Less(i, j int) bool {
	return w[i].Count > w[j].Count
}
func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func MostFrequent(r io.Reader, stripStopWords bool) Words {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	// Regex to strip all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z ]+| +")
	if err != nil {
		log.Fatal(err)
	}
	processedContent := reg.ReplaceAll(buf.Bytes(), []byte(" "))
	if stripStopWords {
		processedContent = removeStopWords(processedContent)
	}

	words := bytes.Fields(processedContent)
	m := make(map[string]int)

	for _, word := range words {
		m[string(word)] += 1
	}
	//return m
	w := collectTopWords(m, 100)

	return w
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

	if top >= len(words) {
		return words
	}
	return words[:top]
}
