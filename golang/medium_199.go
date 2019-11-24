package golang

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (61.91%)
 * Likes:    110
 * Dislikes: 0
 * Total Accepted:    12.5K
 * Total Submissions: 20.2K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 *
 * 示例:
 *
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 *
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := map[int]int{}
	levels := []int{0}
	q := []*TreeNode{root}
	for len(q) != 0 {
		obj, level := q[0], levels[0]
		q, levels = q[1:], levels[1:]
		res[level] = obj.Val
		if nil != obj.Left {
			q = append(q, obj.Left)
			levels = append(levels, level+1)
		}
		if nil != obj.Right {
			q = append(q, obj.Right)
			levels = append(levels, level+1)
		}
	}
	ret := []int{}
	for i := 0; i < len(res); i++ {
		ret = append(ret, res[i])
	}
	return ret
}

// @lc code=end
