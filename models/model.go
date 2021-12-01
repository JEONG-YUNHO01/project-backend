package models

// 공공데이터 RESPONSE DATA 맵핑 모델
type ResData struct {
	Header Header `json:"header"`
	Body   Body   `json:"body"`
}
type Header struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}
type Items struct {
	EntpName            string `json:"entpName"`
	ItemName            string `json:"itemName"`
	ItemSeq             string `json:"itemSeq"`
	EfcyQesitm          string `json:"efcyQesitm"`
	UseMethodQesitm     string `json:"useMethodQesitm"`
	AtpnWarnQesitm      string `json:"atpnWarnQesitm"`
	AtpnQesitm          string `json:"atpnQesitm"`
	IntrcQesitm         string `json:"intrcQesitm"`
	SeQesitm            string `json:"seQesitm"`
	DepositMethodQesitm string `json:"depositMethodQesitm"`
	OpenDe              string `json:"openDe"`
	UpdateDe            string `json:"updateDe"`
	ItemImage           string `json:"itemImage"`
}
type Body struct {
	PageNo     int     `json:"pageNo"`
	TotalCount int     `json:"totalCount"`
	NumOfRows  int     `json:"numOfRows"`
	Items      []Items `json:"items"`
}
