package grawler

/*
import(
    "testing"
    "fmt"
)

func TestQueue(t *testing.T){
    infchan := NewInfChan()
    in := infchan.In()
    out := infchan.Out()
    for i := 0;i<2000;i++{
        in <- i
    }
    infchan.Close()
    for i := range out{
        defer func(){
            if e := recover();e != nil{
                fmt.Println(e)
            }
        }()
        v := i.(int)
        fmt.Println(v)
    }
}
*/
