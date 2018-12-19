package ripper

const (
	errCodeNegetivePage int = iota + 1
	errCodeNegetivePageSize
	errCodeInvalidPageToken
)

var errCodeToField = map[int]string{
	errCodeNegetivePage:     "page",
	errCodeNegetivePageSize: "page_size",
	errCodeInvalidPageToken: "page_token",
}

var errCodeToMessage = map[int]string{
	errCodeNegetivePage:     "page must be greater than 0",
	errCodeNegetivePageSize: "page_size must be greater than 0",
	errCodeInvalidPageToken: "page_token is invalid",
}

// PaginationError is an interface for pagination error.
type PaginationError interface {
	error
	Field() string
	paginationError()
}

type paginationError struct {
	code int
}

func (e *paginationError) Error() string {
	return errCodeToMessage[e.code]
}

func (e *paginationError) Field() string {
	return errCodeToField[e.code]
}

func (*paginationError) paginationError() {}

func newNegetivePageError() *paginationError {
	return &paginationError{
		code: errCodeNegetivePage,
	}
}

func newNegetivePageSizeError() *paginationError {
	return &paginationError{
		code: errCodeNegetivePageSize,
	}
}

func newInvalidPageTokenError() *paginationError {
	return &paginationError{
		code: errCodeInvalidPageToken,
	}
}
