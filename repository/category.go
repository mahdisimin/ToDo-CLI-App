package repository

type CategoryStorage interface {
	Save(string, string)
	Load(entity string) (data string, dataLength int, err error)
	IsExists(entity string, attrib ...any) bool
}
type Category struct {
	storage CategoryStorage
}

func (c *Category) UserHaveThisCategory(userId int, categoryId int) bool {

	ok := c.storage.IsExists("Category", userId, categoryId)

	if !ok {
		return false
	}
	return true
}
