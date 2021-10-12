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

func (jay *OPJay2) c_process_function(line string) string {
	var (
		i   int
		s   string
		spl []string
	)

	if (line[0] != ' ') && (line[0] != '\t') {
		//	fmt.Printf("%s", line)
		i = strings.Index(line, "(")
		if i >= 0 {
			s = line[:i]
			//fmt.Printf("%s\n", s)
			spl = strings.Split(s, " ")
			if len(spl) > 0 {
				s = spl[len(spl)-1]
				s = fmt.Sprintf("%s\t%s\t%d\n", s, jay.fname, jay.line_num-1)
				return s
			}
		}
	}
	return ""
}

func (jay *OPJay2) c_process_struct(line string) string {
	var (
		spl []string
		s   string
	)
	spl = strings.Split(line, " ")
	if (len(spl) == 3) && (spl[2] == "{\n") {
		s = fmt.Sprintf("%s\t%s\t%d\n", spl[1], jay.fname, jay.line_num)
		//fmt.Printf("%s", s)
	} else {
		return ""
	}
	return s
}

func (jay *OPJay2) c_process_file(fname string) error {
	var (
		arr *opgolib.StringArray
		buf *bufio.Reader
		e   error
		f   *os.File
		s   string
		tag string
		//prev string
		idx int
	)
	f, e = os.Open(fname)
	if e != nil {
		return e
	}
	jay.fname = fname
	idx = 0
	arr = opgolib.NewStringArray("", 8192)
	buf = bufio.NewReader(f)
	for {
		jay.line_num = idx + 1
		s, e = buf.ReadString('\n')
		arr.Array_set_at(idx, s)
		if e != nil {
			break
		}

		if strings.HasPrefix(s, "struct ") {
			tag = jay.c_process_struct(s)
			if tag != "" {
				jay.resulted_strings.Array_set_at(jay.idx, tag)
				jay.idx++
			}
		} else if strings.HasPrefix(s, "{") {
			if idx > 0 {
				tag = jay.c_process_function(arr.Array_get_at(idx - 1))
				if tag != "" {
					jay.resulted_strings.Array_set_at(jay.idx, tag)
					jay.idx++
				}
			}
		}
		idx++
	}
	f.Close()
	return nil
}

func (jay *OPJay2) start_c(dirname string) {
	var (
		files []string
		fname string
	)

	files = opgolib.Find(dirname)
	for _, fname = range files {
		if strings.HasSuffix(fname, ".c") ||
			(strings.HasSuffix(fname, ".h")) {
			jay.c_process_file(fname)
		}
	}

}
