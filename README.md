# Wordcount [![Go Report Card](https://goreportcard.com/badge/github.com/mdanzinger/wordcount)](https://goreportcard.com/report/github.com/mdanzinger/wordcount)

wordcount is a useless little library I made to recieve the most frequently used words given any `io.reader`


Examples :
```
var mostFrequentWordsToGet = 10
someText := []byte("The smart fox fox did something somewhere")
someTextReader := bytes.NewReader(someText)

words := wordcount.MostFrequent(someTextReader, mostFrequentWordsToGet)

for _, w := range words {
	 fmt.Printf(w)
	}
```
