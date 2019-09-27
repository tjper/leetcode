package longestsubstring

import "bytes"

func lengthOfLongestSubstring(s string) int {
	var (
		maxLength int
		buf       = make([]byte, 0, len(s))
	)
	for i := range s {
		if index := bytes.IndexByte(buf, s[i]); index != -1 {
			buf = buf[index+1:]
		}
		buf = append(buf, s[i])
		if l := len(buf); l > maxLength {
			maxLength = l
		}
	}
	return maxLength
}
