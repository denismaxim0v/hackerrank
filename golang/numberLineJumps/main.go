package main

import (
	"bufio"
	"fmt"
	"os"
)

// The first kangaroo starts at location `x1` and moves at a rate of `v1` meters per jump.
// The second kangaroo starts at location `x2` and moves at a rate of `v2` meters per jump.

// Input: 0 3 4 2
// Output: YES

func main() {
	var x1, x2, v1, v2 int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &x1)
	fmt.Fscan(io, &x2)
	fmt.Fscan(io, &v1)
	fmt.Fscan(io, &v2)

	result := kangaroo(x1, x2, v1, v2)
	fmt.Print(result)
}

func kangaroo(x1, x2, v1, v2 int) string {
	if (x2-x1)*(v2-v1) < 0 && (x2-x1)%(v2-v1) == 0 {
		return "YES"
	}
	return "NO"
}
