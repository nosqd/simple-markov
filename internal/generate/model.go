package generate

import (
	"encoding/json"
	"os"

	"github.com/mb-14/gomarkov"
	"nsqd.cc/dich-markov/internal/utils"
)

func LoadChain(path string) *gomarkov.Chain {
	data, err := os.ReadFile(path)
	utils.CheckError(err)

	var chain gomarkov.Chain
	err = json.Unmarshal(data, &chain)
	utils.CheckError(err)

	return &chain
}
