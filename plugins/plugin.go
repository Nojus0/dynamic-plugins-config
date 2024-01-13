package plugins

type Impl interface {
	HandlePart(m map[string]any) error
	Info() (string, string)
}
