package model

import (
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/wmentor/kv"
	"github.com/wmentor/tokens"
)

func Generate(in io.Reader) []string {

	var list []string
	var result []string

	tokens.Process(in, func(t string) {

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

		result = append(result, w)

		list = append(list, w)
		if len(list) == workSize {
			list = list[1:]
		}
	}

	return result
}
