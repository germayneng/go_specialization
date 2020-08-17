/**
In this program, since we are performing goroutines on the 2 function, there could be cases where whatIsX is executed
before x is doubled by the double function. This non deterministic result is due to different possible interleaving. The
print statement can either be 4 or 16
*/

package main

import (
	"fmt"
	"time"
)

// create first function here
func double(x *int) {
	*x = *x * *x
}

func whatIsX(x *int) {
	fmt.Println(*x)

}



func main() {
	x := 4

	go double(&x)
	go whatIsX(&x)

	// bad practice but lets add a while to prevent main from finishing
	// good practice to use wait groups
	time.Sleep(100 * time.Millisecond)


}




