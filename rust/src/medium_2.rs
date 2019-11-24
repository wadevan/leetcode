/*
 * @lc app=leetcode.cn id=2 lang=rust
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (36.03%)
 * Likes:    3425
 * Dislikes: 0
 * Total Accepted:    258.2K
 * Total Submissions: 716.7K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 * 
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 * 
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * 
 * 示例：
 * 
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 * 
 * 
 */

// @lc code=start
// Definition for singly-linked list.
// #[derive(PartialEq, Eq, Clone, Debug)]
// pub struct ListNode {
//   pub val: i32,
//   pub next: Option<Box<ListNode>>
// }
// 
// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }
impl Solution {
    pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        Self::add_lists(&l1, &l2, 0)
    }

    fn add_lists(l1: &Option<Box<ListNode>>, l2: &Option<Box<ListNode>>, carry: i32) -> Option<Box<ListNode>> {
        match (l1, l2, carry) {
            (None, None, 0) => None,
            (None, None, c) => Some(Box::new(ListNode::new(c))),
            (Some(l), None, _) | (None, Some(l), _) => {
                let sum = l.val+carry;
                Some(Box::new(ListNode{val: sum%10, next: Self::add_lists(&l.next,&None,sum/10)}))},
            (Some(l1), Some(l2), _) => {
                let sum = l1.val+l2.val+carry;
                Some(Box::new(ListNode{val: sum%10, next: Self::add_lists(&l1.next,&l2.next,sum/10)}))}
        }
    }
}
// @lc code=end

