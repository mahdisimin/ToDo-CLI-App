package contracts

type Storage interface {
	Save(string, string)
}

type Logger interface {
}

const (
	UserFilePath     = ".//userList.tex"
	TaskFilePath     = ".//TaskList.text"
	CategoryFilePath = ".//CategoryList.text"
)
