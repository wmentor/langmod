package model

import (
	"io"
	"strings"

	"github.com/wmentor/kv"
	"github.com/wmentor/tokens"
)

func Load(in io.Reader) {

	var list []string

	tokens.Process(in, func(t string) {

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
