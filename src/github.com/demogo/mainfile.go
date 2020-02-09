package main

import (
	"fmt"

	testlib "github.com/demogo/lib"
	"github.com/demogo/testdir"
)

func main() {
	fmt.Println("value is : " + testlib.GetLang("chirag"))
	testdir.Print()
}
