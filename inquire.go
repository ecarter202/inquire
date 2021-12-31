package inquire

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Command is a representation of an available command.
type Command struct {
	Label       string
	Description string
	Handler     func() error
}

type app struct {
	prefix      string
	prefixColor color.Attribute
	commands    []*Command
	done        bool
}

// Prefix will set the CLI prompt prefix (display name).
func (a *app) Prefix(prefix string) *app {
	a.prefix = prefix
	return a
}

func (a *app) PrefixColor(c color.Attribute) *app {
	a.prefixColor = c
	return a
}

// Commands will set the commands for the CLI app.
func (a *app) Commands(commands []*Command) *app {
	a.commands = commands
	return a
}

// Run will execute the CLI app.
func (a *app) Run() {
	for {
		if a.done {
			break
		}

		a.inquire()
	}
}

// Close kills the CLI app.
func (a *app) Close() {
	a.done = true
}

func (a *app) inquire() {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		color.New(a.prefixColor).Fprint(os.Stderr, a.prefix+" ")
		s, _ = r.ReadString('\n')
		t := strings.TrimSpace(s)
		if t == "q" || t == "quit" {
			fmt.Println("Goodbye!")
			time.Sleep(time.Millisecond * 300)
			a.Close()
			break
		}

		for _, c := range a.commands {
			if c.Label == t {
				if err := c.Handler(); err != nil {
					color.New(color.BgHiRed).Fprintln(os.Stderr, err.Error())
				}
			}
		}
	}
}

// New generates a new CLI app.
func New() *app {
	return &app{
		prefixColor: color.FgWhite,
	}
}
