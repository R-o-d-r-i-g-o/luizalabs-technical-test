package env

import (
	"encoding/json"
	"os"
	"reflect"
)

// LoadStructWithEnvVars loads environment variables into the provided configuration structures.
func LoadStructWithEnvVars(tagName string, configStructures ...interface{}) {
	for _, configStructure := range configStructures {
		reflectElement := reflect.ValueOf(configStructure).Elem()
		environments := make(map[string]string)

		elements := reflectElement.Type()
		for i := 0; i < elements.NumField(); i++ {
			field := elements.Field(i)
			fieldName, envVarName := field.Name, field.Tag.Get(tagName)

			environments[fieldName] = os.Getenv(envVarName)
		}

		parsed := parseMapToJSON(environments)
		json.Unmarshal([]byte(parsed), configStructure)
	}
}

// parseMapToJSON converts a map to a JSON string.
func parseMapToJSON(mp map[string]string) string {
	str, _ := json.Marshal(mp)
	return string(str)
}
