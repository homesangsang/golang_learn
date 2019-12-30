package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// 实现Retriever接口的Get方法
func (r Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(response, true)

	response.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}

func (r Retriever) Get1(url string, queryStr string) string {
	response, err := http.Get(url + "?" + queryStr)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(response, true)
	response.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}
