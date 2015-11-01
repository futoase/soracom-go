package operator

type Request struct {
	Params RequestParams
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
