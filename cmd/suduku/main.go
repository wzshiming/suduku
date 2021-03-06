package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/wzshiming/suduku"
)

var input = ""
var output = ""
var size = 1

func init() {
	flag.StringVar(&input, "i", input, "input")
	flag.StringVar(&output, "o", output, "output")
	flag.IntVar(&size, "s", size, "size")
	flag.Parse()
}

func main() {
	var in []byte
	if input != "" {
		f, err := ioutil.ReadFile(input)
		if err != nil {
			log.Fatal(err)
		}
		in = f
	} else {
		f, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		in = f
	}
	data, err := suduku.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	var out = ioutil.Discard
	if output != "" {
		f, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}
		out = f
		defer f.Close()
	}

	s := suduku.NewSuduku(data)
	i := 0
	s.Solve(func(s *suduku.Suduku) bool {
		if out != ioutil.Discard {
			data := suduku.Encode(s.Board())
			out.Write(data)
		}
		i++
		needNext := i != size
		if needNext {
			out.Write([]byte{'\n'})
		}
		return needNext
	})
}
