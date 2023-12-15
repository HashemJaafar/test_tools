package testTools

import (
	"fmt"
	"testing"
)

func Test_Test(t *testing.T) {
	Test(1, 1)
	Test(1, 2)

	PanicIfError = true
	defer fmt.Println("okkkkkkkkkkkkkkkk")
	Test(1, 2)
}

func TestError_New(t *testing.T) {
	err := Error{}.New(1, "kskks")
	err.Print()

	Error{}.Test(0, "")
	err.Test(1, "testTools")
	err.Test(1, "pp")
	err.Test(2, "pp")
}

func TestError_Error(t *testing.T) {
	err := Error{}.New(1, "")
	s := err.GetError()
	fmt.Println(s)
	fmt.Println(s == nil)
}

func TestError_PanicIfError(t *testing.T) {
	err := Error{}
	err.PanicIfError()

	err = Error{}.New(1, "nooooooooooooooo")
	err.PanicIfError()

}
