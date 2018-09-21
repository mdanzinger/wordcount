package wordcount

import (
	"testing"

	"reflect"
	"bytes"
)

func TestMostFrequent(t *testing.T) {
	expects := map[string]int{
		"fox" : 1,
		"did" : 1,
		"something" : 1,
		"somewhere" : 1,
		"The" : 1,
	}
	b := []byte("The fox did something somewhere")
	s := bytes.NewReader(b)
	//f, err := os.Open("content_test.txt")
	//if err != nil {
	//	t.Errorf("Can't open test content")
	//}

	words := MostFrequent(s)

	if !reflect.DeepEqual(words, expects) {
		t.Errorf("Words is nil")
	}
}