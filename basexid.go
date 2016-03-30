// Package basex generates alpha id (alphanumeric id) for big integers.  This
// is particularly useful for shortening URLs.
package basex

import (
	"math/big"
	"strconv"
)

var (
	dictionary = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	zero       = big.NewInt(0)
)

func pow(base, exp int) int {
	var result int = 1
	for i := exp; i > 0; i-- {
		result = result * base
	}
	return result
}

// Encode converts the big integer to alpha id (an alphanumeric id with mixed cases)
func Encode(val string) string {
	var result []byte

	base := big.NewInt(int64(len(dictionary)))

	a := big.NewInt(0)
	b := big.NewInt(0)
	c := big.NewInt(0)
	d := big.NewInt(0)

	remaining := big.NewInt(0)
	remaining.SetString(val, 10)

	for exponent := 1; remaining.Cmp(zero) != 0; {
		a.Exp(base, big.NewInt(int64(exponent)), nil) //16^1 = 16
		b.Mod(remaining, a)                           //119 % 16 = 7 | 112 % 256 = 112
		c.Exp(base, big.NewInt(int64(exponent-1)), nil)
		d.Div(b, c)

		//if d > dictionary.length, we have a problem. but BigInteger doesnt have
		//a greater than method :-(  hope for the best. theoretically, d is always
		//an index of the dictionary!
		strVal := d.String()
		index, _ := strconv.Atoi(strVal)
		result = append(result, dictionary[index])
		remaining = remaining.Sub(remaining, b) //119 - 7 = 112 | 112 - 112 = 0
		exponent += 1
	}

	//need to reverse it, since the start of the list contains the least significant values
	return string(reverse(result))
}

// Encode converts the big integer to alpha id (an alphanumeric id with mixed cases)
func Encodei(val int) string {
	var result []byte

	base := len(dictionary)

	var a, b, c, d int

	remaining := val

	var exponent int = 1
	for remaining != 0 {
		a = pow(base, exponent)
		b = remaining % a //119 % 16 = 7 | 112 % 256 = 112
		c = pow(base, exponent-1)
		d = b / c

		//if d > dictionary.length, we have a problem. but BigInteger doesnt have
		//a greater than method :-(  hope for the best. theoretically, d is always
		//an index of the dictionary!
		result = append(result, dictionary[d])
		remaining = remaining - b //119 - 7 = 112 | 112 - 112 = 0
		exponent += 1
	}

	//need to reverse it, since the start of the list contains the least significant values
	return string(reverse(result))
}

// Decode converts the alpha id to big integer
func Decode(s string) string {
	//reverse it, coz its already reversed!
	chars2 := reverse([]byte(s))

	//for efficiency, make a map
	dictMap := make(map[byte]*big.Int)

	j := 0
	for _, val := range dictionary {
		dictMap[val] = big.NewInt(int64(j))
		j += 1
	}

	bi := big.NewInt(0)
	base := big.NewInt(int64(len(dictionary)))

	exponent := 0
	a := big.NewInt(0)
	b := big.NewInt(0)
	intermed := big.NewInt(0)

	for _, c := range chars2 {
		a = dictMap[c]
		intermed = intermed.Exp(base, big.NewInt(int64(exponent)), nil)
		b.Mul(intermed, a)
		bi.Add(bi, b)
		exponent += 1
	}

	return bi.String()
}

func reverse(bs []byte) []byte {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return bs
}
