package string

import "fmt"

func isString(a interface{}) (bool, error) {
	switch a.(type) {
	case string:
		return true, nil
	case int32:
		return false, fmt.Errorf("expected string got int: %d", a)
	default:
		return false, fmt.Errorf("expected string got: %T", a)
	}
}
