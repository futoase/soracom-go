package operator

type Response struct {
  CreateDate string `json:"createDate"`
  UpdateDate string `json:"updateDate"`
  Description string `json:"description"`
  Email string `json:"email"`
  RootOperatorId string `json:"rootOperatorId"`
  apiToken string `json:"apiToken"`
  token string `json:"token"`
}
