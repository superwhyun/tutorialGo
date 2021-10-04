// Testing 만들려면, filename_test.go 를 만들면 됨
// go test 실행

package main

import (
	"os"
	"testing"
)

// t: testing handler

// %v 라는게 있네?
// %d, %s 도 먹긴하는것 같음

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

	if d[0] != "A of 스페이드" {
		t.Errorf("Expected 스페이드 A, but got %v", d[0])
	}

	if d[len(d)-1] != "4 of 클로버" {
		t.Errorf("Expected 클로버 A, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFuke(t *testing.T) {
	os.Remove("_deck_testing.tmp")

	d := newDeck()
	d.saveToFile("_deck_testing.tmp")

	loaded_d := newDeckfromFile("_deck_testing.tmp")
	if len(loaded_d) != 16 {
		t.Errorf("Error of len : %v", len(loaded_d))
	}

	os.Remove("_deck_testing.tmp")

}
