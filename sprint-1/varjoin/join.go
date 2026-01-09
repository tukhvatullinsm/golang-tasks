//go:build !solution

package varjoin

func Join(sep string, args ...string) string {
	out := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			out += args[i]
			break
		}
		out += args[i] + sep
	}
	return out
}
