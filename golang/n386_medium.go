package golang

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 *
 * https://leetcode-cn.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (65.25%)
 * Likes:    53
 * Dislikes: 0
 * Total Accepted:    4.4K
 * Total Submissions: 6.7K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n, 返回从 1 到 n 的字典顺序。
 *
 * 例如，
 *
 * 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
 *
 * 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。
 *
 */

// @lc code=start
func lexicalOrder(n int) []int {
	res := []int{}
	for i := 1; i <= 9; i++ {
		dfs(n, i, &res)
	}
	return res
}

func dfs(N, prefix int, res *[]int) {
	if prefix > N {
		return
	}
	*res = append(*res, prefix)
	for i := 0; i <= 9; i++ {
		tmp := prefix*10 + i
		dfs(N, tmp, res)
	}
}

// @lc code=end
