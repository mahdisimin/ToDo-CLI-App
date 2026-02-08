package contracts

type Storage interface {
	Save(string, string)
}

type Logger interface {
}

const (
	UserFilePath     = ".//userList.txt"
	TaskFilePath     = ".//TaskList.txt"
	CategoryFilePath = ".//CategoryList.txt"
)
