package grawler

type Grawler struct {
	workerQueue chan chan interface{}
	pendingTask *InfChan
	quitChan    chan bool
}

type Config struct{
    MaxGoroutine int
}

func NewGrawler(w Worker,c *Config) *Grawler {
	g := &Grawler{
		workerQueue: make(chan chan interface{}, c.MaxGoroutine),
		pendingTask: NewInfChan(),
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
        out := g.pendingTask.Out()
		for {
			select {
            case task := <-out:
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
    in := g.pendingTask.In()
    in <- i
}
