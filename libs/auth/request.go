package auth

type Request struct {
	Email               string `json:"email"`
	Password            string `json:"password"`
	TokenTimeoutSeconds int    `json:"tokenTimeoutSeconds"`
	OneTimeToken        string `json:"token"`
}
