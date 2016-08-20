package main

type Function struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Return      Type        `json:"returns"`
}

func (f *Function) GetNameUpperInitial() string {
	return UpperInitial(f.Name)
}

func (f *Function) GetDescriptionLines() []string {
	return SplitMultilines(f.Description)
}
