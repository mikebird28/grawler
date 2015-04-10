package grawler

import(
    "log"
)

func logs(s string){
    log.Println(s)
}

type Queue []interface{}

func (q Queue) Push(i interface{}){
    q = append(q,i)
}

func (q Queue) Pop()interface{}{
    if len(q) == 0{
        return nil
    }
    out := q[0]
    q = q[1:]
    return out
}
