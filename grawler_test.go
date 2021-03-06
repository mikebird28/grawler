package grawler

import (
    "net/http"
	"testing"
	"time"
)

type TestWorker struct {
}

func (t *TestWorker) Do(interface{}) {
	logs("TEST OUTPUT")
}

func (t *TestWorker) Panic(err interface{}) {
	logs("PANIC")
}

type TestProc struct{}
func (t *TestProc) Done(resp *http.Response){
    logs("ok")
}
func (t *TestProc) Panic(i interface{}){
    logs("panic")
}
var tp = &TestProc{}

func TestGrawler(t *testing.T) {
	logs("Start")
	conf := &Config{
		MaxGoroutine: 4,
	}
	tworker := NewFetcher(tp)
	g := NewGrawler(tworker, conf)
	go g.Run()
	urls := []string{
		"http://b.hatena.ne.jp/hotentry/it",
		"http://b.hatena.ne.jp/hotentry/game",
		"htp://",
	}
	for _, url := range urls {
		req, _ := NewGet(url)
		g.PushTask(req)
	}
	time.Sleep(1 * time.Second)
}
