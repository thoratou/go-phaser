package main

import "strings"

type Function struct {
	Name        string      `json:"name"`
	Suffix      string      `json:"-"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Return      Type        `json:"returns"`
}

func (f *Function) GetNameUpperInitial() string {
	return UpperInitial(f.Name)
}

func (f *Function) GetNameSplit() []string {
	return strings.Split(f.Name, ".")
}

func (f *Function) GetFirstDescriptionLine() string {
	return FirstDescriptionLine(f.Description)
}

func (f *Function) GetNextDescriptionLines() []string {
	return NextDescriptionLines(f.Description)
}

func (f *Function) GetParameterFullString() string {
	result := ""
	firstParameter := true
	for _, parameter := range f.Parameters {
		if !IsBugParameter(parameter.Name) {
			if firstParameter {
				result = result + ConvertGoReserveNames(parameter.Name) + " " + parameter.Type.GetTypeOptionalPointer()
				firstParameter = false
			} else {
				result = result + ", " + ConvertGoReserveNames(parameter.Name) + " " + parameter.Type.GetTypeOptionalPointer()
			}
		}
	}
	return result
}

func (f *Function) GetParameterCallString() string {
	result := ""
	firstParameter := true
	for _, parameter := range f.Parameters {
		if !IsBugParameter(parameter.Name) {
			if firstParameter {
				firstParameter = false
				result = result + ConvertGoReserveNames(parameter.Name)
			} else {
				result = result + ", " + ConvertGoReserveNames(parameter.Name)
			}
		}
	}
	return result
}

func (f *Function) GetParameterCallStringWithCommaPrefix() string {
	result := f.GetParameterCallString()
	if result != "" {
		result = ", " + result
	}
	return result
}
