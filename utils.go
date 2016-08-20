package main

import (
	"fmt"
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
	str = strings.TrimPrefix(str, "PIXI.") //yes twice :)
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
	case "bool":
		return "bool"
	case "integer":
		return "int"
	case "integers":
		return "int"
	case "Integer":
		return "int"
	case "int":
		return "int"
	case "number":
		return "float64"
	case "Number":
		return "float64"
	case "float64":
		return "float64"
	case "float":
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

func GoType(str string, typePackage string) string {
	if native := GoNativeType(str); native != "" {
		return native
	}
	switch str {
	case "any":
		return "interface{}"
	case "object":
		return "interface{}"
	case "Object":
		return "interface{}"
	case "*":
		return "interface{}"
	case "function":
		return "func(...interface{})"
	case "array":
		return "[]interface{}"
	case "Array":
		return "[]interface{}"
	case "Array.<any>":
		return "[]interface{}"
	case "Array.<Array.<any>>":
		return "[][]interface{}"
	default:
		if strings.HasPrefix(str, "Array.<") {
			typeInArray := strings.TrimPrefix(str, "Array.<")
			typeInArray = strings.TrimSuffix(typeInArray, ">")
			if native := GoNativeType(typeInArray); native != "" {
				return "[]" + native
			}
			return "[]" + AppendPackage(typePackage) + TypeNoNamespace(typeInArray)
		}
		if strings.HasPrefix(str, "array.<") {
			typeInArray := strings.TrimPrefix(str, "array.<")
			typeInArray = strings.TrimSuffix(typeInArray, ">")
			if native := GoNativeType(typeInArray); native != "" {
				return "[]" + native
			}
			return "[]" + AppendPackage(typePackage) + TypeNoNamespace(typeInArray)
		}
		return AppendPackage(typePackage) + TypeNoNamespace(str)
	}
}

func GoTypeInArray(str string, typePackage string) string {
	if fullType := GoType(str, typePackage); strings.HasPrefix(fullType, "[]") {
		return strings.TrimPrefix(fullType, "[]")
	}
	return ""
}

func GoTypeNativeInArray(str string, typePackage string) string {
	if fullType := GoType(str, typePackage); strings.HasPrefix(fullType, "[]") {
		typeInArray := strings.TrimPrefix(fullType, "[]")
		return GoNativeType(typeInArray)
	}
	return ""
}

func AppendPackage(typePackage string) string {
	if typePackage != "" {
		return typePackage + "."
	}
	return ""
}

func RemoveDuplicateMembers(elements []Member, className string) []Member {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []Member{}

	for _, v := range elements {
		if encountered[v.Name] == true {
			// Do not add duplicate.
			fmt.Println("warning: remove duplicate member: " + v.Name + " in class: " + className)
		} else {
			// Record this element as an encountered element.
			encountered[v.Name] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}

func RemoveDuplicateFunctions(elements []Function, className string) []Function {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []Function{}

	for _, v := range elements {
		if encountered[v.Name] == true {
			// Do not add duplicate.
			fmt.Println("warning: remove duplicate function: " + v.Name + " in class: " + className)
		} else {
			// Record this element as an encountered element.
			encountered[v.Name] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}

func AddMemberImports(class *Class, elements []Member, additionalPackages map[string]string, additionalImports map[string]string) []Member {
	result := []Member{}
	for _, v := range elements {
		//return type
		if addPackage, exists := additionalPackages[v.Type.GetType()]; exists {
			fmt.Println("note: add package: " + addPackage + " to type: " + v.Type.GetType())
			if addImport, exists := additionalImports[addPackage]; exists {
				class.Imports[addPackage] = addImport
				v.Type.Package = addPackage
			} else {
				fmt.Println("warning: missing import for package: " + addPackage)
			}
		}
		result = append(result, v)
	}
	return result
}

func AddFunctionImports(class *Class, elements []Function, additionalPackages map[string]string, additionalImports map[string]string) []Function {
	result := []Function{}
	for _, v := range elements {
		//return type
		if addPackage, exists := additionalPackages[v.Return.GetType()]; exists {
			fmt.Println("note: add package: " + addPackage + " to type: " + v.Return.GetType())
			if addImport, exists := additionalImports[addPackage]; exists {
				class.Imports[addPackage] = addImport
				v.Return.Package = addPackage
			} else {
				fmt.Println("warning: missing import for package: " + addPackage)
			}
		}
		result = append(result, v)
	}
	return result
}

func AddMemberWrappers(elements []Member, wrappers map[string]string) []Member {
	result := []Member{}
	for _, v := range elements {
		//return type
		if wrapper, exists := wrappers[v.Type.GetType()]; exists {
			fmt.Println("note: add wrapper: " + wrapper + " to type: " + v.Type.GetType())
			v.Type.Wrapper = wrapper
		}
		result = append(result, v)
	}
	return result
}

func AddFunctionWrappers(elements []Function, wrappers map[string]string) []Function {
	result := []Function{}
	for _, v := range elements {
		//return type
		if wrapper, exists := wrappers[v.Return.GetType()]; exists {
			fmt.Println("note: add wrapper: " + wrapper + " to type: " + v.Return.GetType())
			v.Return.Wrapper = wrapper
		}
		result = append(result, v)
	}
	return result
}
