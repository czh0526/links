package links

type Friend struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Owner    string `json:"owner"`
}

func GetMyFriends(id string) (map[string]*Friend, error) {
	return map[string]*Friend{
		"friend-12345": {
			Id:       "friend-12345",
			Nickname: "梁川",
		},
		"friend-12346": {
			Id:       "friend-12346",
			Nickname: "阿鸿",
		},
	}, nil
}
