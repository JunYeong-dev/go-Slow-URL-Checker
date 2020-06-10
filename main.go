package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// 접속해 볼 url
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.apple.com/",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

// hit : 인터넷 웹 서버의 파일 1개에 접속하는 것을 뜻함
// hitURL - 해당하는 url에 접속
func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	// StatusCode는 0부터 100, 200, 300까지는 리다이렉션을 해주지만
	// 400부터는 무엇인가 문제가 있기 때문에 에러에 포함
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
