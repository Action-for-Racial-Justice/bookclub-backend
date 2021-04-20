package bcerrors

//Incremental integer wrappers for different error types
const (
	InternalError = iota
	ValidationError
	DecodeError
)

func getTypeCode(errorDescriptor int) int {

	switch errorDescriptor {
	case InternalError:
		return 500

	case DecodeError:
		return 404

	default:
		return 404
	}
}

func getTypeString(errorDescriptor int) string {

	switch errorDescriptor {
	case InternalError:
		return "Internal Server Error"

	case DecodeError:
		return "JSON Decode Error"

	default:
		return "Validation Error"
	}

}
