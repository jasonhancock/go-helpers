package helpers

import "strings"

// CleanupCommaDelimited cleans up a comma delimted list of strings, removing
// empty entries and trimming whitespace.
func CleanupCommaDelimited(str string) []string {
	data := make([]string, 0)
	for _, v := range strings.Split(str, ",") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		data = append(data, v)
	}

	return data
}
