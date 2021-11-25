package tinyopt

import (
	"testing"
)

func TestOpt(t *testing.T) {
	s := "          key=val1, flagOn, flagOff=0,key2=val2"

	opt := NewOpt()
	opt.Parse(s)

	flagOn, ok := opt.GetBool("flagOn")

	if !ok {
		t.Error("flagOn should be true")
	}

	if !flagOn {
		t.Error("flagOn should be true")
	}

	v, _ := opt.GetStr("key")
	if v != "val1" {
		t.Error(v)
	}

	vv, _ := opt.GetStr("flagOn")
	if vv != "true" {
		t.Error(vv)
	}
	vvv, _ := opt.GetBool("flagOff")
	if vvv != false {
		t.Error(vvv)
	}

	s = "flag"
	opt.Parse(s)
	f, _ := opt.GetBool("flag")
	if f != true {
		t.Error(f)
	}
}
