package contracts

type storage interface {
	Save(string) error
	Load(string) ([]byte, error)
}
