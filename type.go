package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Type struct {
	Names       []string
	Description string
	Package     string
	Wrapper     string
}

func (t *Type) UnmarshalJSON(buf []byte) error {
	//Type could be eith string type or []string type
	//specific Unmarshall to cope with it :/
	if string(buf) != "\"\"" {
		tmp := map[string]interface{}{}
		if err := json.Unmarshal(buf, &tmp); err != nil {
			return err
		}

		//fmt.Println(tmp)

		returnType, exists := tmp["type"]
		if !exists {
			returnType, exists = tmp["names"]
		}
		if !exists {
			return fmt.Errorf("missing names field in map")
		}

		switch returnType.(type) {
		case string:
			t.Names = make([]string, 1, 1)
			t.Names[0] = returnType.(string)
		default:
			typeArray := returnType.([]interface{})
			length := len(typeArray)
			t.Names = make([]string, length, length)
			for i := 0; i < length; i++ {
				t.Names[i] = typeArray[i].(string)
			}
		}

		returnDescription, exists := tmp["description"]
		if exists {
			t.Description = returnDescription.(string)
		}
	}
	return nil
}

func (t *Type) GetType() string {
	if len(t.Names) == 1 {
		return GoType(t.Names[0], t.Package)
	}
	return "interface{}"
}

func (t *Type) GetTypeOptionalPointer() string {
	if !t.IsVoid() && !t.IsAnyType() && !t.IsNativeType() && !t.IsArray() && !t.IsCallback() && !t.HasWrapper() {
		return "*" + t.GetType()
	}
	return t.GetType()
}

func (t *Type) IsVoid() bool {
	return len(t.Names) == 0 || (len(t.Names) == 1 && t.Names[0] == "")
}

func (t *Type) IsCallback() bool {
	return t.GetType() == "func(...interface{})"
}

func (t *Type) IsAnyType() bool {
	return t.GetType() == "interface{}"
}

func (t *Type) GetNativeType() string {
	if len(t.Names) == 1 {
		return GoNativeType(t.Names[0])
	}
	return ""
}

func (t *Type) IsNativeType() bool {
	return t.GetNativeType() != ""
}

func (t *Type) IsGenericArray() bool {
	return t.GetType() == "[]interface{}"
}

func (t *Type) IsArray() bool {
	return strings.HasPrefix(t.GetType(), "[]")
}

func (t *Type) GetTypeInArray() string {
	if len(t.Names) == 1 {
		return GoTypeInArray(t.Names[0], t.Package)
	}
	return "interface{}"
}

func (t *Type) GetNativeTypeInArray() string {
	if len(t.Names) == 1 {
		return GoTypeNativeInArray(t.Names[0], t.Package)
	}
	return ""
}

func (t *Type) IsNativeTypeInArray() bool {
	return t.GetNativeTypeInArray() != ""
}

func (t *Type) GetGopherjsCallName() string {
	if t.IsNativeType() {
		native := t.GetNativeType()
		if native == "float64" {
			return "Float"
		}
		return UpperInitial(native)
	}
	return ""
}

func (t *Type) GetGopherjsCallNameInArray() string {
	if t.IsNativeTypeInArray() {
		native := t.GetNativeTypeInArray()
		if native == "float64" {
			return "Float"
		}
		return UpperInitial(native)
	}
	return ""
}

func (t *Type) HasWrapper() bool {
	return t.Wrapper != ""
}
