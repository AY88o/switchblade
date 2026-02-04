package profile

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveProfile(p Profile) error {

	data, err1 := json.MarshalIndent(p.Apps, "", " ")

	if err1 != nil {
		return fmt.Errorf("couldn't freeze profile %w", err1)
	}

	filename := p.Name + ".json"

	err2 := os.WriteFile(filename, data, 0644)

	if err2 != nil {
		return fmt.Errorf("couldn't save the file to the disk: %w", err2)
	}

	return nil

}

func LoadProfile(name string) ([]string, error) {

	filename := name + ".json"

	data, err := os.ReadFile(filename)
	if err != nil {

		return nil, fmt.Errorf("file not found: %s", filename)

	}

	var p Profile

	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, fmt.Errorf("file corrupted: %w", err)

	}

	return p.Apps, nil

}
