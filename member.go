package main

import "strings"

type Member struct {
	Name        string `json:"name"`
	Access      string `json:"access"`
	Virtual     bool   `json:"virtual"`
	Type        Type   `json:"type"`
	Description string `json:"description"`
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
		return GoType(m.Type.Names[0], m.Type.Package)
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
		return GoTypeInArray(m.Type.Names[0], m.Type.Package)
	}
	return "interface{}"
}

func (m *Member) GetGopherjsCallName() string {
	if m.IsGoNativeType() {
		native := m.GetGoNativeType()
		if native == "float64" {
			return "Float"
		}
		return UpperInitial(native)
	}
	return ""
}
