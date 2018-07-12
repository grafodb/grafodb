package middleware

type ContextKey string

const (
	KeyRequestID     ContextKey = "__req_id__"
	KeyRequestLogger ContextKey = "__req_logger__"
)
