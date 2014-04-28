package main

func init() {
	n0 := &node{nil, nil, '0'}
	n1 := &node{nil, nil, '1'}
	n2 := &node{nil, nil, '2'}
	n3 := &node{nil, nil, '3'}
	n4 := &node{nil, nil, '4'}
	n5 := &node{nil, nil, '5'}
	n6 := &node{nil, nil, '6'}
	n7 := &node{nil, nil, '7'}
	n8 := &node{nil, nil, '8'}
	n9 := &node{nil, nil, '9'}
	npl := &node{nil, nil, '+'}
	neq := &node{nil, nil, '='}
	nsl := &node{nil, nil, '/'}

	h := &node{n5, n4, 'h'}
	v := &node{nil, n3, 'v'}
	f := &node{nil, nil, 'f'}
	i2 := &node{nil, n2, ' '}
	l := &node{nil, nil, 'l'}
	ipl := &node{npl, nil, ' '}
	p := &node{nil, nil, 'p'}
	j := &node{nil, n1, 'j'}
	b := &node{n6, neq, 'b'}
	x := &node{nsl, nil, 'x'}
	c := &node{nil, nil, 'c'}
	y := &node{nil, nil, 'y'}
	z := &node{n7, nil, 'z'}
	q := &node{nil, nil, 'q'}
	i8 := &node{n8, nil, ' '}
	i90 := &node{n9, n0, ' '}

	s := &node{h, v, 's'}
	u := &node{f, i2, 'u'}
	// _
	r := &node{l, ipl, 'r'}
	// _
	w := &node{p, j, 'w'}
	d := &node{b, x, 'd'}
	k := &node{c, y, 'k'}
	g := &node{z, q, 'g'}
	o := &node{i8, i90, 'o'}

	i := &node{s, u, 'i'}
	a := &node{r, w, 'a'}
	n := &node{d, k, 'n'}
	m := &node{g, o, 'm'}

	e := &node{i, a, 'e'}
	t := &node{n, m, 't'}
	root = &node{e, t, ' '}

	root.init(enc)
}
