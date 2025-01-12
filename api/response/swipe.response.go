package response

type Swipe struct {
	Message string `json:"message"`
	Match   bool   `json:"match"`
}
