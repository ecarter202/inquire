package inquire

type Command struct {
	label       string
	description string
	handler     func()
}

type app struct {
	prefix   string
	commands []*Command
}

func (a *app) Prefix(prefix string) *app {
	a.prefix = prefix
	return a
}

func (a *app) Commands(commands []*Command) *app {
	a.commands = commands
	return a
}

func New() *app {
	return new(app)
}
