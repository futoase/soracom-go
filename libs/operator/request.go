package operator

type Request struct {
  XSoracomApiKey string `json:"X-SORACOM-API-KEY"`
  XSoracomToken string `json:"X-SORACOM-TOKEN"`
  OperatorId string `json:"operatorId"`
  TokenTimeoutSeconds int `json:"timeout_seconds"`
  CurrentPassword string `json:"currentPassword"`
  NewPassword string `json:"newPassword"`
}

type RequestPassword struct {
  CurrentPassword string `json:"currentPassword"`
  NewPassword string `json:"newPassword"`
}
