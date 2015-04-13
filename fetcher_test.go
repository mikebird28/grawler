package grawler

import(
    "testing"
)

func TestFetch(t *testing.T){
    m := map[string]string{
        "this":"that",
        "test":"testing",
    }
    logs(joinMapString(m))
}
