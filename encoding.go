package suduku

import "bytes"

func Encode(grid Grid) []byte {
	buf := make([]byte, 0, 2*size*size)
	for i := uint8(0); i != size*size; i++ {
		x := i % size
		y := i / size
		buf = append(buf, grid[x][y]+'0')
		if x+1 == size {
			buf = append(buf, '\n')
		} else {
			buf = append(buf, ' ')
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
			}
		}
	}
	return grid, nil
}
