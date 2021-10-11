package main

import "fmt"

func main() {

	/////////////////
	//  Step 1
	//  variable := map[Key타입]Value타입
	colors := map[string]string{
		"red":   "0xff0000",
		"green": "0x00ff00",
		"white": "0xffffff",
	}
	fmt.Println(colors)

	/////////////////
	// Step 2
	// - 추가, 삭제
	// var colors2 map[string]string
	//   ==> 여기까지만 쓰게 되면 nil Map으로 데이터를 쓸 수 없음.
	//		 make를 해 줘야 함.

	colors2 := make(map[string]string)
	colors2["white"] = "0xffffff"
	fmt.Println(colors2)
	delete(colors2, "white")

	/////////////////
	//  Step 3
	//	- Map iteration
	printMap(colors)

}

/////////////////
//  Step 3
//	- Map iteration
func printMap(c map[string]string) {
	for color, hex := range c {
		// fmt.Println("Color :", color, "Hex :", hex)
		fmt.Printf("%-10s : %s\n", color, hex)

	}
}
