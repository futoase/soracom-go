package operator

type Request struct {
	XSoracomApiKey      string `json:"X-SORACOM-API-KEY"`
	XSoracomToken       string `json:"X-SORACOM-TOKEN"`
	OperatorId          string `json:"operatorId"`
	TokenTimeoutSeconds int    `json:"timeout_seconds"`
	CurrentPassword     string `json:"currentPassword"`
	NewPassword         string `json:"newPassword"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	VerifyToken         string `json:"token"`
}

type RequestPassword struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type RequestOperators struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestVerify struct {
	VerifyToken string `json:"token"`
}
