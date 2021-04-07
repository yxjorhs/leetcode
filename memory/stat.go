package memory

import (
	"fmt"
	"runtime"
)

// PrintUsed 打印占用的内存
func PrintUsed() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%fM\n", float64(m.Sys)/1024/1024)
}
