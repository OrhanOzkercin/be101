package blog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type LimitType struct {
	HasValue bool
	Value    int
}

func GetBlogTitles(limit LimitType) []string {
	jsonFile, err := os.Open("dummydata/titles.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var titles []string
	json.Unmarshal(byteValue, &titles)

	if limit.HasValue {
		return titles[:limit.Value]
	}

	return titles
}
