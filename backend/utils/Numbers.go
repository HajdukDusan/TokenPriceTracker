package utils

import (
	"backendtask/config"
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// Safely scales a float64 to big.Int value using strings.
// The scaling factor of decimal places is determined in the config.
func ScaleFloatToBigInt(num float64) (*big.Int, error) {

	float_str := strconv.FormatFloat(num, 'f', config.DecimalPoints+1, 64)

	comma_pos := -1

	for i := 0; i < len(float_str); i++ {
		if float_str[i] == '.' {
			comma_pos = i
			break
		}
	}

	var int_str string

	// remove the comma
	if comma_pos != -1 {
		int_str = float_str[:comma_pos] + float_str[comma_pos+1:len(float_str)-1]
	} else {
		int_str = float_str
	}

	result := new(big.Int)
	result, ok := result.SetString(int_str, 10)
	if !ok {
        return nil, errors.New("Failed to parse float!")
    }

	return result, nil
}

// Returns a float value from scaled big.Int.
// The scaling factor of decimal places is determined in the config.
func ScaleBigIntToFloat(num *big.Int) (float64, error) {

	int_str := strings.Repeat("0", config.DecimalPoints) + num.String()

	comma_point := len(int_str)-config.DecimalPoints

	float_str := int_str[:comma_point] + "." + int_str[comma_point:] 

	return strconv.ParseFloat(float_str, 64)
}

// Return an abs value of big.int.
func GetAbs(num *big.Int) *big.Int {
	if num.Cmp(big.NewInt(0)) == -1 {
		num.Neg(num)
	}
	return num
}