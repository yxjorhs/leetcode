package main

import (
	"fmt"
)

func main() {
	for {
		v := 0
		sc, e := fmt.Scan(&v)

		println("scan", v, sc, e)

		if sc == 0 {
			break
		}
	}
	return
}
