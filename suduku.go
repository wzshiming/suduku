package suduku

import (
	"bytes"
	"strings"
)

const (
	size = 9
)

type Grid [size][size]uint8

type Suduku struct {
	grid
}

func NewSuduku() *Suduku {
	return &Suduku{}
}

func (s *Suduku) Import(data Grid) {
	s.grid = grid{}
	for y, line := range data {
		for x, row := range line {
			if row == 0 || row > size {
				continue
			}
			s.Set(uint8(x), uint8(y), row)
		}
	}
}

func (s *Suduku) Export() Grid {
	out := Grid{}
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		out[y][x] = s.subgrid(x, y).clearly
	}
	return out
}

func (s *Suduku) SolveAll() []Grid {
	out := []Grid{}
	s.Solve(func(s *Suduku) bool {
		out = append(out, s.Export())
		return true
	})
	return out
}

func (s *Suduku) SolveFirst() Grid {
	out := Grid{}
	s.Solve(func(s *Suduku) bool {
		out = s.Export()
		return false
	})
	return out
}

func (s *Suduku) Get(x, y uint8) uint8 {
	return s.subgrid(x, y).clearly
}

func (s *Suduku) Set(x, y uint8, num uint8) {
	if !s.subgrid(x, y).IsPossible(num) {
		return
	}
	s.subgrid(x, y).Clearly(num)

	for i := uint8(0); i != size; i++ {
		if i == x {
			continue
		}
		p := s.subgrid(i, y)
		p.Impossible(num)
	}

	for i := uint8(0); i != size; i++ {
		if i == y {
			continue
		}
		p := s.subgrid(x, i)
		p.Impossible(num)
	}

	for j := uint8(0); j != size/3; j++ {
		for i := uint8(0); i != size/3; i++ {
			m := x/3*3 + i
			n := y/3*3 + j
			if m == x || n == y {
				continue
			}
			p := s.subgrid(m, n)
			p.Impossible(num)
		}
	}
}

func (s *Suduku) Clone() *Suduku {
	return &Suduku{
		grid: s.grid,
	}
}

func (s *Suduku) order() ([][2]uint8, uint8, bool) {
	min := uint8(size)
	var order [][2]uint8
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		p := s.subgrid(x, y)
		if p.IsClearly() {
			continue
		}
		si := p.PossibilitySize()
		switch {
		case si == 0:
			return nil, 0, false
		case si > min:
			continue
		case si < min:
			min = si
			order = order[:0]
		}
		order = append(order, [2]uint8{x, y})
	}
	return order, min, true
}

func (s *Suduku) solve(unique map[string]struct{}, fun func(s *Suduku) bool) bool {
	u := s.uniq()
	if _, ok := unique[u]; ok {
		return false
	}
	unique[u] = struct{}{}
	order, min, ok := s.order()
	if !ok {
		return false
	}
	if len(order) == 0 {
		return !fun(s)
	}

	if min == 1 {
		x := order[0][0]
		y := order[0][1]
		point := s.subgrid(x, y)
		s.Set(x, y, point.FirstPossibility())
		return s.solve(unique, fun)
	}
	for _, o := range order {
		x := o[0]
		y := o[1]
		point := s.subgrid(x, y)
		ns := s.Clone()
		ns.Set(x, y, point.FirstPossibility())
		if ns.solve(unique, fun) {
			return true
		}
	}
	return false
}

func (s *Suduku) Solve(fun func(s *Suduku) bool) {
	unique := map[string]struct{}{}
	s.solve(unique, fun)
}

type grid [size * size]subgrid

func (p *grid) subgrid(x, y uint8) *subgrid {
	return &(*p)[x*size+y]
}

func (p *grid) String() string {
	table := [][]string{}
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		if x == 0 {
			table = append(table, []string{})
		}
		table[len(table)-1] = append(table[len(table)-1], p.subgrid(x, y).String())
		if x == 2 || x == 5 {
			table[len(table)-1] = append(table[len(table)-1], "|")
		}
		if x == size-1 && (y == 2 || y == 5) {
			table = append(table, lineSep)
		}
	}
	return strings.Join(formatTable(table), "\n") + "\n"
}

var lineSep = []string{"-", "-", "-", "+", "-", "-", "-", "+", "-", "-", "-"}

func (p *grid) uniq() string {
	var tmp [size * size]byte
	buf := tmp[:0]
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		buf = append(buf, p.subgrid(x, y).clearly+'0')
	}
	return string(buf)
}

type subgrid struct {
	flag    uint64
	clearly uint8
}

func (p *subgrid) IsClearly() bool {
	return p.clearly != 0
}

func (p *subgrid) PossibilitySize() uint8 {
	var s uint8
	for i := uint8(0); i != size; i++ {
		if p.IsPossible(i + 1) {
			s++
		}
	}
	return s
}

func (p *subgrid) FirstPossibility() uint8 {
	for i := uint8(0); i != size; i++ {
		if p.IsPossible(i + 1) {
			return i + 1
		}
	}
	return 0
}

func (p *subgrid) Possibility() []uint8 {
	s := []uint8{}
	for i := uint8(0); i != size; i++ {
		if p.IsPossible(i + 1) {
			s = append(s, i+1)
		}
	}
	return s
}

func (p *subgrid) IsPossible(n uint8) bool {
	return p.flag&(1<<(n-1)) == 0
}

func (p *subgrid) Impossible(n uint8) {
	if p.clearly != 0 {
		return
	}
	p.flag |= 1 << (n - 1)
}

func (p *subgrid) Clearly(n uint8) {
	p.clearly = n
}

func (p *subgrid) String() string {
	if p.clearly != 0 {
		return formatUint(p.clearly)
	}
	s := p.Possibility()
	return "[" + formatRange(s) + "]"
}

func formatRange(s []uint8) string {
	if len(s) == 0 {
		return ""
	}
	tmp := [][]uint8{{s[0]}}
	for _, n := range s[1:] {
		a := tmp[len(tmp)-1]
		if len(a) == 0 || a[len(a)-1]+1 == n {
			tmp[len(tmp)-1] = append(tmp[len(tmp)-1], n)
		} else {
			tmp = append(tmp, []uint8{n})
		}
	}
	ss := []string{}
	for _, o := range tmp {
		if len(o) == 1 {
			ss = append(ss, formatUint(o[0]))
		} else if len(o) == 2 {
			ss = append(ss, formatUint(o[0]), formatUint(o[1]))
		} else {
			ss = append(ss, formatUint(o[0])+"-"+formatUint(o[len(o)-1]))
		}
	}
	return strings.Join(ss, ",")
}

const uintList = ".123456789"

func formatUint(u uint8) string {
	return uintList[u : u+1]
}

func formatTable(b [][]string) (ss []string) {
	max := []int{}
	for _, v1 := range b {
		for k, v2 := range v1 {
			if len(max) == k {
				max = append(max, 0)
			}
			if b := len(v2); max[k] < b {
				max[k] = b
			}
		}
	}
	buf := bytes.NewBuffer(nil)
	for _, v1 := range b {
		buf.Reset()
		for k, v2 := range v1 {
			buf.WriteString(v2)
			ps := max[k] - len(v2) + 1
			for i := 0; i != ps; i++ {
				buf.WriteByte(' ')
			}
		}
		ss = append(ss, buf.String())
	}
	return
}
