package test

import (
	"bytes"
	"testing"

	//"testing"
	"github.com/mdanzinger/wordcount"
	//"reflect"
	//"bytes"
)

func TestMostFrequent(t *testing.T) {
	expects := map[string]int{
		"fox":   2,
		"smart": 1,
	}
	b := []byte("The smart fox fox did something somewhere")
	s := bytes.NewReader(b)
	//f, err := os.Open("content_test.txt")
	//if err != nil {
	//	t.Errorf("Can't open test content")
	//}
	//
	words := wordcount.MostFrequent(s, 5)

	//for _, w := range words {
	//	fmt.Printf("Word: %s %v \n", w.Word, w.Count)
	//}

	for _, w := range words {
		if ew, ok := expects[w.Word]; ok {
			if ew != w.Count {
				t.Errorf("Expected different values")
			}
		}
	}
}
