package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stlpid.io",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://discomco.pl",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkLink(site, c)
	}

	for l := range c {
		go func(lnk string) {
			time.Sleep(5 * time.Second)
			checkLink(lnk, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!! Error: ", err)
		c <- link
		return
	}
	fmt.Println(link, "is Up!")
	c <- link
}
