/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package main

import (
	"flag"
	"fmt"
	"opgolib"
	"os"
	"sort"
)

type OPJay2 struct {
	line_num         int
	fname            string
	resulted_strings *opgolib.StringArray
	idx              int
	lang             string
}

func (jay *OPJay2) make_tags() {
	var (
		e        error
		f        *os.File
		tags_arr []string
	)
	tags_arr = jay.resulted_strings.Array_get_array()[:jay.idx]
	sort.Strings(tags_arr)
	f, e = os.Create("./tags")
	if e != nil {
		fmt.Printf("Can't create file ./tags")
		return
	}

	for _, tag := range tags_arr {
		f.WriteString(tag)
	}
	f.Close()
}

func main() {
	var (
		lang    *string
		dirname string
	)

	lang = flag.String("lang", "golang", "")
	flag.Parse()

	jay := new(OPJay2)
	jay.lang = *lang
	jay.resulted_strings = opgolib.NewStringArray("", 4096)
	jay.idx = 0
	if len(os.Args) < 2 {
		fmt.Printf("One argument (dirname) needed.\n")
		return
	}

	dirname = os.Args[len(os.Args)-1]
	if jay.lang == "golang" {
		jay.start_go(dirname)
	} else if jay.lang == "clang" {
		jay.start_c(dirname)
	} else if jay.lang == "cpplang" {
		jay.start_cpp(dirname)
	}
	jay.make_tags()
}
