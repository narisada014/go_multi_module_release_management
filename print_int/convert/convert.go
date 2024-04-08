package convert

import "strconv"

// Int returns the decimal reversal of the integer i.
func ConvertStrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
