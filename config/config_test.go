package config

import "testing"

func TestConfigStruct(t *testing.T) {
	c := DefaultConfig()
	if c == nil {
		t.Errorf("DefaultConfig() is returning a nil object.")
	}
	if c.Quiet != true {
		t.Errorf("Quiet should be true, it is not")
	}
	if c.ListenPort != "9000" {
		t.Errorf("ListenPort should be 9000, it is not")
	}
	if c.Verb != false {
		t.Errorf("Verb should be false, it is not")
	}
}
