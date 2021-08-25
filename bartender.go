package bartender

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const chunkSize = 5000

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

func ShakeOneCocktail(tableName string, field interface{}) string {
	queryString := `
		INSERT INTO
			%s (%s)
		VALUES
			%s`
	var queryHeaders []string
	var queryValues []string
	t := reflect.TypeOf(field)
	v := reflect.ValueOf(field)
	queryValue := `(%s)`
	for i := 0; i < t.NumField(); i++ {
		queryHeaders = append(queryHeaders, t.Field(i).Tag.Get("db"))
		val := v.FieldByName(t.Field(i).Name).Interface()
		queryValues = append(queryValues, toString(val))
	}
	return fmt.Sprintf(queryString, tableName, strings.Join(queryHeaders, ","), fmt.Sprintf(queryValue, strings.Join(queryValues, ",")))
}

func ShakeCocktails(tableName string, list ...interface{}) string {
	//var queryHeaders []string
	//var queryValues []string
	//queryString := `
	//	INSERT INTO
	//		%s (%s)
	//	VALUES
	//		%s`
	for _, ptr := range list {
		fmt.Printf("%+v", ptr)
		val, _ := tableForPointer(ptr)
		fmt.Printf("+++%+v", val.Field(0))
	}
	//for i, field := range list {
	//	fmt.Printf("******%+v", field)
	//	t := reflect.TypeOf(field)
	//	v := reflect.ValueOf(field)
	//	fmt.Printf("++++%+v", t)
	//	queryValue := `(%s)`
	//	var values []string
	//	for j := 0; j < v.NumField(); j++ {
	//		if i == 0 {
	//			queryHeaders = append(queryHeaders, t.Field(j).Tag.Get("db"))
	//		}
	//		values = append(values, toString(v.FieldByName(t.Field(j).Name).Interface()))
	//	}
	//	queryValues = append(queryValues, fmt.Sprintf(queryValue, strings.Join(values, ",")))
	//	if i >= chunkSize {
	//		return fmt.Sprintf(queryString, tableName, strings.Join(queryHeaders, ","), strings.Join(queryValues, ","))
	//	}
	//}
	//return fmt.Sprintf(queryString, tableName, strings.Join(queryHeaders, ","), strings.Join(queryValues, ","))
	return ""
}

func tableForPointer(ptr interface{}) (reflect.Value, error) {
	ptrv := reflect.ValueOf(ptr)
	if ptrv.Kind() != reflect.Ptr {
		e := fmt.Sprintf("gorp: passed non-pointer: %v (kind=%v)", ptr,
			ptrv.Kind())
		return reflect.Value{}, errors.New(e)
	}
	elem := ptrv.Elem()
	return elem, nil
}