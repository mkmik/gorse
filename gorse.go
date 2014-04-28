package main

import (
	"flag"
	"fmt"
)

var (
	root *node
	enc  = make(map[string]string)
)

type node struct {
	left, right *node
	value       rune
}

// init initalizes the the encoding map from the decoding tree.
func (n *node) init(enc map[string]string) {
	n.walk("", func(d string, n *node) {
		if n.value == ' ' {
			return
		}
		enc[string(n.value)] = d
	})
}

// walk descends recursively in the decoding tree
// and invokes the provided function on each node
// passing the morse code accumulated along the path.
func (n *node) walk(d string, f func(string, *node)) {
	f(d, n)
	if n.left != nil {
		n.left.walk(d+".", f)
	}
	if n.right != nil {
		n.right.walk(d+"-", f)
	}
}

// decode recursively descends the decoding tree
// branching left or right according to the input
// morse code.
func (n *node) decode(s string) (string, error) {
	if n == nil {
		return "", fmt.Errorf("undefined sequence")
	}

	if s == "" {
		return string(n.value), nil
	}

	var next *node
	if s[0] == '.' {
		next = n.left
	} else if s[0] == '-' {
		next = n.right
	} else {
		return "", fmt.Errorf("invalid morse char %q", s[0])
	}

	return next.decode(s[1:])
}

// encode uses the encoding map to emit the morse
// code for each rune of an input string, separated by spaces.
func encode(s string) (string, error) {
	res := ""
	for _, r := range s {
		m, ok := enc[string(r)]
		if !ok {
			m = "?"
		}
		res = fmt.Sprint(res, m, " ")
	}
	return res, nil
}

// isMorse is a dummy euristics used to choose
// encoding or decoding operations depending on the
// first char of the input.
func isMorse(s string) bool {
	return s[0] == '-' || s[0] == '.'
}

func main() {
	flag.Parse()

	for _, a := range flag.Args() {
		f := encode
		if isMorse(a) {
			f = root.decode
		}
		if r, err := f(a); err == nil {
			fmt.Println(r)
		} else {
			fmt.Printf("<%s>\n", err)
		}
	}
}
