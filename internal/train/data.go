package train

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"nsqd.cc/dich-markov/internal/utils"
)

func LoadData(jsonPath string) [][]string {
	jsonFileData, err := os.ReadFile(jsonPath)
	utils.CheckError(err)

	var jsonData map[string]interface{}
	utils.CheckError(json.Unmarshal(jsonFileData, &jsonData))

	data := make([][]string, 0)

	if jsonData["messages"] != nil {
		var messages []interface{} = jsonData["messages"].([]interface{})
		for _, message := range messages {
			messageData := make([]string, 0)
			switch v := message.(map[string]interface{})["text"].(type) {
			default:
				panic(errors.New("unknown \"text\" type"))
			case string:
				if !(len(strings.TrimSpace(v)) == 0) {
					messageData = append(messageData, strings.Split(strings.TrimSpace(strings.ToLower(v)), " ")...)
				}
			case []interface{}:
				for _, entity := range v {
					switch e := entity.(type) {
					default:
						panic(errors.New("unknown \"text\" entity type"))
					case string:
						if !(len(strings.TrimSpace(e)) == 0) {
							messageData = append(messageData, strings.Split(strings.ToLower(e), " ")...)
						}
					case map[string]interface{}:
						if !(len(strings.TrimSpace(e["text"].(string))) == 0) {
							messageData = append(messageData, strings.Split(strings.ToLower(e["text"].(string)), " ")...)
						}
					}
				}
			}
			if len(messageData) != 0 {
				data = append(data, messageData)
			}
		}
	} else {
		panic(errors.New("no \"messages\" in json"))
	}

	return data
}
