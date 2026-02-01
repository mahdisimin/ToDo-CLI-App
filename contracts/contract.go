package contracts

type Storage interface {
	Save(string)

	Load(string) ([]byte, error)
}
