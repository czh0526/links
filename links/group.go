package links

type Group struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  int32  `json:"type"`
	Owner string `json:"owner"`
}

func GetMyGroups(id string) (map[string]*Group, error) {
	return map[string]*Group{
		"group-12345": {
			Id:   "group-12345",
			Name: "测试群1",
		},
		"friend-12346": {
			Id:   "friend-12346",
			Name: "测试群2",
		},
	}, nil
}
