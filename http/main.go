package main

import ( 
	"net/http"
	"fmt"
	"os"
	"log"
	"io/ioutil"
)


func main() { 
  rsp, err := http.Get("http://google.com")
  if err != nil {
		log.Fatalln(err)
    //fmt.Println("Error: ", err)
    os.Exit(1)
  }

	fmt.Println(rsp.StatusCode)

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	rsp.Body.Close() // 이것은 잘 이해가 안 간다. close 해야 한다는 것은 알겠는데, 뭔가 함수 사용법에서 아귀가 잘 안 맞는 느낌?

	sb  := string(body)
	log.Printf(sb)


}
