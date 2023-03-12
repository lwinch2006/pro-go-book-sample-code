package models

type CategoryNotExistError struct {
	Category string
}

func (e *CategoryNotExistError) Error() string {
	return "Category " + e.Category + " does not exist in product list"
}
