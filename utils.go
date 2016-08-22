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
	if strings.Contains(str, "(") || strings.Contains(str, ")") || strings.Contains(str, "|") {
		return "interface{}"
	}
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
		return "int"
	case "Number":
		return "int"
	case "float64":
		return "float64"
	case "float":
		return "float64"
	case "numer":
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
	case "...interface{}":
		return "...interface{}"
	case "":
		return "interface{}"
	case "any":
		return "interface{}"
	case "object":
		return "interface{}"
	case "Object":
		return "interface{}"
	case "*":
		return "interface{}"
	case "Bounds-like":
		return "interface{}"
	case "Rectangle-like":
		return "interface{}"
	case "function":
		return "func(...interface{})"
	case "Functon":
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
		//parameters
		v.Parameters = AddParamaterImports(class, v.Parameters, additionalPackages, additionalImports)
		result = append(result, v)
	}
	return result
}

func AddParamaterImports(class *Class, elements []Parameter, additionalPackages map[string]string, additionalImports map[string]string) []Parameter {
	result := []Parameter{}
	for _, v := range elements {
		//return type
		if addPackage, exists := additionalPackages[v.Type.GetType()]; exists {
			fmt.Println("note: add package: " + addPackage + " to parameter type: " + v.Type.GetType())
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

func AddParameteredFunctions(elements []Function) []Function {
	result := []Function{}
	for _, v := range elements {
		//most specific function => keep the default one
		result = append(result, v)

		//optional parameter removal from last to start
		/*optionalNumber := 1
		for f.IsLastParameterOptional() {
			v.Suffix = "O" + string(optionalNumber)
			v.Parameters = v.Parameters[0 : len(v.Parameters)-1]
			optionalNumber++
		}*/

		//generic parameters
		v.Suffix = "I"
		v.Parameters = make([]Parameter, 1, 1)
		v.Parameters[0] = Parameter{
			Name: "args",
			Type: Type{
				Names:       []string{"...interface{}"},
				Description: "",
				Package:     "",
				Wrapper:     "",
			},
			Description: "",
			Optional:    false,
			Nullable:    false,
		}
		result = append(result, v)
	}
	return result
}

func ConvertGoReserveNames(str string) string {
	switch str {
	case "type":
		return "type_"
	case "map":
		return "map_"
	default:
		return str
	}
}

func IsBugParameter(str string) bool {
	if strings.HasPrefix(str, "options.") {
		//PhysicsP2.go
		//options interface{}, options.optimalDecomp bool, options.skipSimpleCheck bool, options.removeCollinearPoints interface{}
		//only option must be kept
		return true
	}
	if strings.HasPrefix(str, "style.") {
		//Text.go
		//style interface{}, style.font string, style.fontStyle string, style.fontVariant string, style.fontWeight string, style.fontSize interface{}, style.backgroundColor string, style.fill string, style.align string, style.boundsAlignH string, style.boundsAlignV string, style.stroke string, style.strokeThickness int, style.wordWrap bool, style.wordWrapWidth int, style.maxLines int, style.tabs interface{}
		//only style must be kept
		return true
	}
	return false
}
