package todo

type todoModel struct {
	title      string
	isComplete bool
}

func (tm todoModel) GetTitle() string {
	return tm.title
}

func (tm *todoModel) SetTitle(newTitle string) {
	tm.title = newTitle
}

func (tm *todoModel) Complete() {
	tm.isComplete = true
}
