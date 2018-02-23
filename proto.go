package gopher4j

const (
	ESCAPE_CHAR = "\\"

	// Entry point
	ENTRY_POINT_OBJECT_ID         = "t"
	CONNECTION_PROPERTY_OBJECT_ID = "c"
	GATEWAY_SERVER_OBJECT_ID      = "GATEWAY_SERVER"
	STATIC_PREFIX                 = "z:"

	// JVM
	DEFAULT_JVM_ID   = "rj"
	DEFAULT_JVM_NAME = "default"

	// Types
	BYTES_TYPE        = "j"
	INTEGER_TYPE      = "i"
	LONG_TYPE         = "L"
	BOOLEAN_TYPE      = "b"
	DOUBLE_TYPE       = "d"
	DECIMAL_TYPE      = "D"
	STRING_TYPE       = "s"
	REFERENCE_TYPE    = "r"
	ARRAY_TYPE        = "t"
	SET_TYPE          = "h"
	LIST_TYPE         = "l"
	MAP_TYPE          = "a"
	NULL_TYPE         = "n"
	PACKAGE_TYPE      = "p"
	CLASS_TYPE        = "c"
	METHOD_TYPE       = "m"
	NO_MEMBER         = "o"
	VOID_TYPE         = "v"
	ITERATOR_TYPE     = "g"
	PYTHON_PROXY_TYPE = "f"

	// Protocol
	END            = "e"
	ERROR          = "x"
	FATAL_ERROR    = "z"
	SUCCESS        = "y"
	RETURN_MESSAGE = "!"

	// Shortcuts
	SUCCESS_PACKAGE   = SUCCESS + PACKAGE_TYPE
	SUCCESS_CLASS     = SUCCESS + CLASS_TYPE
	CLASS_FQN_START   = 2
	END_COMMAND_PART  = END + "\n"
	NO_MEMBER_COMMAND = SUCCESS + NO_MEMBER

	CALL_COMMAND_NAME             = "c\n"
	FIELD_COMMAND_NAME            = "f\n"
	CONSTRUCTOR_COMMAND_NAME      = "i\n"
	SHUTDOWN_GATEWAY_COMMAND_NAME = "s\n"
	LIST_COMMAND_NAME             = "l\n"
	REFLECTION_COMMAND_NAME       = "r\n"
	MEMORY_COMMAND_NAME           = "m\n"
	HELP_COMMAND_NAME             = "h\n"
	ARRAY_COMMAND_NAME            = "a\n"
	JVMVIEW_COMMAND_NAME          = "j\n"
	EXCEPTION_COMMAND_NAME        = "p\n"
	DIR_COMMAND_NAME              = "d\n"
	STREAM_COMMAND_NAME           = "S\n"

	// Array subcommands
	ARRAY_GET_SUB_COMMAND_NAME    = "g\n"
	ARRAY_SET_SUB_COMMAND_NAME    = "s\n"
	ARRAY_SLICE_SUB_COMMAND_NAME  = "l\n"
	ARRAY_LEN_SUB_COMMAND_NAME    = "e\n"
	ARRAY_CREATE_SUB_COMMAND_NAME = "c\n"

	// Reflection subcommands
	REFL_GET_UNKNOWN_SUB_COMMAND_NAME         = "u\n"
	REFL_GET_MEMBER_SUB_COMMAND_NAME          = "m\n"
	REFL_GET_JAVA_LANG_CLASS_SUB_COMMAND_NAME = "c\n"
)
