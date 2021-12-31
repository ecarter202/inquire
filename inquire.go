package inquire

type Command struct {
	label       string
	description string
	handler     func()
}

type app struct {
	prefix   string
	commands []*Command
	done     <-chan bool
}

func (a *app) Prefix(prefix string) *app {
	a.prefix = prefix
	return a
}

func (a *app) Commands(commands []*Command) *app {
	a.commands = commands
	return a
}

func (a *app) Run() {
	for {
		if <-a.done {
			break
		}
	}
}

func New() *app {
	return &app{
		done: make(chan bool, 1),
	}
}
