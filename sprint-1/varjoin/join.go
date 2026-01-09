//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	stringTemp := ""
	for index, arg := range args {
		if index == len(args)-1 {
			stringTemp += arg
		} else {
			stringTemp += arg + sep
		}
	}
	return stringTemp
}
