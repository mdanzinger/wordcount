package wordcount

import (
	"io"
	"strings"
	"bytes"
	"io/ioutil"
	"log"
	"fmt"
)


type Words []Word

type Word struct {
	Word string
	Count int
}

func MostFrequent(r io.Reader) map[string]int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	//bufstring := buf.String()

	bufstring := string(removeStopWords(buf.Bytes()))

	words := strings.Fields(bufstring)
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}


func removeStopWords(content []byte) []byte {
	var result string
	result = strings.ToLower(string(content))
	fmt.Println("Content: " + result)
	stopfile, err := ioutil.ReadFile("stopwords.txt")
	if err != nil {
		log.Printf("Error opening file: %s", err)
	}

	stopwords := strings.Split(string(stopfile), "\n")
	for i, word:= range stopwords{
		//re := regexp.MustCompile(`(?i)\b`+word+`\b`)
		//result = re.ReplaceAllString(result, "")
		if i == 1 {
			result = strings.Replace(string(result), word+" ", " ", -1)
		}
		if i == len(stopwords) - 1 {
			fmt.Println("whaaa")
			result = strings.Replace(string(result), " "+word, " ", -1)
		}
		result = strings.Replace(string(result), " "+word + " ", " ", -1)
	}

	fmt.Println(result)
	return nil
}
