package suduku

import (
	"bytes"
)

func Encode(grid Grid) []byte {
	buf := make([]byte, 0, 2*(size+2)*(size+2))
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		buf = append(buf, uintList[grid[x][y]])
		if x+1 == size {
			buf = append(buf, '\n')
		} else {
			buf = append(buf, ' ')
		}
		if x == 2 || x == 5 {
			buf = append(buf, "| "...)
		}
		if x == size-1 && (y == 2 || y == 5) {
			buf = append(buf, "- - - + - - - + - - -\n"...)
		}
	}
	return buf
}

func Decode(data []byte) (Grid, error) {
	grid := Grid{}
	reader := bytes.NewReader(data)
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		for {
			ch, err := reader.ReadByte()
			if err != nil {
				return grid, err
			}
			if ch >= '0' && ch <= '9' {
				grid[x][y] = ch - '0'
				break
			} else if ch == uintList[0] {
				break
			}
		}
	}
	return grid, nil
}
