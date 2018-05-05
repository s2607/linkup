package main
import (
	"fmt"
	"testing"
)

func TestCheckstr(t *testing.T) {
	go Dbwriter()
	r := new(response)
	c := new(criterion)
	fmt.Println("testing checkstr")
	c.regex="a.*"
	r.value="abc";if(!c.checkstr(r.value)){fmt.Println("checkstr did not match a.* to abc");t.Fail()}
	r.value="bc";if(c.checkstr(r.value)){fmt.Println("checkstr match a.* to bc");t.Fail()}
	r.value="";if(c.checkstr(r.value)){fmt.Println("checkstr match a.* to null string");t.Fail()}
	Killchan <- true
}
