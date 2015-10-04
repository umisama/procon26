package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const PPROF_DIR = "pprof"

func main() {
	f, err := os.Create(generatePathToPprofFile())
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	game, err := NewGameMemo(os.Args[1], 1*time.Minute, 300)
	if err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	plan := game.Run()

	fmt.Println(plan)
	fmt.Println(plan.Score())
}

func generatePathToPprofFile() string {
	return fmt.Sprintf("%s/%d.dat", PPROF_DIR, time.Now().Unix())
}
