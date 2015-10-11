package main

import (
	"flag"
	"fmt"
	"github.com/umisama/procon26/plan"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

const PPROF_DIR = "pprof"

var (
	flagFilePath      = flag.String("f", "", "file name")
	flagServerAddress = flag.String("i", "", "server address")
	flagTeamToken     = flag.String("t", "", "team token")
	flagQuestNo       = flag.Int("q", -1, "quest number")
	flagTimeLimit     = flag.Int("s", 180, "time(sec)")
	flagRepeatCount   = flag.Int("n", 3, "max repeat count")
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

	if *flagFilePath != "" {
		lines, err := getLinesFromFile(*flagFilePath)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		game, err := NewGameMemo(lines, time.Duration(*flagTimeLimit)*time.Second, *flagRepeatCount)
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
	} else if *flagServerAddress != "" && *flagTeamToken != "" && *flagQuestNo != -1 {
		resp, err := http.Get(fmt.Sprintf("http://%s/quest%d.txt?token=%s", *flagServerAddress, *flagQuestNo, *flagTeamToken))
		println(fmt.Sprintf("http://%s/quest%d.txt?token=%s", *flagServerAddress, *flagQuestNo, *flagTeamToken))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		lines := strings.Split(string(body), "\r\n")
		game, err := NewGameMemo(lines, time.Duration(*flagTimeLimit)*time.Second, *flagRepeatCount)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		plan := game.Run()
		resp, err = http.PostForm(fmt.Sprintf("http://%s/answer", *flagServerAddress), url.Values{
			"token":  []string{*flagTeamToken},
			"answer": []string{plan.String()},
		})
		if resp.StatusCode != http.StatusOK {
			fmt.Println("send failed")
		}
		err = outputFile("output.txt", plan)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		flag.Usage()
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
