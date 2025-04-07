package train

import (
	"github.com/mb-14/gomarkov"
)

func Train(data [][]string) *gomarkov.Chain {
	chain := gomarkov.NewChain(1)

	for _, message := range data {
		chain.Add(message)
	}

	return chain
}
