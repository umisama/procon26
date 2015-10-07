package main

import (
	"flag"
	"fmt"
	"github.com/umisama/procon26/plan"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

const PPROF_DIR = "pprof"

var (
	flagFilePath = flag.String("f", "", "file name")
)

func main() {
	flag.Parse()
	pproff, err := os.Create(generatePathToPprofFile())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer pproff.Close()

	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix())
	pprof.StartCPUProfile(pproff)
	defer pprof.StopCPUProfile()

	lines, err := getLinesFromFile(*flagFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	game, err := NewGameMemo(lines, 1*time.Minute, 300)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	plan := game.Run()
	err = outputFile("output.txt", plan)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func getLinesFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(body), "\n")
	return lines, err
}

func outputFile(path string, p *plan.Plan) error {
	str := p.String()
	err := ioutil.WriteFile(path, []byte(str), 0666)
	if err != nil {
		return err
	}
	return nil
}

func generatePathToPprofFile() string {
	return fmt.Sprintf("%s/%d.dat", PPROF_DIR, time.Now().Unix())
}
