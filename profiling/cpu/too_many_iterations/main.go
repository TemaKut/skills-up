package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	file, err := os.Create("profiling/cpu/too_many_iterations/cpu2.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()

	for i := 0; i < 5000; i++ {
		DoStuff()
	}
}

func DoStuff() {
	DoStuff2()
	DoStuff3()
}

func DoStuff2() {
	for i := 0; i < 20000; i++ {
		continue
	}
}

func DoStuff3() {
	for i := 0; i < 1000000; i++ {
		continue
	}
}
