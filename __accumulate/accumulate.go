//
package accumulate

//
func Accumulate(strings []string, fx func(string) string) []string {
	result := make([]string, len(strings))
	for i := 0; i < len(strings); i++ {
		result[i] = fx(strings[i])
	}
	return result
}
