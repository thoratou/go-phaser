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

func (m *Member) GetDescriptionLines() []string {
	lines := strings.Replace(m.Description, "\r", "\n", -1)
	return strings.Split(lines, "\n")
}

func (m *Member) IsPrivate() bool {
	return m.Access == "private"
}

func (m *Member) GetGoType() string {
	if len(m.Type.Names) == 1 {
		if m.IsGoNativeType() {
			return m.GetGoNativeType()
		}

		goType := m.Type.Names[0]
		switch goType {
		case "any":
			return "interface{}"
		case "object":
			return "interface{}"
		default:
			return goType
		}
	}
	return "interface{}"
}

func (m *Member) IsGoAnyType() bool {
	return m.GetGoType() == "interface{}"
}

func (m *Member) GetGoNativeType() string {
	if len(m.Type.Names) == 1 {
		switch m.Type.Names[0] {
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
		case "string":
			return "string"
		case "String":
			return "string"
		default:
			return ""
		}
	}
	return ""
}

func (m *Member) IsGoNativeType() bool {
	return m.GetGoNativeType() != ""
}

func (m *Member) GetGopherjsCallName() string {
	if m.IsGoNativeType() {
		return UpperInitial(m.GetGoNativeType())
	}
	return ""
}
