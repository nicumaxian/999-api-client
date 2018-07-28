package categories

type Category struct {
	Id    string
	Title string
	Url   string
}

type CategoryList struct {
	Categories []Category
}
