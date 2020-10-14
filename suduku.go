package suduku

const (
	size = 9
)

type Grid [size][size]uint8

type Suduku struct {
	row, col, box [size]int
	board         Grid
}

func NewSuduku(board Grid) *Suduku {
	m := &Suduku{
		board: board,
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != 0 {
				var idx = i/3*3 + j/3
				var v = 0x1 << board[i][j]
				m.row[i] |= v
				m.col[j] |= v
				m.box[idx] |= v
			}
		}
	}
	return m
}

func (m *Suduku) SolveFirst() Grid {
	out := Grid{}
	m.Solve(func(m *Suduku) bool {
		out = m.Board()
		return false
	})
	return out
}

func (m *Suduku) SolveAll() []Grid {
	out := []Grid{}
	m.Solve(func(m *Suduku) bool {
		out = append(out, m.Board())
		return true
	})
	return out
}

func (m *Suduku) Board() Grid {
	return m.board
}

func (m *Suduku) Solve(next func(*Suduku) bool) {
	m.solve(0, next)
	return
}

func (m *Suduku) solve(step int, next func(*Suduku) bool) bool {
	if step >= size*size {
		if next != nil && next(m) {
			return false
		}
		return true
	}

	var i, j = step / size, step % size

	if m.board[i][j] != 0 {
		return m.solve(step+1, next)
	}

	var idx = i/3*3 + j/3
	for k := 1; k < 10; k++ {
		var v = 0x1 << k
		if m.row[i]&v > 0 || m.col[j]&v > 0 || m.box[idx]&v > 0 {
			continue
		}

		m.board[i][j] = uint8(k)
		m.row[i] |= v
		m.col[j] |= v
		m.box[idx] |= v

		ok := m.solve(step+1, next)

		m.row[i] ^= v
		m.col[j] ^= v
		m.box[idx] ^= v
		m.board[i][j] = 0

		if ok {
			return true
		}
	}
	return false
}
