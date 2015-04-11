package grawler

type Worker interface {
	Do(interface{})
    Panic(err interface{})
}

type worker struct {
	grawler *Grawler
	wchan   chan interface{}
	qchan   chan bool
	w       Worker
}

func newWorker(g *Grawler, w Worker) chan interface{} {
	nw := &worker{
		grawler: g,
		wchan:   make(chan interface{}),
		qchan:   make(chan bool),
		w:       w,
	}
	nw.start()
	return nw.wchan
}

func (sw *worker) start() {
	go func() {
		for {
			select {
			case t := <-sw.wchan:
                sw.do(t)
                sw.grawler.workerQueue <-sw.wchan
			case <-sw.qchan:
				break
			}
		}
	}()
}

func (sw *worker) do(i interface{}){
    defer func(){
        if err := recover(); err != nil{
            sw.w.Panic(err)
        }
    }()
    sw.w.Do(i)
}
