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

// JavaObject represents a reference to a object in the JVM
type JavaObject struct {
	gw                  *JavaGateway
	id                  string
	wasGarbageCollected bool
}

// GarbageCollect will tell the JVM that we no longer hold a reference to a given object.
//
// Should be called on every acquired JavaObject after it is no longer used.
//
// Must be called manually, since finalizers are not guaranteed to run.
func (obj JavaObject) GarbageCollect() {
	if !obj.wasGarbageCollected {
		cmd := MEMORY_COMMAND_NAME +
			MEMORY_DEL_SUBCOMMAND_NAME +
			obj.id + "\n" +
			END_COMMAND_PART
		_, err := obj.gw.sendCommand(cmd)
		if err != nil && obj.gw.DebugLogger != nil {
			obj.gw.DebugLogger.Printf("Failed to garbage-collect Java object %q", obj.id)
		}
		obj.wasGarbageCollected = true
	}
}
