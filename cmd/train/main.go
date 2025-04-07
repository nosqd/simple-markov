package main

import (
	"encoding/json"
	"os"

	"nsqd.cc/dich-markov/internal/train"
	"nsqd.cc/dich-markov/internal/utils"
)

func main() {
	data := train.LoadData("./data-hack/result.json")

	jsonData, err := json.Marshal(data)
	utils.CheckError(err)

	chain := train.Train(data)
	jsonChainData, err := json.Marshal(chain)
	utils.CheckError(err)

	err = os.WriteFile("text.json", jsonData, 0644)
	utils.CheckError(err)

	err = os.WriteFile("model.json", jsonChainData, 0644)
	utils.CheckError(err)
}
