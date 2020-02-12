package leetcode_test

import (
	"github.com/randomSignal/algorithmWorkbook/leetcode"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum(t *testing.T) {
	args := []int{2, 7, 11, 15}
	res := leetcode.TwoSum(args, 9)
	require.Equal(t, res, []int{0, 1})

	args = []int{3, 2, 4}
	res = leetcode.TwoSum(args, 6)
	require.Equal(t, res, []int{1, 2})

	args = []int{3, 3}
	res = leetcode.TwoSum(args, 6)
	require.Equal(t, res, []int{0, 1})
}
