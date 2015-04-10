package grawler

const (
	bufferSize = 128
)

//InfChan provides the channel which has infinite buffer
type InfChan struct {
	q        queue
	inchan   chan interface{}
	outchan  chan interface{}
	quitchan chan bool
}

//NewInfChan returns a new InfChan which is already initialized
func NewInfChan() *InfChan {
	out := &InfChan{
		q:        newQueue(),
		inchan:   make(chan interface{}, bufferSize),
		outchan:  make(chan interface{}, bufferSize),
		quitchan: make(chan bool),
	}
	go func() {
		for {
			i := out.q.top()
			if i == nil {
				i, ok := <-out.inchan
				if ok {
					out.q.push(i)
				} else {
					break
				}
			} else {
				select {
				case v, ok := <-out.inchan:
					if ok {
						out.q.push(v)
					} else {
						break
					}
				case out.outchan <- i:
					out.q.pop()
				}
			}
		}
		for v := out.q.pop(); v != nil; v = out.q.pop() {
			out.outchan <- v
		}
		close(out.outchan)
	}()
	return out
}

//Close close InfChannel
func (c *InfChan) Close() {
	close(c.inchan)
}

//In returns a channel for input
func (c *InfChan) In() chan<- interface{} {
	return c.inchan
}

//Out returns a channel for output
func (c *InfChan) Out() <-chan interface{} {
	return c.outchan
}

type queue []interface{}

func newQueue() queue {
	return make([]interface{}, 0)
}

func (q *queue) push(i interface{}) {
	*q = append(*q, i)
}

func (q *queue) pop() interface{} {
	if len(*q) == 0 {
		return nil
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *queue) top() interface{} {
	if len(*q) == 0 {
		return nil
	}
	return (*q)[0]
}
