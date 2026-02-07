package contracts

type Storage interface {
	Save(string, string)
}

type Logger interface {
}

const (
	UserFilePath     = ".//userList.text"
	TaskFilePath     = ".//TaskList.text"
	CategoryFilePath = ".//CategoryList.text"
)
