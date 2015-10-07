package auth

type Response struct {
  ApiKey string `json:"apiKey"`
  OperatorId string `json:"operatorId"`
  Token string `json:"token"`
}
