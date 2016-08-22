package main

type Function struct {
	Name        string      `json:"name"`
	Suffix      string      `json:"-"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Return      Type        `json:"returns"`
}

func (f *Function) GetNameUpperInitial() string {
	return UpperInitial(f.Name) + f.Suffix
}

func (f *Function) GetDescriptionLines() []string {
	return SplitMultilines(f.Description)
}

func (f *Function) GetParameterFullString() string {
	result := ""
	firstParameter := true
	for _, parameter := range f.Parameters {
		if firstParameter {
			result = result + parameter.Name + " " + parameter.Type.GetTypeOptionalPointer()
			firstParameter = false
		} else {
			result = result + ", " + parameter.Name + " " + parameter.Type.GetTypeOptionalPointer()
		}
	}
	return result
}

func (f *Function) GetParameterCallString() string {
	result := ""
	for _, parameter := range f.Parameters {
		result = result + ", " + parameter.Name
	}
	return result
}
