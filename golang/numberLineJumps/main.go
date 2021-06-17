package main

import (
	"bufio"
	"fmt"
	"os"
)

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
