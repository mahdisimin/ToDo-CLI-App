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

// Network configuration
const (
	network = "tcp"
	address = "127.0.0.1:8080"
)
