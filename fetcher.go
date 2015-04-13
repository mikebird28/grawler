package grawler

import (
    "io"
    "strings"
	"errors"
	"net/http"
)

var URLError = errors.New("Cannot parse the url.")

type Fetcher struct {
}

func (f *Fetcher) Do(i interface{}) {
	request, ok := i.(http.Request)
	if !ok {
		panic(URLError)
	}
	client := http.DefaultClient
	resp, err := client.Do(&request)
	if err != nil {
		panic(err)
	}
	f.Done(resp)
}

func (f *Fetcher) Panic(i interface{}) {
}

func (f *Fetcher) Done(resp *http.Response) {
}

func NewGet(url string) (http.Request,error){
    req,err :=  http.NewRequest("GET",url,nil)
    return *req,err
}
func NewPost(url string, values map[string]string)(http.Request,error){
    req,err := http.NewRequest("POST",url,createArgReader(values))
    return *req,err
}

func createArgReader(values map[string]string) io.Reader{
    str := joinMapString(values)
    reader := strings.NewReader(str)
    return reader
}

func joinMapString(m map[string]string) string{
    s := ""
    for k,v := range m{
        s = s+k+"="+v+"&"
    }
    s = s[0:len(s)-1]
    return s
}
