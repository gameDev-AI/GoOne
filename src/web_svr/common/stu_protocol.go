package define

type MsgSecCheckReq struct {
	AccountId string `json:"account_id"`
	NickName  string `json:"nick_name"`
	Time      string `json:"time"`
}
