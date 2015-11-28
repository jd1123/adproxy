package modules

import "testing"

func TestPrintNewMetaDataStruct(t *testing.T) {
	m := NewMetaStruct("test", "test0", "test1")
	if m.ModuleName != "test" {
		t.Errorf("ModuleName should be \"test\", it is not")
	}
	if m.VersionNumber != "test0" {
		t.Errorf("ModuleName should be \"test0\", it is not")
	}
	if m.Service != "test1" {
		t.Errorf("ModuleName should be \"test1\", it is not")
	}
}
