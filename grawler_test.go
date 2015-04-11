package grawler

import (
	"testing"
	"time"
)

type TestWorker struct {
}

func (t *TestWorker) Do(interface{}) {
	logs("TEST OUTPUT")
}

func (t *TestWorker) Panic(err interface{}){
    logs("PANIC")
}

type TestTask struct{}

func TestGrawler(t *testing.T) {
    logs("Start")
	conf := &Config{
		MaxGoroutine: 4,
	}
	tworker := &Fetcher{}
	g := NewGrawler(tworker, conf)
	g.Run()
    urls := []string{
        "http://b.hatena.ne.jp/hotentry/it",
        "http://b.hatena.ne.jp/hotentry/game",
        "htp://",
    }
    for _,url := range urls{
		g.PushTask(url)
	}
	time.Sleep(1 * time.Second)
}
