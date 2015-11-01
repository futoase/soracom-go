package auth

type Request struct {
	Email               string
	Password            string
	TokenTimeoutSeconds int
	OneTimeToken        string
	VerifyToken         string
	NewPassword         string
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
