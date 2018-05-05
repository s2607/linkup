package main
import (
	"fmt"
	"testing"
)

func TestDelold(t *testing.T) {
	go Dbwriter()
	q := new(question)
	rr := new(responder)
	Init(rr)
	rrkey := rr.key
	Init(q)
	fmt.Println("testing delold")
	Sstore(q); Sstore(rr)
	r := q.New_response(rr)
	rr.responses = append(rr.responses,r)
	badkey := r.key
	r.value = "a"
	Sstore(r)
	r = q.New_response(rr)
	rr.responses = append(rr.responses,r)
	r.value = "b"
	goodkey := r.key
	Sstore(r)
	Sstore(rr)
	rr2 := new(responder)
	rr2.key = rrkey
	Sget(rr2)
	rd := new(response)
	rd.key = badkey
	Sget(rd)
	if rd.key != 0 {
		fmt.Println("old answer persists")
		t.Fail()
	}
	rd.key = goodkey
	Sget(rd)
	if rd == nil {
		fmt.Println("new answer missing")
		t.Fail()
	}
	if rd.value != "b" {
		fmt.Println("new answer blank")
		t.Fail()
	}
	Killchan <-true
}
