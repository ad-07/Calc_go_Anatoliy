package calculation

import (
	"regexp"
	"strconv"
	"strings"
)

func isSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

func Calc(expression string) (float64, error) {
	validExpression := regexp.MustCompile(`^[0-9+\-*/()]+$`)
	if !validExpression.MatchString(expression) {
		return 0, ErrInvalidExpression
	}
	if len(expression) < 3 {
		return 0, ErrInvalidExpression
	}

	var res float64
	var b string
	var c rune = 0
	var resflag bool = false
	var isc int
	var countc int = 0

	for _, value := range expression {
		if isSign(value) {
			countc++
		}
	}
	
	if isSign(rune(expression[0])) || isSign(rune(expression[len(expression)-1])) {
		return 0, ErrInvalidExpression
	}
	for i, value := range expression {
		if value == '(' {
			isc = i
		}
		if value == ')' {
			calc, err := Calc(expression[isc+1 : i])
			if err != nil {
				return 0, ErrInvalidExpression
			}
			calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
			i2 := i
			i -= len(expression[isc:i+1]) - len(calcstr)
			expression = strings.Replace(expression, expression[isc:i2+1], calcstr, 1) // Меняем скобки на результат выражения в них
		}
	}
	if countc > 1 {
		for i := 1; i < len(expression); i++ {
			value := rune(expression[i])
			
			//Умножение и деление
			if value == '*' || value == '/' {
				var imin int = i - 1
				if imin != 0 {
					for !isSign(rune(expression[imin])) && imin > 0 {
						imin--
					}
					imin++
				}
				var imax int = i + 1
				if imax == len(expression) {
					imax--
				} else {
					for !isSign(rune(expression[imax])) && imax < len(expression)-1 {
						imax++
					}
				}
				if imax == len(expression)-1 {
					imax++
				}
				calc, err := Calc(expression[imin:imax])
				if err != nil {
					return 0, ErrInvalidExpression
				}
				calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
				i -= len(expression[isc:i+1]) - len(calcstr) - 1
				expression = strings.Replace(expression, expression[imin:imax], calcstr, 1) // Меняем скобки на результат выражения в них
			}
			if value == '+' || value == '-' || value == '*' || value == '/' {
				c = value
			}
		}
	}
	
	for _, value := range expression + "s" {
		switch {
		case value == ' ':
			continue
		case value > 47 && value < 58: // Если это цифра
			b += string(value)
		case isSign(value) || value == 's': // Если это знак
			if resflag {
				switch c {
				case '+':
					newNum,_ := strconv.ParseFloat(b, 64)
					res += newNum
				case '-':
					newNum,_ := strconv.ParseFloat(b, 64)
					res -= newNum
				case '*':
					newNum,_ := strconv.ParseFloat(b, 64)
					res *= newNum
				case '/':
					if b == "0" {
						return 0, ErrInternalServer
					}
					newNum,_ := strconv.ParseFloat(b, 64)
					res /= newNum
				}
			} else {
				resflag = true
				res,_ = strconv.ParseFloat(b, 64)
			}
			b = strings.ReplaceAll(b, b, "")
			c = value

		case value == 's':
		default:
			return 0, ErrInternalServer
		}
	}
	return res, nil
}
