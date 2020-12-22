package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/wmentor/kv"
	"github.com/wmentor/tokens"
)

const (
	workSize int = 4
)

var (
	loadMode    bool
	limitOutput int
	dataDir     string
)

func init() {
}

func main() {

	flag.StringVar(&dataDir, "data", "./data", "-data dir")
	flag.BoolVar(&loadMode, "load", false, "load data mode")
	flag.IntVar(&limitOutput, "output-size", 100, "limit output words number")

	flag.Parse()

	if _, err := kv.Open("path=" + dataDir + " global=1"); err != nil {
		panic(err)
	}
	defer kv.Close()

	if loadMode {
		loadData()
		return
	}

	generateData()

	fmt.Println("")
}

func generateData() {
	var list []string

	tokens.Process(os.Stdin, func(t string) {

		list = append(list, t)
		if len(list) == workSize {
			list = list[1:]
		}

	})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	output := []string{}

	for i := 0; i < limitOutput; i++ {
		key := strings.Join(list, " ") + " "
		output = output[:0]

		kv.Prefix([]byte(key), func(k, v []byte) bool {
			output = append(output, string(k))
			return true
		})

		if len(output) == 0 {
			break
		}

		val := output[r.Intn(len(output))]
		val = val[len(key):]
		items := strings.Fields(val)
		if len(items) == 0 {
			break
		}

		w := items[0]

		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(w)

		list = append(list, w)
		if len(list) == workSize {
			list = list[1:]
		}
	}
}

func loadData() {

	var list []string

	tokens.Process(os.Stdin, func(t string) {

		if t == "\"" {
			list = nil
			return
		}

		list = append(list, t)
		if len(list) == workSize {
			key := strings.Join(list, " ")
			kv.Set([]byte(key), []byte("1"))
			list = list[1:]
		}

	})
}
