package categories

type SubCategory struct {
	Id    string
	Title string
	Url   string
}

type SubCategoryList struct {
	SubCategories []Category
}
