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

func (jay *OPJay2) go_parse_func(s string) string {
	var (
		i  int
		s2 string
	)
	i = 0
	if s[5] == '(' {
		i = strings.Index(s, ")")
		if i < 0 {
			return ""
		}
		//fmt.Printf("%d ", i)
		i++
		s2 = s[i:]
		i = strings.Index(s2, "(")
		if i < 0 {
			return ""
		}
		//fmt.Printf("%d\n", i)
		s2 = s2[:i]
		s2 = strings.TrimSpace(s2)
	} else {
		i = strings.Index(s, "(")
		if i < 0 {
			return ""
		}
		s2 = s[5:i]
	}
	return fmt.Sprintf("%s\t%s\t%d\n", s2, jay.fname, jay.line_num)
}

func (jay *OPJay2) go_parse_type(s string) string {
	var (
		spl []string
	)

	spl = strings.Split(s, " ")
	if len(spl) < 3 {
		return ""
	}

	return fmt.Sprintf("%s\t%s\t%d\n", spl[1], jay.fname, jay.line_num)
}

func (jay *OPJay2) go_parse_package(s string) string {
	var (
		spl []string
		tag string
	)
	spl = strings.Split(s, " ")
	if len(spl) < 2 {
		return ""
	}
	tag = strings.TrimRight(spl[1], "\n")
	return fmt.Sprintf("%s\t%s\t%d ; p\n", tag, jay.fname, jay.line_num)
}

func (jay *OPJay2) go_process_line(s string) string {
	var (
		s2 string
		i  int
	)

	i = strings.Index(s, " ")
	if i <= 0 {
		return ""
	}
	s2 = s[:i]
	if s2 == "func" {
		return jay.go_parse_func(s)
	} else if s2 == "type" {
		return jay.go_parse_type(s)
	} else if s2 == "package" {
		return jay.go_parse_package(s)
	}
	return ""
}

func (jay *OPJay2) go_process_file() error {
	var (
		f   *os.File
		e   error
		buf *bufio.Reader
		s   string
		s2  string
	)
	f, e = os.Open(jay.fname)
	if e != nil {
		return e
	}
	//println(jay.fname)
	buf = bufio.NewReader(f)
	jay.line_num = 1
	for {
		s, e = buf.ReadString('\n')
		if e != nil {
			break
		}
		s2 = jay.go_process_line(s)
		if s2 != "" {
			jay.resulted_strings.Array_set_at(jay.idx, s2)
			jay.idx++
		}
		jay.line_num += 1
	}
	f.Close()
	return nil
}

func (jay *OPJay2) start_go(dirname string) {
	var (
		arr []string
	)

	arr = opgolib.Find(dirname)
	for _, fname := range arr {
		//	fmt.Printf("%s\n", fname)
		if strings.HasSuffix(fname, ".go") {
			jay.fname = fname
			jay.go_process_file()
		}
	}
}
