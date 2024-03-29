/*
 * @lc app=leetcode id=9 lang=rust
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (45.62%)
 * Likes:    1727
 * Dislikes: 1422
 * Total Accepted:    728.7K
 * Total Submissions: 1.6M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 * 
 * Example 1:
 * 
 * 
 * Input: 121
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 * 
 * 
 * Follow up:
 * 
 * Coud you solve it without converting the integer to a string?
 * 
 */

pub struct Solution {} 
// @lc code=start
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0{
            return false
        }
        let mut cur: i32 = x;
        let mut res: i32 = 0;
        while cur != 0 {
            match res.checked_mul(10) {
                None => return false,
                Some(tmp) => match tmp.checked_add(cur % 10) {
                    None => return false,
                    Some(fine) => {
                        res = fine;
                    }
                },
            }
            cur = cur / 10;
        }
        return cur == x
    }
}
// @lc code=end


mod tests {
    use super::*;

    #[test]
    fn test_9() {
        assert_eq!(true, Solution::is_palindrome(121));
    }
}