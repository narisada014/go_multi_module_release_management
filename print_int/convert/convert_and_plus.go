package convert

import "strconv"

// Int returns the decimal reversal of the integer i.
func ConvertStrToIntPlus(s string) int {
	i, _ := strconv.Atoi(s)
	i = i + 1
	return i
}
