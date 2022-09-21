package util

import (
	"bytes"
)

func MaskingString(startLen int, endLen int, stringToMask string) string {
	var total = len(stringToMask)
	var masklen = total - (startLen + endLen)
	var maskedBuff bytes.Buffer

	maskedBuff.WriteString(stringToMask[:startLen])
	for i := 0; i < masklen; i++ {
		maskedBuff.WriteString("x")
	}
	maskedBuff.WriteString(stringToMask[masklen+endLen:])

	return maskedBuff.String()
}
