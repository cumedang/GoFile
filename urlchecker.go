package urlchecker

import (
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func Urlcheck(term string) map[string]string {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		term,
	}
	for _, url := range urls {
		go hitURL(url, c)

	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	// 대충 스트링 변수
	a := make(map[string]string)
	for url, status := range results {
		a = map[string]string{url: status}
	}
	return a
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
