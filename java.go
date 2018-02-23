package gopher4j

// Callable reperesents either a function or a class constructor
type Callable interface {
	Call(params ...interface{}) (interface{}, error)
}

//var _ Callable = JavaClass{}

type JavaClass struct {
	gw   *JavaGateway
	name string
}

func (c JavaClass) Name() string {
	return c.name
}

func (c JavaClass) Call(params ...interface{}) (interface{}, error) {
	command := CONSTRUCTOR_COMMAND_NAME +
		c.name + "\n" +
		// FIXME params
		END_COMMAND_PART
	answer, err := c.gw.sendCommand(command)
	if err != nil {
		return nil, err
	}

	return getReturnValue(answer, c.gw)
}

type JavaObject struct {
	gw *JavaGateway
	id string
}
