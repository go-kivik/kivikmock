package kivikmock

import "fmt"

func optionsString(opt map[string]interface{}) string {
	if opt == nil {
		return "\n\t- has any options"
	}
	return fmt.Sprintf("\n\t- has options: %v", opt)
}

func errorString(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("\n\t- should return error: %s", err)
}
