package action

type ActionType string

const (
	SelectVersion ActionType = "SelectVersion"
	SelectDay     ActionType = "SelectDay"
	ConfirmDeploy ActionType = "ConfirmDeploy"
)

func (t ActionType) Value() string {
	return string(t)
}
