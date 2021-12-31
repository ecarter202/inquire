package inquire

// Command is a representation of an available command.
type Command struct {
	Label       string
	Description string
	Handler     func() error
}
