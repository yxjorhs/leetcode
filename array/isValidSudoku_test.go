package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSudoku(t *testing.T) {
	data := make([][]byte, 9)

	for i := 0; i < 9; i++ {
		data[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			data[i][j] = '.'
		}
	}

	data[0][0] = '5'
	data[0][1] = '3'
	data[0][4] = '7'

	data[1][0] = '6'
	data[1][3] = '1'
	data[1][4] = '9'
	data[1][5] = '5'

	data[2][1] = '9'
	data[2][2] = '8'
	data[2][7] = '6'

	data[3][0] = '8'
	data[3][4] = '6'
	data[3][8] = '3'

	data[4][0] = '4'
	data[4][3] = '8'
	data[4][5] = '3'
	data[4][8] = '1'

	data[5][0] = '7'
	data[5][4] = '2'
	data[5][8] = '6'

	data[6][1] = '6'
	data[6][6] = '2'
	data[6][7] = '8'

	data[7][3] = '4'
	data[7][4] = '1'
	data[7][5] = '9'
	data[7][8] = '5'

	data[8][4] = '8'
	data[8][7] = '7'
	data[8][8] = '9'

	assert.Equal(t, true, isValidSudoku(data))
}
