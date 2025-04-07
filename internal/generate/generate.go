package generate

import (
	"fmt"
	"strings"

	"github.com/mb-14/gomarkov"
)

func Generate(chain *gomarkov.Chain, start string) string {
	tokens := []string{gomarkov.StartToken}
	tokens = append(tokens, strings.Split(strings.TrimSpace(start), " ")...)
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[1:len(tokens)-1], " ")
}

func GenerateAdvanced(chain *gomarkov.Chain, start string) string {
	maxRetries := 50
	attempts := 0

	for attempts < maxRetries {
		text := Generate(chain, start)

		fmt.Printf("attempt %d/%d |%s| |%s|\n", attempts, maxRetries, strings.TrimSpace(text), strings.TrimSpace(start))
		if strings.Compare(strings.TrimSpace(text), strings.TrimSpace(start)) != 0 {
			return text
		}
		attempts += 1
	}

	return "не удалось сгенерировать продолжение"
}
