package regexpress

import (
	"fmt"
	"regexp"
)

func Compile2(str string) (reg *regexp.Regexp, err error) {
	result, err := regexp.Compile(str)
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Caught error: %v", e)
			result = nil // Clear return value.
			//err = string(e) // Will re-panic if not a parse error.
			//panic(e) // re-panic <-- temporary suppress panic
		}
	}()
	if err != nil {
		panic("Incorrect regexp")
	}
	return result, nil
}

func RunRegExp() {
	var validID, _ = Compile2(`^[a-z]+\[[0-9]+\]$`)
	//var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	//
	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))

	var _, err = Compile2(`^[a-+\[[0-9]+\]$`)
	if err != nil {
		fmt.Printf("Invalid re-expression: %v", err)
	}
}
