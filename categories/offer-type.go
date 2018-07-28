package categories

type OfferType struct {
	Id    string
	Title string
}

type OfferTypeList struct {
	OfferTypes []OfferType `json:"offer_types"`
}
