/*
 * (c) Oleg Puchinin 2021
 * puchininolegigorevich@gmail.com
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"opgolib"
)

func (jay *OPJay2) cpp_process_line(s string) string {
	var (
		res string
		spl []string
		idx int
	)
	res = ""
	s = strings.TrimSuffix(s, "\n")
	idx = strings.Index(s, "::")
	if strings.HasPrefix(s, "class ") {
		spl = strings.Split(s, " ")
		res = fmt.Sprintf("%s\t%s\t%d\n", spl[1], jay.fname, jay.line_num-1)
	} else if idx > 0 {
		s = s[idx+2:]
		idx = strings.Index(s, "(")
		s = s[:idx]
		if idx > 0 {
			res = fmt.Sprintf("%s\t%s\t%d\n", s, jay.fname, jay.line_num-1)
		}
	}
	return res
}

func (jay *OPJay2) cpp_process_file(fname string) error {
	var (
		buf         bufio.Reader
		e           error
		f           *os.File
		s           string
		prev_string string
		res         string
	)

	//fmt.Printf("%s\n", fname)
	f, e = os.Open(fname)
	if e != nil {
		return e
	}
	jay.fname = fname
	jay.line_num = 1
	buf = *bufio.NewReader(f)
	s = ""
	for {
		prev_string = s
		s, e = buf.ReadString('\n')
		if e != nil {
			break
		}
		if s == "{\n" {
			res = jay.cpp_process_line(prev_string)
			if res != "" {
				jay.resulted_strings.Array_set_at(jay.idx, res)
				jay.idx++
			}
		}
		jay.line_num++
	}
	f.Close()
	return nil
}

func (jay *OPJay2) start_cpp(dirname string) {
	var (
		files []string
		name  string
	)
	//fmt.Printf("%s\n", "CPP module started !")
	files = opgolib.Find(dirname)
	for _, name = range files {
		if strings.HasSuffix(name, ".cpp") || strings.HasSuffix(name, ".hpp") {
			jay.cpp_process_file(name)
		}
	}
}
