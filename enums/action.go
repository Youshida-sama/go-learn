package enums

type Action string

const (
	ActionLogin  Action = "login"
	ActionPut    Action = "put"
	ActionGet    Action = "get"
	ActionDelete Action = "delete"
	ActionLogout Action = "logout"
)

func (a Action) String() string {
	return string(a)
}

func (s Action) IsValid() bool {
	switch s {
	case ActionLogin, ActionPut, ActionGet, ActionDelete, ActionLogout:
		return true
	}
	return false
}

func GetAction(value string) Action {
	return map[string]Action{
		"login":  ActionLogin,
		"put":    ActionPut,
		"get":    ActionGet,
		"delete": ActionDelete,
		"logout": ActionLogout,
	}[value]
}
