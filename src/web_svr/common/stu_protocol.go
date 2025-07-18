package define

type MsgSecCheckReq struct {
	AccountId  string `json:"account_id"`
	MsgContent string `json:"msg_content"`
	Time       string `json:"time"`
}
