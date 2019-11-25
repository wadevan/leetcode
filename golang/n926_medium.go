package golang

/*
 * @lc app=leetcode.cn id=926 lang=golang
 *
 * [926] 将字符串翻转到单调递增
 *
 * https://leetcode-cn.com/problems/flip-string-to-monotone-increasing/description/
 *
 * algorithms
 * Medium (42.39%)
 * Likes:    26
 * Dislikes: 0
 * Total Accepted:    1.7K
 * Total Submissions: 3.9K
 * Testcase Example:  '"00110"'
 *
 * 如果一个由 '0' 和 '1' 组成的字符串，是以一些 '0'（可能没有 '0'）后面跟着一些 '1'（也可能没有
 * '1'）的形式组成的，那么该字符串是单调递增的。
 *
 * 我们给出一个由字符 '0' 和 '1' 组成的字符串 S，我们可以将任何 '0' 翻转为 '1' 或者将 '1' 翻转为 '0'。
 *
 * 返回使 S 单调递增的最小翻转次数。
 *
 *
 *
 * 示例 1：
 *
 * 输入："00110"
 * 输出：1
 * 解释：我们翻转最后一位得到 00111.
 *
 *
 * 示例 2：
 *
 * 输入："010110"
 * 输出：2
 * 解释：我们翻转得到 011111，或者是 000111。
 *
 *
 * 示例 3：
 *
 * 输入："00011000"
 * 输出：2
 * 解释：我们翻转得到 00000000。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= S.length <= 20000
 * S 中只包含字符 '0' 和 '1'
 *
 *
 */

// @lc code=start
func minFlipsMonoIncr1(S string) int {
	data := []byte(S)

	dp := [20000]int{}
	numOfOne := 0
	if data[0] == '1' {
		numOfOne = 1
	}

	for i := 1; i < len(data); i++ {
		if data[i] == '1' {
			numOfOne += 1
			dp[i] = min(dp[i-1], numOfOne)
		} else {
			dp[i] = min(dp[i-1]+1, numOfOne)
		}
	}
	return dp[len(data)-1]

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFlipsMonoIncr(S string) int {
	data := []byte(S)
	e := make([]int, len(S))
	cntOfZero, cntOfOne := 0, 0
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] == '0' {
			cntOfZero++
		} else {
			e[i] = cntOfZero
		}
	}
	res := 20000
	if len(data)-cntOfZero < res {
		res = len(data) - cntOfZero
	}
	for i := 0; i < len(data); i++ {
		if data[i] == '0' {
			continue
		}
		res = min(cntOfOne+e[i], res)
		cntOfOne++
	}
	return res
}

// @lc code=end
