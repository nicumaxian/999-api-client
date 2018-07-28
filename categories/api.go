package categories

import (
	"github.com/maxiannicu/999-api-client/client"
	"fmt"
)

func GetCategories(client *client.Client) (*CategoryList, error) {
	list := &CategoryList{}
	err := client.Get("/categories", list)

	return list, err
}

func GetSubCategories(client *client.Client, categoryId string) (*SubCategoryList, error) {
	list := &SubCategoryList{}
	err := client.Get(fmt.Sprintf("/categories/%v/subcategories", categoryId), list)

	return list, err
}

func GetOfferTypes(client *client.Client, categoryId string, subCategoryId string) (*OfferTypeList, error) {
	list := &OfferTypeList{}
	err := client.Get(fmt.Sprintf("/categories/%v/subcategories/%v/offer-types", categoryId, subCategoryId), list)

	return list, err

}
