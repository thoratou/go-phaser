package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MemberType struct {
	Names []string
}

type Member struct {
	Name        string     `json:"name"`
	Access      string     `json:"access"`
	Virtual     bool       `json:"virtual"`
	Type        MemberType `json:"type,omitempty"`
	Description string     `json:"description"`
}

func (t *MemberType) UnmarshalJSON(buf []byte) error {
	//strange Json output issue, 1 empty type
	if string(buf) != "\"\"" {
		tmp := map[string][]string{}
		if err := json.Unmarshal(buf, &tmp); err != nil {
			return err
		}
		if len(tmp) != 1 {
			return fmt.Errorf("wrong number of fields in MemberType: %d != 1", len(tmp))
		}

		names, exists := tmp["names"]
		if !exists {
			return fmt.Errorf("missing names field in map")
		}

		//fmt.Println(names)
		t.Names = names
	}
	return nil
}

func (m *Member) GetNameUpperInitial() string {
	return UpperInitial(m.Name)
}

func (m *Member) GetDescriptionLines() []string {
	return SplitMultilines(m.Description)
}

func (m *Member) IsPrivate() bool {
	return m.Access == "private"
}

func (m *Member) GetGoType() string {
	if len(m.Type.Names) == 1 {
		return GoType(m.Type.Names[0])
	}
	return "interface{}"
}

func (m *Member) IsGoAnyType() bool {
	return m.GetGoType() == "interface{}"
}

func (m *Member) GetGoNativeType() string {
	if len(m.Type.Names) == 1 {
		return GoNativeType(m.Type.Names[0])
	}
	return ""
}

func (m *Member) IsGoNativeType() bool {
	return m.GetGoNativeType() != ""
}

func (m *Member) IsCallback() bool {
	return m.GetGoType() == "func(...interface{})"
}

func (m *Member) IsGenericArray() bool {
	return m.GetGoType() == "[]interface{}"
}

func (m *Member) IsArray() bool {
	return strings.HasPrefix(m.GetGoType(), "[]")
}

func (m *Member) GetGoTypeInArray() string {
	if len(m.Type.Names) == 1 {
		return GoTypeInArray(m.Type.Names[0])
	}
	return "interface{}"
}

func (m *Member) GetGopherjsCallName() string {
	if m.IsGoNativeType() {
		return UpperInitial(m.GetGoNativeType())
	}
	return ""
}
