package gopher4j

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

// JavaGateway reperesents a client connection to the java gateway
type JavaGateway struct {
	connection     io.ReadWriteCloser
	bufferedReader *bufio.Reader
	mutex          sync.Mutex
	jvmID          string
}

// NewGatewayFromConnection returns a gateway client instance using a provided connection
func NewGatewayFromConnection(rwc io.ReadWriteCloser) (*JavaGateway, error) {
	return &JavaGateway{connection: rwc, jvmID: DEFAULT_JVM_ID}, nil
}

// NewGateway connects to a java gateway at the given host/port and returns a client connection
func NewGateway(host, port string) (*JavaGateway, error) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}
	return NewGatewayFromConnection(conn)
}

// ErrEmptyResponse is returned whenever the java gateway returned an empty response
var ErrEmptyResponse = errors.New("the gateway returned an empty response")

func (gw *JavaGateway) sendCommand(cmd string) (string, error) {
	gw.mutex.Lock()
	defer gw.mutex.Unlock()

	_, err := gw.connection.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	if gw.bufferedReader == nil {
		gw.bufferedReader = bufio.NewReader(gw.connection)
	}

	line, _, err := gw.bufferedReader.ReadLine()
	if err != nil {
		return "", err
	}

	answer := string(line)
	if strings.HasPrefix(answer, "!") {
		answer = answer[1:]
	}

	if len(answer) == 0 {
		return "", ErrEmptyResponse
	}

	return answer, nil
}

// Get retrieves a java package or class described by the given name
func (gw *JavaGateway) Get(name string) (interface{}, error) {
	cmd := REFLECTION_COMMAND_NAME +
		REFL_GET_UNKNOWN_SUB_COMMAND_NAME +
		name + "\n" +
		gw.jvmID + "\n" +
		END_COMMAND_PART
	answer, err := gw.sendCommand(cmd)
	if err != nil {
		return nil, err
	}

	if answer == SUCCESS_PACKAGE {
		return nil, fmt.Errorf("%q is a package - gopher4j doesn't support getting these yet", name)
	} else if strings.HasPrefix(answer, SUCCESS_CLASS) {
		return JavaClass{gw: gw, name: answer[CLASS_FQN_START:]}, nil
	} else {
		return nil, fmt.Errorf("%s does not exist in the JVM", name)
	}
}
