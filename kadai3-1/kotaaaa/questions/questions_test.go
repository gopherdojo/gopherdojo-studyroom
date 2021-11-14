package questions

import (
	"testing"
)

func TestReadWords(t *testing.T) {
	const wordPath string = "../testdata/ielts_words.txt"
	vc := ReadWords(wordPath)
	expected := 4223
	if len(vc.Words) != expected {
		t.Errorf("Word count is illegal. Result: %v, Expected: %v", expected, vc.Words)
	}
	if len(vc.Words) != expected {
		t.Errorf("Word meaning count is illegal.Result: %v, Expected: %v", expected, vc.Meanings)
	}
}

func TestCreateProblem(t *testing.T) {
	vc := Vocab{
		Words:    []string{"apple", "orange"},
		Meanings: []string{"red fruit", "orange fruit"},
	}
	result := CreateProblem(&vc)
	if result != 0 && result != 1 {
		t.Errorf("Target Seed is illegal. Result: %v, Expected: 0 or 1", result)
	}
}
