package main

import (
	"fmt"
)

/* 기본
  - Go는 Object Oriented 언어가 아님.
    - 근데 흉내낼 수 있음?

/****************
  GO의 기본 Type
   - bool
   - string
   - int
   - float64
   - array (incl. slice)
   - map

  근디 type 키워드를 이용해 새로운 변수를 만들 수 있음.


*****************/

/*****************
* 배열(Array)
  - Fixed length
* Slice
  - 일종의 배열
  - Grow/shrink 가능
  - 안에 있는 아이템들은 모두 같은 타입을 가져야 함.
******************/

func main() {

	//////////////////////////////////
	// Step 1
	//////////////////////////////////
	fmt.Println("Hello Fucker")

	//////////////////////////////////
	// Step 2
	//////////////////////////////////
	card := "스페이드 A" // var card string = "스페이드 A"
	// 위에 있는 거랑 같은 효과. 단, 컴파일러가 알아서 변수를 할당하는 것에 차이가 있음.
	// 단, 처음에 선언할 떄만 :을 넣어주고, 이후에 사용할때는 일반 변수와 똑같이 사용함.

	fmt.Println(card)

	//////////////////////////////////
	// Step 3
	//////////////////////////////////
	new_card := newCard()
	fmt.Println(new_card)

	//////////////////////////////////
	// Step 4
	//////////////////////////////////

	/***********************
	* slice
	  - slice의 선언은 "[]type{}"을 이용
	  - 초기값을 넣으려면 {} 에 늘어 놓음.
	  - slice에 값을 넣으려면 append 함수를 이용
	  - 일반 배열처럼 slice[0] 등으로 접근 가능
	***********************/

	cards := []string{newCard(), newCard()}

	//
	// 그런데, 조금 방식이 평소와는 다르군..
	cards = append(cards, "fucker")

	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}

	//////////////////////////////////
	// Step 5
	//////////////////////////////////
	// 새로운 type을 선언하고, receiver를 이용해 보자.
	deck_cards := deck{newCard(), newCard()}
	for i, card := range deck_cards {
		fmt.Println(i, card)
	}
	deck_cards.print()

	//////////////////////////////////
	// Step 6: slice를 쪼개는 것을 해 보자
	// - card[0:2] ==> 0, 1 번째 것만 끊어옴. 앞이 0일 경우 생략가능
	// - card[2:] ==> 2번째부터 끝까지
	//////////////////////////////////
	fmt.Println("Step 6 ============================")
	deck := newDeck()
	deck.print()

	//////////////////////////////////
	// Step 7: 함수에서 multiple return을 알아보자.
	fmt.Println("Step 7 ============================")
	hand, remainingDeck := deal(deck, 5)
	hand.print()
	remainingDeck.print()

	//////////////////////////////////
	// Step 8: 파일저장
	// 	- Type conversion (string/.../... ==> []byte )
	//	- string str을 conversion 하려면,
	// 		-[]byte(str)
	// deck --> []string --> string ---> []byte 의 단계를 거쳐야 함.
	fmt.Println("Step 8 ============================")
	fmt.Println(deck.toString())
	deck.saveToFile("fucker.sav")

	//////////////////////////////////
	// Step 9: 파일읽기
	fmt.Println("Step 9 ============================")
	deck2 := newDeckfromFile("fucker.sav")
	deck2.print()

	//////////////////////////////////
	// Step 10: Random 함수 사용하기
	fmt.Println("Step 10: Shuffle ============================")
	deck2.shuffle()
	deck2.print()
}

// 함수 선언
func newCard() string {
	return "다이아몬드 5"
}
