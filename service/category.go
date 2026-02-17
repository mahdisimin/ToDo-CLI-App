package service

type CategoryServiceRepository interface {
	UserHaveThisCategory(int, int) bool
	PersistCategory()
}
