package schemas

type Address struct {
	StreetAddress string `json:"street_address" bson:"street_address"`
	City          string `json:"city" bson:"city"`
	Province      string `json:"province" bson:"province"`
	PostalCode    string `json:"postal_code" bson:"postal_code"`
	Country       string `json:"country" bson:"country"`
}
