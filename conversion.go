package bartender

import "strconv"

func toString(arg interface{}) string {
	switch arg := arg.(type) {
	case string:
		return arg
	case bool:
		return strconv.FormatBool(arg)
	case uint:
		return strconv.Itoa(int(arg))
	case uint8:
		return strconv.Itoa(int(arg))
	case uint32:
		return strconv.Itoa(int(arg))
	case uint64:
		return strconv.Itoa(int(arg))
	case int:
		return strconv.Itoa(arg)
	case int8:
		return strconv.Itoa(int(arg))
	case int32:
		return strconv.Itoa(int(arg))
	case int64:
		return strconv.Itoa(int(arg))
	default:
		return ""
	}
}
