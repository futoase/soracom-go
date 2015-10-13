package operator

type Request struct {
  XSoracomApiKey string `json:"X-SORACOM-API-KEY"`
  XSoracomToken string `json:"X-SORACOM-TOKEN"`
  OperatorId string `json:"operatorId"`
  TokenTimeoutSeconds int `json:"timeout_seconds"`
}