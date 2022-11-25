package leetcode

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Example 1:

Input: s = "egg", t = "add"
Output: true
Example 2:

Input: s = "foo", t = "bar"
Output: false
Example 3:

Input: s = "paper", t = "title"
Output: true

https://leetcode.com/problems/isomorphic-strings/
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	return oneWayIsomorphic(s, t) && oneWayIsomorphic(t, s)
}

func oneWayIsomorphic(l string, r string) bool {
	mapping := make(map[uint8]uint8)

	for i := range l {

		lc, rc := l[i], r[i]

		if mapped, ok := mapping[lc]; ok {
			if mapped != rc {
				return false
			}
		} else {
			mapping[lc] = rc
		}
	}
	return true
}
