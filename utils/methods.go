package utils

import (
	"errors"
	"strconv"
)

func ParseToString(data interface{}) (string, error) {
	switch d := data.(type) {
	case string:
		return d, nil
	case int:
		return strconv.Itoa(d), nil
	case int64:
		return strconv.FormatInt(d, 10), nil
	case float64:
		return strconv.FormatFloat(d, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(d), 'f', -1, 64), nil
	default:
		return "", errors.New("unsupported type")
	}
}

func ParseToInt(data interface{}) (int, error) {
	switch d := data.(type) {
	case string:
		i, err := strconv.Atoi(d)
		if err != nil {
			return 0, err
		}
		return i, nil
	case int:
		return d, nil
	case int64:
		return int(d), nil
	case float64:
		return int(d), nil
	case float32:
		return int(d), nil
	default:
		return 0, errors.New("unsupported type")
	}
}

func ParseToFloat(data interface{}) (float64, error) {
	switch d := data.(type) {
	case string:
		f, err := strconv.ParseFloat(d, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	case int:
		return float64(d), nil
	case int64:
		return float64(d), nil
	case float32:
		return float64(d), nil
	default:
		return 0, errors.New("unsupported type")
	}
}
