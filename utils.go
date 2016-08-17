package main

import (
	"strings"
	"unicode"
)

func LowerInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func UpperInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func TypeNoNamespace(str string) string {
	str = strings.TrimPrefix(str, "Phaser.")
	str = strings.TrimPrefix(str, "PIXI.")
	/*lastPoint := strings.LastIndex(str, ".")
	if lastPoint != -1 {
		return str[lastPoint+1 : len(str)]
	}*/
	str = strings.Replace(str, ".", "", -1)
	str = UpperInitial(str)
	return str
}

func SplitMultilines(str string) []string {
	lines := strings.Replace(str, "\r", "\n", -1)
	return strings.Split(lines, "\n")
}

func GoNativeType(str string) string {
	switch str {
	case "boolean":
		return "bool"
	case "Boolean":
		return "bool"
	case "integer":
		return "int"
	case "Integer":
		return "int"
	case "number":
		return "float64"
	case "Number":
		return "float64"
	case "Bounds-like":
		return "float64"
	case "string":
		return "string"
	case "String":
		return "string"
	default:
		return ""
	}
}

func GoType(str string) string {
	if native := GoNativeType(str); native != "" {
		return native
	}
	switch str {
	case "any":
		return "interface{}"
	case "object":
		return "interface{}"
	case "function":
		return "func(...interface{})"
	case "array":
		return "[]interface{}"
	case "Array.<any>":
		return "[]interface{}"
	default:
		if strings.HasPrefix(str, "Array.<") {
			typeInArray := strings.TrimPrefix(str, "Array.<")
			typeInArray = strings.TrimSuffix(typeInArray, ">")
			return "[]" + TypeNoNamespace(typeInArray)
		}
		if strings.HasPrefix(str, "array.<") {
			typeInArray := strings.TrimPrefix(str, "array.<")
			typeInArray = strings.TrimSuffix(typeInArray, ">")
			return "[]" + TypeNoNamespace(typeInArray)
		}
		return TypeNoNamespace(str)
	}
}

func GoTypeInArray(str string) string {
	if fullType := GoType(str); strings.HasPrefix(fullType, "[]") {
		return strings.TrimPrefix(fullType, "[]")
	}
	return ""
}
