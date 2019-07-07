package display

type Display interface {
	Refresh() bool
	Render() string
}
