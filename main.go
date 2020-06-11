package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	// make() - map을 초기화 시켜주는 함수
	// make()또는 map끝에 {}를 붙여주지 않으면 map은 nil이 되어버리기 때문에 컴파일러시 찾을 수 없게 됨
	// url과 접속 결과를 저장할 map
	var results = make(map[string]string)
	// var results = map[string]string{}

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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		// 접속 후 결과를 저장
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
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
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
