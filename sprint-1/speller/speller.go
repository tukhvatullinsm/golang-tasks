//go:build !solution

package speller

var upToTen = map[int64]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var upToHundred = map[int64]string{
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var upToTrillion = map[int64]string{
	1: "thousand",
	2: "million",
	3: "billion",
	4: "trillion",
}

var upToTwenty = map[int64]string{
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

func Spell(n int64) string {
	scaleCount := int64(0)
	stringTmp := ""
	stringSlice := make([]string, 0)
	if n == 0 {
		return "zero"
	}
	if n < 0 {
		stringTmp += "minus "
		n *= -1
	}

	remainDivHn := func(n int64) string {
		remainSlice := make([]string, 0)
		tempString := ""
		if n >= 100 {
			remainSlice = append(remainSlice, upToTen[n/100]+" hundred")
			n %= 100
		}
		switch {
		case n < 10 && n > 0:

			remainSlice = append(remainSlice, upToTen[n])
		case n <= 19 && n >= 11:

			remainSlice = append(remainSlice, upToTwenty[n])
		case n%10 == 0 && n > 0:
			remainSlice = append(remainSlice, upToHundred[n/10])
		case n%10 != 0 && n > 0:
			remainSlice = append(remainSlice, upToHundred[n/10]+"-"+upToTen[n%10])
		}

		for i, v := range remainSlice {
			if i == len(remainSlice)-1 {
				tempString += v
			} else {
				tempString += v + " "
			}

		}

		return tempString
	}

	for n > 0 {
		prefix, ok := upToTrillion[scaleCount]
		resultDiv := n % 1000
		if ok && resultDiv != 0 {
			stringSlice = append(stringSlice, prefix)
		}
		if resultDiv != 0 {
			stringSlice = append(stringSlice, remainDivHn(resultDiv))
		}
		scaleCount++
		n /= 1000
	}
	for i := len(stringSlice) - 1; i >= 0; i-- {
		if i == 0 {
			stringTmp += stringSlice[i]
		} else {
			stringTmp += stringSlice[i] + " "
		}
	}

	return stringTmp

}
