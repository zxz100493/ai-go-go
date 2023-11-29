package ai

type AiChat interface {
	Chat(string) string
	ParseResult(string)
}
