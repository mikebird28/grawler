package grawler

type Grawler struct {
	workerQueue chan chan interface{}
	pendingTask chan interface{}
	quitChan    chan bool
}

type Config struct{
    MaxGoroutine int
}

func NewGrawler(w Worker,c *Config) *Grawler {
	g := &Grawler{
		workerQueue: make(chan chan interface{}, c.MaxGoroutine),
		pendingTask: make(chan interface{}, 10000),
		quitChan:    make(chan bool),
	}
    for i:=0;i<c.MaxGoroutine;i++{
        nw := newWorker(g,w)
        g.workerQueue <- nw
    }
	return g
}


func (g *Grawler) Run() {
	go func() {
		for {
			select {
            case task := <-g.pendingTask:
                worker := <-g.workerQueue
                worker <- task
			case <-g.quitChan:
				break
			}
		}
	}()
}

func (g *Grawler) Quit() {
	g.quitChan <- true
}

func (g *Grawler) PushTask(i interface{}){
    g.pendingTask <- i
}
