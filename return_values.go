package gopher4j

import (
	"fmt"
	"strconv"
	"strings"
)

func getReturnValue(answer string, gw *JavaGateway) (interface{}, error) {
	if isAnswerError(answer) {
		return nil, fmt.Errorf("answer returned an error") // FIXME more error handling and more descriptive error message
	}

	typ := answer[1:2]
	if typ == VOID_TYPE {
		return nil, nil
	}

	return convertReturnValue(typ, answer[2:], gw)
}

func convertReturnValue(typ, value string, gw *JavaGateway) (interface{}, error) {
	switch typ {
	case NULL_TYPE:
		return nil, nil
	case BOOLEAN_TYPE:
		return strconv.ParseBool(value)
	case LONG_TYPE:
		return strconv.ParseInt(value, 10, 32)
	case DECIMAL_TYPE:
		return strconv.ParseInt(value, 10, 32)
	case INTEGER_TYPE:
		return strconv.ParseInt(value, 10, 32)
	case BYTES_TYPE:
		return nil, fmt.Errorf("byte arrays are not yet supported by gopher4j")
	case DOUBLE_TYPE:
		return strconv.ParseFloat(value, 64)
	case STRING_TYPE:
		return unescapeString(value), nil
	case MAP_TYPE:
		return nil, fmt.Errorf("maps are not yet supported by gopher4j")
	case LIST_TYPE:
		return nil, fmt.Errorf("lists are not yet supported by gopher4j")
	case ARRAY_TYPE:
		return nil, fmt.Errorf("arrays are not yet supported by gopher4j")
	case SET_TYPE:
		return nil, fmt.Errorf("sets are not yet supported by gopher4j")
	case ITERATOR_TYPE:
		return nil, fmt.Errorf("iterators are not yet supported by gopher4j")
	case REFERENCE_TYPE:
		return JavaObject{id: value, gw: gw}, nil
	default:
		return nil, fmt.Errorf("unknown value type %q", typ)
	}
}

func unescapeString(input string) string {
	rSplit := strings.Split(input, ESCAPE_CHAR+"r")
	r := strings.Join(rSplit, "\r")
	nSplit := strings.Split(r, ESCAPE_CHAR+"n")
	n := strings.Join(nSplit, "\n")
	sSplit := strings.Split(n, ESCAPE_CHAR+ESCAPE_CHAR)
	s := strings.Join(sSplit, ESCAPE_CHAR)
	return s
}

func isAnswerError(answer string) bool {
	return len(answer) == 0 || (answer[:1] != SUCCESS)
}
