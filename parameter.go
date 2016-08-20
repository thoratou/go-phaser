package main

import (
	"encoding/json"
	"fmt"
)

type Parameter struct {
	Name        string `json:"name"`
	Type        Type   `json:"type"`
	Description string `json:"description"`
	//	Default     interface{} `json:"default"`
	Optional bool `json:"optional"`
	Nullable bool `json:"nullable"`
}

func (p *Parameter) UnmarshalJSON(buf []byte) error {
	//Type could be eith string type or []string type
	//specific Unmarshall to cope with it :/
	if string(buf) != "\"\"" {
		tmp := map[string]interface{}{}
		if err := json.Unmarshal(buf, &tmp); err != nil {
			return err
		}

		//fmt.Println(tmp)

		name, exists := tmp["name"]
		if exists {
			p.Name = name.(string)
		} else {
			fmt.Println("warning: missing parameter name, put args as default:")
			fmt.Println(tmp)
			p.Name = "args"
		}

		parType, exists := tmp["type"]
		if !exists {
			return fmt.Errorf("missing type field in map")
		}

		switch parType.(type) {
		case string:
			p.Type.Names = make([]string, 1, 1)
			p.Type.Names[0] = parType.(string)
		default:
			typeArray := parType.([]interface{})
			length := len(typeArray)
			p.Type.Names = make([]string, length, length)
			for i := 0; i < length; i++ {
				p.Type.Names[i] = typeArray[i].(string)
			}
		}

		description, exists := tmp["description"]
		if exists {
			p.Description = description.(string)
		}

		//no default suport in Go

		optional, exists := tmp["optional"]
		if !exists {
			return fmt.Errorf("missing optional field in map")
		}

		switch optional.(type) {
		case bool:
			p.Optional = optional.(bool)
		default:
			p.Optional = false
		}

		nullable, exists := tmp["nullable"]
		if !exists {
			return fmt.Errorf("missing nullable field in map")
		}

		switch nullable.(type) {
		case bool:
			p.Nullable = nullable.(bool)
		default:
			p.Nullable = false
		}
	}
	return nil
}
