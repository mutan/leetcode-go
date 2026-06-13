// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/description/
// Easy

// Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false

// Constraints:
// 1 <= s.length, t.length <= 5 * 10^4
// `s` and `t` consist of lowercase English letters.

// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

package validanagram

// Time complexity: O(n)
// Space complexity: O(1) – for lowercase English letters, O(n) for common case
// Inputs can contain Unicode characters
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, ch := range s {
		counts[ch]++
	}
	for _, ch := range t {
		counts[ch]--
		if counts[ch] < 0 {
			return false
		}
	}

	return true
}

// Time complexity: O(n)
// Space complexity: O(1)
// Only lowercase English letters
func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := make([]int, 26)

	for _, ch := range s {
		chars[ch-'a']++
	}

	for _, ch := range t {
		chars[ch-'a']--
		if chars[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
