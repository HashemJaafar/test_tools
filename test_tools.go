package testTools

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

const (
	colorGreen  = "\033[32m"
	colorRed    = "\033[31m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

var (
	failTestNumber uint
	PanicIfError   bool = false
)

func Test[t any](actual, expected t) {
	mainTest(actual, expected)
}

func mainTest[t any](actual, expected t) {
	now := time.Now().Format("2006-01-02 15:04:05")
	stack := strings.Split(string(debug.Stack()), "\n")[8]

	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("%v%#v\n", colorBlue, actual)
		fmt.Printf("%v%#v\n", colorYellow, expected)

		failTestNumber++
		fmt.Println(colorRed, now, "\t", stack, "\t", failTestNumber, colorReset)

		if PanicIfError {
			log.Panic()
		}
	} else {
		fmt.Println(colorGreen, now, "\t", stack, colorReset)
	}
}

type Error struct {
	Index       uint
	PackageName string
	Text        string
	CallStack   string
}

func (Error) New(index uint, text string) Error {
	if index == 0 {
		log.Panicln("You cannot use zero as an error index because it is reserved for the absence of error")
	}

	var standardError Error

	standardError.Index = index
	standardError.Text = text

	stackSlice := strings.Split(string(debug.Stack()), "\n")
	stackSlice[6] = colorRed + stackSlice[6] + colorReset

	standardError.CallStack = strings.Join(stackSlice, "\n")

	for _, r := range stackSlice[5] {
		if r == '.' {
			break
		}
		standardError.PackageName += string(r)
	}

	return standardError
}

func (err Error) Test(index uint, packageName string) {
	mainTest(err.Index, index)
	mainTest(err.PackageName, packageName)
}

func (err Error) NotNil() bool {
	return !reflect.DeepEqual(err, Error{})
}

func (err Error) Print() {
	fmt.Println(colorGreen, "Index\t:", err.Index)
	fmt.Println(colorBlue, "Package:", err.PackageName)
	fmt.Println(colorYellow, "Text\t:", err.Text)
	fmt.Println(colorReset, err.CallStack)
}

func (err Error) GetError() error {
	if err.Text == "" {
		return nil
	}
	return errors.New(err.Text)
}

func (err Error) PanicIfError() {
	if err.NotNil() {
		err.Print()
		log.Panic()
	}
}
