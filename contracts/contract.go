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

// Network configuration
const (
	network = "tcp"
	address = "127.0.0.1:8080"
)
