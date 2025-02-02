package parser

type Parser interface {
	Parse(string) map[string]any
}
