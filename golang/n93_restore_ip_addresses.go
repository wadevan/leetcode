package golang

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (32.90%)
 * Likes:    851
 * Dislikes: 361
 * Total Accepted:    159.1K
 * Total Submissions: 483.5K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string containing only digits, restore it by returning all possible
 * valid IP address combinations.
 *
 * Example:
 *
 *
 * Input: "25525511135"
 * Output: ["255.255.11.135", "255.255.111.35"]
 *
 *
 */

// @lc code=start

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	data := []byte(s)
	strings := gen(0, data[:])
	return strings
}

func gen(index int, data []byte) []string{

	res := []string{}
	end := len(data)
	for i := 0; i < end; i++ {
		cur := string(data[:i+1])
		if isValid(cur) {
			if index > 3 {
				return []string{}
			}
			if i == end-1 && index == 3 {
				return []string{cur}
			}
			tmp := gen(index+1, data[i+1:])

			for _, item := range tmp {
				ip := fmt.Sprintf("%s.%s", cur, item)
				res = append(res, ip)
			}
		} else {
			return res
		}
	}
	return res
}

func isValid(cur string) bool {

	if len(cur)>1 && strings.HasPrefix(cur, "0"){
		return false
	}

	i, _ := strconv.Atoi(cur)
	if i <= 255 {
		return true
	}
	return false
}

// @lc code=end
