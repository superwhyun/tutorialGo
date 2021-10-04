package main

import "fmt"

//////////////////////
// Step 1: Basics for structure

type person struct {
	firstname string
	lastname  string
}

//////////////////////
// Step 2: embeded structure

type contactInfo struct {
	email   string
	zipcode int
}

type personInfo struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type personInfo2 struct {
	firstname string
	lastname  string

	contactInfo
	// 아래와 동일한 효과
	// 변수명과 타입명을 동일하게 가져갈 때 같은 의미
	// contactInfo   contactInfo
}

func main() {

	//////////////////////
	// Step 1: Basics for structure
	// 초기값을 넣으면서 선언하려면 모든 내부 값을 다 채워줘야 함
	alex := person{"Alex", "Anderson"}
	// alex := person{firstname: "Alex", lastname: "Anderson"}

	var tom person

	fmt.Println(alex)

	// 값 넣기
	tom.firstname = "tom"
	// tom.lastname = "ripper"
	// %+v를 넣으면, 변수명과 값을 같이 출력해 줌.
	fmt.Printf("%+v", tom)

	//////////////////////
	// Step 2: embeded structure
	// multiline으로 값을 할당할 때는 반드시 모든 라인의 끝에 ','를 넣어야 함.
	tumy := personInfo{
		firstname: "tumi",
		lastname:  "bag",
		contact: contactInfo{
			email:   "obj@home.com",
			zipcode: 12345,
		},
	}
	fmt.Printf("%+v", tumy)

	//////////////////////
	// Step 3: receiver and pointer
	tumy._updateName("handbag") // 반영이 안될거임.

	tumyPtr := &tumy
	tumyPtr.updateName("handbag") // 잘 될거임.
	tumy.print()

	tumy.updateName("footbag") // 어머, 이렇게 해도 된대.
	tumy.print()

	//////////////////////
	// Step 4: slice는 조금 다르다.
	// slice는 내부적으로 array를 선언해서 refer해서 쓰는데
	// 자체적으로 얘가 pointer의 형태로 되어 있다.
	// 그래서 slice는 자연스럽게 pointer로 넘어간다.
	// 이런 것을 reference types라 함
	//  - slices
	//  - maps
	//  - channels
	//  - pointers
	//  - functions

}

//////////////////////
// Step 3: receiver and pointer

func (p personInfo) _updateName(newFirstname string) {
	// 아래처럼 하면 될것 같은데, 반영이 전혀 안 될거임.
	// go는 기본적으로 pass by value 이기 때문에
	p.firstname = newFirstname
}

func (ptr *personInfo) updateName(newFirstname string) {
	// 아래처럼 하면 될것 같은데, 반영이 전혀 안 될거임.
	// go는 기본적으로 pass by value 이기 때문에
	(*ptr).firstname = newFirstname
}

func (p personInfo) print() {
	fmt.Printf("%+v", p)
}
