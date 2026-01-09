//go:build !solution

package speller

func Spell(n int64) string {
	const dl1000 int64 = 1000
	const dl100 int64 = 100
	const dl10 int64 = 10
	var iteration int64 = 0
	tempResult := make(map[int64][]int64)
	var a, b, c int64
	var tmpSlice []int64
	var tmpSliceStr []string
	var m int64
	result := ""
	tempUpTo10 := map[int64]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
	tempUpTo20 := map[int64]string{11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen"}
	tempUpTo100 := map[int64]string{1: "ten", 2: "twenty", 3: "thirty", 4: "forty", 5: "fifty", 6: "sixty", 7: "seventy", 8: "eighty", 9: "ninety"}
	switch {
	case n < 0:
		result += "minus "
		n = n * (-1)
	case n == 0:
		result = "zero"
	}
	for n > 0 {
		_, ok := tempResult[iteration]
		if !ok {
			m = n % dl1000
			n /= dl1000
			a = m / dl100
			b = m % dl100
			switch {
			case b >= 11 && b <= 19:
				c = 0
			case b%dl10 == 0:
				c = 0
			default:
				b = b / dl10
				c = m % 10
			}
			tmpSlice = append(tmpSlice, a, b, c)
			tempResult[iteration] = tmpSlice
			tmpSlice = nil
			iteration++
		} else {
			panic("Something wrong")
		}
	}

	for x := len(tempResult) - 1; x >= 0; x-- {
		tmpSlice = tempResult[int64(x)]
		if tmpSlice[0] != 0 {
			tmpSliceStr = append(tmpSliceStr, tempUpTo10[tmpSlice[0]]+" hundred")
			//result += tempUpTo10[tmpSlice[0]] + " hundred"
		}
		if tmpSlice[1] != 0 {
			switch {
			case tmpSlice[1]%dl10 == 0:
				tmpSliceStr = append(tmpSliceStr, tempUpTo100[tmpSlice[1]/dl10])
				//result += tempUpTo100[tmpSlice[1]/dl10]
			case tmpSlice[1] >= 11 && tmpSlice[1] <= 19:
				tmpSliceStr = append(tmpSliceStr, tempUpTo20[tmpSlice[1]])
				//result += tempUpTo20[tmpSlice[1]]
			default:
				tmpSliceStr = append(tmpSliceStr, tempUpTo100[tmpSlice[1]]+"-"+tempUpTo10[tmpSlice[2]])
				//result += tempUpTo100[tmpSlice[1]] + "-"
			}
		}
		if tmpSlice[2] != 0 && tmpSlice[1] == 0 {
			tmpSliceStr = append(tmpSliceStr, tempUpTo10[tmpSlice[2]])
			//result += tempUpTo10[tmpSlice[2]]
		}

		if tmpSlice[0] != 0 || tmpSlice[1] != 0 || tmpSlice[2] != 0 {
			switch {
			case x == 3:
				tmpSliceStr = append(tmpSliceStr, "billion")
				//result += " billion"
			case x == 2:
				tmpSliceStr = append(tmpSliceStr, "million")
				//result += " million"
			case x == 1:
				tmpSliceStr = append(tmpSliceStr, "thousand")
				//result += " thousand"
			}
		}

	}

	for i, v := range tmpSliceStr {
		if i == len(tmpSliceStr)-1 {
			result += v
			break
		}
		result += v + " "
	}
	return result
}
