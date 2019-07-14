package display

type Display interface {
	Refresh() error
	Render() string
}
