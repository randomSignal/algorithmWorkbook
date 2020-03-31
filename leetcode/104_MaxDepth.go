package leetcode

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
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var maxDepthInt = 0

func maxDepth(root *TreeNode) int {
    maxDepthInt = 0
    if root == nil {
        return 0
    }
    scanTree(root, 1)
    return maxDepthInt
}

func scanTree(root *TreeNode, depth int) {
    if  depth > maxDepthInt {
        maxDepthInt = depth
    }

    if root.Left != nil {
        scanTree(root.Left, depth + 1) 
    }
    if root.Right != nil {
        scanTree(root.Right, depth + 1) 
    }
}
