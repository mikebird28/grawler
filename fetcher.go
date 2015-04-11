package grawler

import (
	"net/http"
    "errors"
)

var URLError = errors.New("Cannot parse the url.")

type Fetcher struct {
}

func (f *Fetcher) Do(i interface{}) {
	url, ok := i.(string)
	if !ok {
        panic(URLError)
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f.Done(resp)
}

func (f *Fetcher) Panic(i interface{}) {
}

func (f *Fetcher) Done(resp *http.Response) {
}
