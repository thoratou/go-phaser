package main

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
