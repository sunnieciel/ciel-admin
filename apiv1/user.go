package apiv1

type LoginVo struct {
	Uname      string `json:"uname"`
	Nickname   string `json:"nickname"`
	Icon       string `json:"icon"`
	Summary    string `json:"summary"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	GoldStatus uint   `json:"gold_status"`
	Token      string `json:"token"`
}
