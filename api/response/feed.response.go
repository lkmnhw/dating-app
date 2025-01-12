package response

type Feed struct {
	Message string            `json:"message"`
	Data    []DataFeedProfile `json:"data"`
}

type DataFeedProfile struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
	Age         int    `json:"date_of_birth"`
}
