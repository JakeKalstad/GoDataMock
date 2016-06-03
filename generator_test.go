package generate

import (
	"fmt"
	"testing"
)

type SubTestData struct {
	Key  int64
	Meta string
}

type Grouper interface {
}

type TestGroup struct {
	GroupKey  int64
	GroupName string
}

type TestData struct {
	Name  string
	Age   int
	Male  bool
	Info  SubTestData
	Group Grouper
}

func TestGen(t *testing.T) {
	gen := &Generator{}
	data := []TestData{TestData{Group: &TestGroup{}}}
	data = gen.Get(data, 10).([]TestData)
	fmt.Printf("%v+", data)
}
