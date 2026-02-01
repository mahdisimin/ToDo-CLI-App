package storage

type InMemoryMap struct {
	data map[int][]byte
}

func (m *InMemoryMap) Save(string2 string) {
	println(string2)
}

func (m *InMemoryMap) Load(string) ([]byte, error) {

	return _, nil
}
