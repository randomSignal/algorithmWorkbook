package leetcode

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？


例如，上图是一个7 x 3 的网格。有多少可能的路径？


示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28


提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 10 ^ 9
通过次数85,793提交次数143,826

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if m > n {
		return UniquePaths(n, m)
	}

	var a [][]int
	for i := 0; i < m; i++ {
		list := make([]int, n)
		a = append(a, list)
	}

	for i := 0; i < m; i++ {
		a[i][0] = 1
	}

	for i := 0; i < n; i++ {
		a[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			a[i][j] = a[i-1][j] + a[i][j-1]
		}
	}

	return a[m-1][n-1]
}
