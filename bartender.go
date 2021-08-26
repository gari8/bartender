package bartender

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	chunkSize   = 5000
	queryFormat = `
		INSERT INTO
			%s
		VALUES
			%s`
	headerFormat = `%s (%s)`
	bodyFormat   = `(%s)`
)

type DBScheme struct {
	TableName string        `json:"table_name,omitempty"`
	Type      reflect.Type  `json:"type,omitempty"`
	Value     reflect.Value `json:"value,omitempty"`
}

func NewDBScheme(tableName string) DBScheme {
	return DBScheme{
		TableName: tableName,
	}
}

func (s *DBScheme) Reload(field interface{}) {
	s.Type = reflect.TypeOf(field)
	s.Value = reflect.ValueOf(field)
}

func (s *DBScheme) ReloadTableName(tableName string) {
	s.TableName = tableName
}

func (s *DBScheme) Serve(field interface{}) (string, error) {
	s.Reload(field)
	var header string
	var body string
	if s.Value.Kind() == reflect.Slice {
		var b []string
		length := s.Value.Len()
		for i := 0; i < length; i++ {
			//s.Reload(s.Value.Index(i).Interface())
			fmt.Printf("%+v", s.Value.Index(i).Interface())
			//if i == 0 {
			//	header = strings.Join(s.readHeader(), ",")
			//}
			//b = append(b, s.readBody())
		}
		body = strings.Join(b, ",")
	} else {
		header = fmt.Sprintf(headerFormat, s.TableName, strings.Join(s.readHeader(), ","))
		body = s.readBody()
	}
	return fmt.Sprintf(queryFormat, header, body), nil
}

func (s DBScheme) readHeader() []string {
	var _header []string
	for i := 0; i < s.Type.NumField(); i++ {
		ff := s.Type.Field(i)
		tag := ff.Tag.Get("db")
		_header = append(_header, tag)
	}
	return _header
}

func (s DBScheme) readBody() string {
	return fmt.Sprintf(bodyFormat, strings.Join(s.getValue(), ","))
}

func (s DBScheme) getValue() []string {
	var fields []string
	for i := 0; i < s.Type.NumField(); i++ {
		field := s.Value.FieldByName(s.Type.Field(i).Name).Interface()
		fields = append(fields, toString(field))
	}
	return fields
}

func field2String(field interface{}) (string, string) {
	t := reflect.TypeOf(field)
	v := reflect.ValueOf(field)
	var queryHeaders []string
	var queryValues []string
	queryHeader := `(%s)`
	queryValue := `(%s)`
	for i := 0; i < t.NumField(); i++ {
		queryHeaders = append(queryHeaders, t.Field(i).Tag.Get("db"))
		val := v.FieldByName(t.Field(i).Name).Interface()
		queryValues = append(queryValues, toString(val))
	}
	return fmt.Sprintf(queryHeader, strings.Join(queryHeaders, ",")),
		fmt.Sprintf(queryValue, strings.Join(queryValues, ","))
}
