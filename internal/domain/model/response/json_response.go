package response

type Base struct {
	StatusCode int    `json:"status_code"`
	UUID       string `json:"uuid"`
	DateTime   string `json:"date_time"`
}

type Error struct {
	Base
	Error string `json:"error"`
}

type Generic struct {
	Base
	Description string `json:"description"`
	Data        any    `json:"data"`
}
