package roaringcompany

import "strings"

func Index(s string) (startascii uint32, lseascii uint32, meascii uint32, proceed bool) {
	slen := uint32(PureLength(s))
	if slen <= ZERO {
		return ZERO, ZERO, ZERO, FALSE
	}

	start := uint32(s[0])
	end := uint32(s[slen-ONE])
	midpoint := uint32(s[slen/TWO]) // integer division, decimals are truncated

	return start, slen + start + end, midpoint + end, TRUE
}

func PureLength(s string) int {
	return len(strings.TrimSpace(s))
}

func VaildString(s string) bool {
	if len(strings.TrimSpace(s)) > ZERO {
		return TRUE
	}
	return FALSE
}
