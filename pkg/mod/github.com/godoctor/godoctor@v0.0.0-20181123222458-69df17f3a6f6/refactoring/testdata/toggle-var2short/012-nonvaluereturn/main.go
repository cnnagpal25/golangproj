// <<<<< toggle,10,2,10,22,pass
package main

import "fmt"

func f() (int,int,int) {
	return 5,8,3
}
func main() {
	var _, y, _ int = f()
	fmt.Println("Value of y :",y)
}
