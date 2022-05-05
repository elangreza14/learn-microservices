package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://tokopedia.com",
		"http://bukalapak.com",
		"http://instagram.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(lk string) {
			time.Sleep(5 * time.Second)
			checkLink(lk, c)
		}(l)
	}
}

func checkLink(lnk string, c chan string) {
	_, err := http.Get(lnk)
	if err != nil {
		fmt.Println(lnk, "Somethin' error")
		c <- lnk
		return
	}

	fmt.Println(lnk, "ALL IZZ WELL")
	c <- lnk

}
