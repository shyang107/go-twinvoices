package util

import (
	"fmt"
	"reflect"
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
	ignoreFields ...string) (fields, types, kinds, tags []string) {
	Sf := fmt.Sprintf
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		if isIgnore(t.Field(i).Name, ignoreFields...) {
			continue
		}
		fields = append(fields, t.Field(i).Name)
		types = append(types, Sf("%v", t.Field(i).Type))
		kinds = append(kinds, Sf("%v", t.Field(i).Type.Kind()))
		tags = append(tags, t.Field(i).Tag.Get(tagname))
	}
	return fields, types, kinds, tags
}

func isIgnore(fieldName string, ignoreFields ...string) bool {
	for i := 0; i < len(ignoreFields); i++ {
		if fieldName == ignoreFields[i] {
			return true
		}
	}
	return false
}
