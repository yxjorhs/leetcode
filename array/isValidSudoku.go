package array

const len = 9
const lens = 3

func isValidSudoku(board [][]byte) bool {
	dict := map[byte]bool{}
	// 扫描行
	for i := 0; i < len; i++ {
		dictInit(&dict)
		for j := 0; j < len; j++ {
			if dictSet(&dict, board[i][j]) == false {
				return false
			}
		}
	}

	// 扫描列
	for i := 0; i < len; i++ {
		dictInit(&dict)
		for j := 0; j < len; j++ {
			if dictSet(&dict, board[j][i]) == false {
				return false
			}
		}
	}

	// 扫描九宫格
	for i := 0; i < lens; i++ {
		for j := 0; j < lens; j++ {
			dictInit(&dict)
			for k := 0; k < lens; k++ {
				for l := 0; l < lens; l++ {
					if dictSet(&dict, board[i*3+k][j*3+l]) == false {
						return false
					}
				}
			}
		}
	}

	return true
}

func dictInit(dict *map[byte]bool) {
	*dict = map[byte]bool{}
}

func dictSet(dict *map[byte]bool, val byte) bool {
	if val == '.' {
		return true
	}

	if _, ok := (*dict)[val]; ok {
		return false
	}

	(*dict)[val] = true

	return true
}
