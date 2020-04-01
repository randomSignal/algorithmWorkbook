package leetcode_test

import (
//"github.com/randomSignal/algorithmWorkbook/leetcode"
//"github.com/stretchr/testify/require"
//"testing"
)

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7

返回它的最大深度 3 。
通过次数149,156
提交次数205,259
*/

/*
func TestMaxDepth(t *testing.T) {
	var res int

	type Sample struct {
		Args   []int
		Result int
	}

	sampleList := []Sample{
        {Args: []int{3,9,20,0,0,15,7}, Result: 3},
	}

	for i := 0; i < len(sampleList); i++ {
		res = leetcode.MaxDepth(sampleList[i].Args)
		require.Equal(t, sampleList[i].Result, res)
	}
}
*/
