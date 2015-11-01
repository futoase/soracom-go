package auth

type Request struct {
	Params RequestParams
}

type AuthRequest struct {
	Email               string `json:"email"`
	Password            string `json:"password"`
	TokenTimeoutSeconds int    `json:"tokenTimeoutSeconds"`
}

type IssueRequest struct {
	Email string `json:"email"`
}

type VerifyRequest struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}
