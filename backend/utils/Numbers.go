package utils

import (
	"backendtask/config"
	"errors"
	"strconv"
	"strings"
)

func ScaleFloatToInt(num float64) (int64, error) {

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

	result, err := strconv.ParseInt(int_str, 10, 64)
	if err != nil {
		return 0, errors.New("Failed to parse float!")
	}

	return result, nil
}

func ScaleIntToFloat(num int64) (float64, error) {

	int_str := strings.Repeat("0", config.DecimalPoints) + strconv.FormatInt(num, 10)

	comma_point := len(int_str)-config.DecimalPoints

	float_str := int_str[:comma_point] + "." + int_str[comma_point:] 

	return strconv.ParseFloat(float_str, 64)
}