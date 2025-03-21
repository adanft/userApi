package maps

var ApiErrorStatusCode = map[string]int{
	"BAD_REQUEST": 400,
	"VALIDATION":  400,
	"NOT_FOUND":   404,
	"CONFLICT":    409,
	"SERVER":      500,
}
