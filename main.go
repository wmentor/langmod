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

var (
	loadMode    bool
	debugMode   bool
	limitOutput int
	dataDir     string
	chars       map[rune]bool
)

func init() {

	chars = make(map[rune]bool)

	for _, r := range "1234567890qwertyuiopasdfghjklzxcvbnm-'" {
		chars[r] = true
	}

}

func main() {

	flag.StringVar(&dataDir, "data", "./data", "-data dir")
	flag.BoolVar(&loadMode, "load", false, "load data mode")
	flag.BoolVar(&debugMode, "debug", false, "debug mode")
	flag.IntVar(&limitOutput, "limit-output", 100, "limit output words number")

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

		if t == "'" || t == "-" {
			return
		}

		for _, r := range t {
			if !chars[r] {
				return
			}
		}

		list = append(list, t)
		if len(list) == 4 {
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
		if len(list) == 4 {
			list = list[1:]
		}
	}
}

func loadData() {

	var list []string

	tokens.Process(os.Stdin, func(t string) {

		if t == "'" || t == "-" {
			return
		}

		for _, r := range t {
			if !chars[r] {
				return
			}
		}

		list = append(list, t)
		if len(list) == 4 {
			key := strings.Join(list, " ")
			kv.Set([]byte(key), []byte("1"))
			if debugMode {
				fmt.Println(key)
			}
			list = list[1:]
		}

	})
}
