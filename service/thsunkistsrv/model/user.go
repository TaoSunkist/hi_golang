package model

type UserInfo struct {
	Uid    uint64  `json:"uid"`
	Name   *string `json:"name"`
	Age    uint8   `json:"age"`
	Avatar *string `json:"avatar"`
}

func Fake() UserInfo {
	var name = "Sunkist"
	var userInfo = UserInfo{10, &name, 20, nil}
	return userInfo
}
