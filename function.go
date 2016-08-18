package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Return struct {
	Type        []string `json:"-"`
	Description string   `json:"-"`
}

type Function struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Return      Return      `json:"returns"`
}

func (r *Return) UnmarshalJSON(buf []byte) error {
	//Type could be eith string type or []string type
	//specific Unmarshall to cope with it :/
	if string(buf) != "\"\"" {
		tmp := map[string]interface{}{}
		if err := json.Unmarshal(buf, &tmp); err != nil {
			return err
		}
		if len(tmp) != 2 {
			return fmt.Errorf("wrong number of fields in MemberType: %d != 2", len(tmp))
		}

		//fmt.Println(tmp)

		returnType, exists := tmp["type"]
		if !exists {
			return fmt.Errorf("missing names field in map")
		}

		switch returnType.(type) {
		case string:
			r.Type = make([]string, 1, 1)
			r.Type[0] = returnType.(string)
		default:
			typeArray := returnType.([]interface{})
			length := len(typeArray)
			r.Type = make([]string, length, length)
			for i := 0; i < length; i++ {
				r.Type[i] = typeArray[i].(string)
			}
		}

		returnDescription, exists := tmp["description"]
		if !exists {
			return fmt.Errorf("missing names field in map")
		}

		r.Description = returnDescription.(string)

	}
	return nil
}

func (f *Function) GetNameUpperInitial() string {
	return UpperInitial(f.Name)
}

func (f *Function) GetDescriptionLines() []string {
	return SplitMultilines(f.Description)
}

func (f *Function) GetGoReturnType() string {
	if len(f.Return.Type) == 1 {
		return GoType(f.Return.Type[0])
	}
	return "interface{}"
}

func (f *Function) IsReturnVoid() bool {
	return len(f.Return.Type) == 0 || (len(f.Return.Type) == 1 && f.Return.Type[0] == "")
}

func (f *Function) IsGoReturnAnyType() bool {
	return f.GetGoReturnType() == "interface{}"
}

func (f *Function) GetGoReturnNativeType() string {
	if len(f.Return.Type) == 1 {
		return GoNativeType(f.Return.Type[0])
	}
	return ""
}

func (f *Function) IsGoReturnNativeType() bool {
	return f.GetGoReturnNativeType() != ""
}

func (f *Function) IsReturnGenericArray() bool {
	return f.GetGoReturnType() == "[]interface{}"
}

func (f *Function) IsReturnArray() bool {
	return strings.HasPrefix(f.GetGoReturnType(), "[]")
}

func (f *Function) GetGoReturnTypeInArray() string {
	if len(f.Return.Type) == 1 {
		return GoTypeInArray(f.Return.Type[0])
	}
	return "interface{}"
}

func (f *Function) GetGoReturnNativeTypeInArray() string {
	if len(f.Return.Type) == 1 {
		return GoTypeNativeInArray(f.Return.Type[0])
	}
	return ""
}

func (f *Function) IsGoReturnNativeTypeInArray() bool {
	return f.GetGoReturnNativeTypeInArray() != ""
}

func (f *Function) GetReturnGopherjsCallName() string {
	if f.IsGoReturnNativeType() {
		return UpperInitial(f.GetGoReturnNativeType())
	}
	return ""
}

func (f *Function) GetReturnGopherjsCallNameInArray() string {
	if f.IsGoReturnNativeTypeInArray() {
		return UpperInitial(f.GetGoReturnNativeTypeInArray())
	}
	return ""
}
