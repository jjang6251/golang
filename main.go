package main

import (
	// "study/mydict"
	"errors"
	"fmt"
	"net/http"
)

// "fmt"
// "log"
// "study/accounts"

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

	testUrl();
}