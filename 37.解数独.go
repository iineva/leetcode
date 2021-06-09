/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start

type solveSudokuer struct {
	board [][]byte
	// key = row * 1000 + column
	maybe map[int][]byte
}

// 获取可能性索引
func (s *solveSudokuer) get(row, column int) []byte {
	return s.maybe[row*1000+column]
}

func (s *solveSudokuer) add(row, column int, l []byte) {
	s.maybe[row*1000+column] = l
}

func (s *solveSudokuer) del(row, column int) {
	delete(s.maybe, row*1000+column)
}

func (s *solveSudokuer) Print() {
	for i, row := range s.board {
		fmt.Print("\n")
		for j, v := range row {
			fmt.Printf("%s  ", string(v))
			if j%3 == 2 {
				fmt.Print("   ")
			}
		}
		if i%3 == 2 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}

func del(l []byte, s byte) []byte {
	ll := []byte{}
	for _, v := range l {
		if v != s {
			ll = append(ll, v)
		}
	}
	return ll
}

// 推理填入一定能确定的数字
func (s *solveSudokuer) opt() {
	loop := 9 * 9

	defer func() {
		log.Print(loop, " ", s.maybe)
	}()

	// 通过逻辑填空
	for {
		match := false

		for i, rowList := range s.board {
			for j, v := range rowList {
				if v != '.' {
					continue
				}
				b := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

				// 排除同3x3
				row := i / 3 * 3
				column := j / 3 * 3
				for ii := 0; ii < 3; ii++ {
					for jj := 0; jj < 3; jj++ {
						vv := s.board[row+ii][column+jj]
						if vv == '.' {
							continue
						}
						b = del(b, vv)
					}
				}
				if len(b) == 1 {
					// log.Printf("%v: %v - %v : %v", loop, i, j, string(b))
					s.board[i][j] = b[0]
					s.del(i, j)
					match = true
					continue
				}

				for ii := 0; ii < 9; ii++ {
					// 排除row
					if ii != i {
						vv := s.board[ii][j]
						if vv != '.' {
							b = del(b, vv)
						}
					}
					// 排除column
					if ii != j {
						vv := s.board[i][ii]
						if vv != '.' {
							b = del(b, vv)
						}
					}
				}
				if len(b) == 1 {
					// log.Printf("%v: %v - %v : %v", loop, i, j, string(b))
					s.board[i][j] = b[0]
					s.del(i, j)
					match = true
					continue
				}

				for _, vv := range b {
					n := 0
					for ii := 0; ii < 3; ii++ {
						if row+ii != i {
							for k := 0; k < 9; k++ {
								if vv == s.board[row+ii][k] {
									n++
								}
							}
						}
						if column+ii != j {
							for k := 0; k < 9; k++ {
								if vv == s.board[k][column+ii] {
									n++
								}
							}
						}
					}
					if n == 4 {
						s.board[i][j] = vv
						s.del(i, j)
						match = true
						continue
					}
				}

				if len(b) == 0 {
					log.Printf("%v: %v - %v : %v", loop, i, j, string(b))
				}

				// TODO 继续优化
				s.add(i, j, b)

			}
		}

		loop--
		if loop <= 0 || len(s.maybe) == 0 || !match {
			break
		}
	}

}

func (s *solveSudokuer) Solve() [][]byte {

	s.opt()

	if len(s.maybe) == 0 {
		return s.board
	}

	row := [9][9]bool{}
	column := [9][9]bool{}
	block := [3][3][9]bool{}
	space := [][2]int{}

	for i, r := range s.board {
		for j, v := range r {
			if v == '.' {
				space = append(space, [2]int{i, j})
			} else {
				d := v - '1'
				row[i][d] = true
				column[j][d] = true
				block[i/3][j/3][d] = true
			}
		}
	}

	// 暴力解法
	var def func(off int) bool
	def = func(off int) bool {

		if off >= len(space) {
			return true
		}

		for k := 0; k < 9; k++ {
			i := space[off][0]
			j := space[off][1]
			v := byte(k + '1')

			finded := false
			if l, ok := s.maybe[i*1000+j]; ok {
				for _, vv := range l {
					if vv == v {
						finded = true
						break
					}
				}
			}
			if !finded {
				continue
			}

			if !row[i][k] && !column[j][k] && !block[i/3][j/3][k] {
				row[i][k] = true
				column[j][k] = true
				block[i/3][j/3][k] = true
				s.board[i][j] = v
				if def(off + 1) {
					return true
				}
				row[i][k] = false
				column[j][k] = false
				block[i/3][j/3][k] = false
			}
		}

		return false
	}

	def(0)

	return s.board
}

func solveSudoku(board [][]byte) {
	s := solveSudokuer{board: board, maybe: map[int][]byte{}}
	s.Print()
	s.Solve()
	s.Print()
}

// @lc code=end

