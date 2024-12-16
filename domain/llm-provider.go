package domain

type LLMProvider interface {
	MakeChangelog(prompt string) (string, error)
}
