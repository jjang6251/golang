package main

import (
	// "study/mydict"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// "fmt"
// "log"
// "study/accounts"

// UrlChecker
var errRequestFailed = errors.New("Request failed")

func testUrl() {
	//var results = map[string]string{}
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	results["Gello"] = "Hello"
	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err)
		return errRequestFailed
	}
	return nil
}

// GoRoutine

//순차적 실행
func testSequential() {
	sexyCount("nico")
	sexyCount("flynn")
}

//병렬 실행(GoRoutine)
func testGo() {
	go sexyCount("nico")
	sexyCount("flynn")
}

//병렬 실행 실패
func goFail() {
	go sexyCount("nico")
	go sexyCount("flynn")

	//이런 식으로 작성하면 main에는 실행되는 함수가 없으므로 비동기 처리 즉 GoRoutine이 실행이 되지 않고 바로 종료가 된다.
}
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}


//---------------------------------------------
//Channels
//GoRoutine 간의 데이터를 주고 받기 위해서는 채널을 생성하고 채널을 통해 주고 받는다.
func testChannels () {
	c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	
	//time sleep이 없더라도 Go에서는 채널에서 응답값을 받기 위해 기다린다.
	// result := <-c
	// fmt.Println(result)
	//fmt.Println(<-c) //위와 같음
	//fmt.Println(<-c)

	//사람 2명에 대한 채널 수신 후 한번 더 호출 시 deadlock 발생
	//fmt.Println(<-c)

	// <- 이런 식으로 채널로부터 수신한다는 뜻은 Blocking Operation이라고 할 수 있다.
	for i:=0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}
func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}

//---------------------------------------------
//UrlChecker + GoRoutine

type result struct {
	url string
	status string
}

func testUrl2() {
	//var results = map[string]string{}
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	results["Gello"] = "Hello"
	for _, url := range urls {
		hitUrl2(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <- c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// chan<- 이런식으로 채널의 성격을 send-only인지 아니면 receive-only인지 정할 수 있다.
func hitUrl2(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}

func main() {
	// account := accounts.NewAccount("nico")
	// account.Deposit(10)
	// err := account.Withdraw(11)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())

	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, err := dictionary.Search(word)
	// fmt.Println(hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// dictionary.Search(baseWord)
	// dictionary.Delete(baseWord)
	// word, err := dictionary.Search(baseWord)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(word)

	//testUrl();
	testGo()
}