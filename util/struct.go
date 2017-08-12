package util

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/intelligentpos/structextract"
)

// GetTypes returns map["Field"]="Type"
func GetTypes(obj interface{}) map[string]string {
	d := reflect.ValueOf(obj)
	t := d.Type()
	n := t.NumField()
	types := make(map[string]string)
	for i := 0; i < n; i++ {
		types[t.Field(i).Name] = t.Field(i).Type.String() //io.Sf("%v", t.Field(i).Type)
	}
	return types
}

// GetTags returns map["Field"]="tag_value"
func GetTags(obj interface{}, tag string) map[string]string {
	d := reflect.ValueOf(obj)
	t := d.Type()
	n := t.NumField()
	tags := make(map[string]string)
	for i := 0; i < n; i++ {
		tags[t.Field(i).Name] = t.Field(i).Tag.Get(tag) //io.Sf("%v", t.Field(i).Tag.Get(tag))
	}
	return tags
}

// GetFieldsInfo return information of fields
func GetFieldsInfo(obj interface{},
	tagname string,
	ignoreFields ...string) (fields, types, kinds, tags []string, err error) {
	Sf := fmt.Sprintf

	if err = isValidStruct(obj); err != nil {
		return nil, nil, nil, nil, err
	}

	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		t := v.Type().Field(i)
		if isIgnored(t.Name, ignoreFields...) {
			continue
		}
		fields = append(fields, t.Name)
		types = append(types, Sf("%v", t.Type))
		kinds = append(kinds, Sf("%v", t.Type.Kind()))
		if name, ok := t.Tag.Lookup(tagname); ok {
			tags = append(tags, name)
		}
	}
	return fields, types, kinds, tags, nil
}

func isIgnored(fieldName string, ignoreFields ...string) bool {
	for _, l := range ignoreFields {
		if l == fieldName {
			return true
		}
	}
	return false
}

func isValidStruct(e interface{}) error {
	stVal := reflect.ValueOf(e)
	if stVal.Kind() != reflect.Ptr || stVal.IsNil() {
		return errors.New("struct passed is not valid, a pointer was expected")
	}
	structVal := stVal.Elem()
	if structVal.Kind() != reflect.Struct {
		return errors.New("struct passed is not valid, a pointer to struct was expected")
	}

	return nil
}

// StringValues returns an string array with all the values
func StringValues(e interface{}, ignoreFields ...string) ([]string, error) {
	extract := structextract.New(e).IgnoreField(ignoreFields...)
	values, err := extract.Values()
	if err != nil {
		return nil, err
	}
	var s = make([]string, 0)
	for _, v := range values {
		s = append(s, fmt.Sprintf("%v", v))
	}
	return s, nil
}
