package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// deck이라는 새로운 변수를 만듬. 걔는 []string 타입임.
type deck []string

// receiver function
//  d ==> 변수명 --> 관례적으로 receiver 인자는 단어 하나(또는 둘)를 씀.
//  deck ==> 변수타입
//  print ==> function name
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// 일반 함수 선언과 함 비교해 보렴
// func fuckCard() string {
// 	return "다이아몬드 5"
// }

// 2중 for loop를 이용해 deck을 만드는 함수
func newDeck() deck {
	cards := deck{} //deck type으로 empty 된 변수 선언
	cardSuits := []string{"스페이드", "다이아몬드", "하트", "클로버"}
	cardValues := []string{"A", "2", "3", "4"}

	// for loop에서 index를 안 쓸거면 '_'를 넣음.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//////////////////////////////////
// Step 7: 함수에서..multiple return을 알아보자.
// 	- 여러개 리턴할 수 있다.

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//////////////////////////////////
// Step 8: 파일저장
//	- receiver로 만들자.
//	- slice를 합쳐보자.

func (d deck) toString() string {
	// deck이 어차피 []string 이라 전혀 문제 없음.
	return strings.Join([]string(d), ",")

}

// return 타입을 'error'로 넣을 수 있음. 왜냐! WriteFile의 return type이 error니까.. 대단한 이유 없다.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//////////////////////////////////
// Step 9: 파일읽기
// []byte ==> string ==> []string ==> deck

func newDeckfromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1) // 프로그램 종료하려면
		return nil // nil을 리턴해 주려면
	} else {
		// str1 := string(bs)
		// str2 := strings.Split(str1, ",")
		// fmt.Println(str2)
		return deck(strings.Split(string(bs), ","))
	}
}

//////////////////////////////////
// Step 10: Random 함수 사용하기

func (d deck) shuffle() {
	// 아래 2줄을 넣어주지 않으면, Intn 함수가 항상 같은 값 넘김
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		//////////////////////////////////
		// Swapping in Go
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
