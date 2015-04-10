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

type TestTask struct{}

func TestGrawler(t *testing.T) {
	println("Start")
	conf := &Config{
		MaxGoroutine: 4,
	}
	tworker := &TestWorker{}
	g := NewGrawler(tworker, conf)
	g.Run()
	for i := 0; i < 10; i++ {
		g.PushTask(TestTask{})
	}
	time.Sleep(1 * time.Second)
}
