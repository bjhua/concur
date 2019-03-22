package lib

import (
	"testing"
)

type MyString string

func (s MyString)String()string{
	return (string)(s)
}

func Wrap(s string)MyString{
	return MyString(s)
}

func TestDot(t *testing.T){
	d := NewDot("test")
	d.AppendEdge(Wrap("1"), Wrap("2"))
	d.AppendEdge(Wrap("2"), Wrap("3"))
	d.Layout()

}
