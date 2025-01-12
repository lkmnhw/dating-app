package response

type LogIn struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
