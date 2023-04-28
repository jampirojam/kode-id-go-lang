package constant

type ProductType string

var Product = struct {
	CAR   ProductType
	MOTOR ProductType
	CYCLE ProductType
}{
	CAR:   "MOBIL",
	MOTOR: "MOTOR",
	CYCLE: "SEPEDA",
}

type ProductTypeId int

var ProductId = struct {
	CAR   ProductTypeId
	MOTOR ProductTypeId
	CYCLE ProductTypeId
}{
	CAR:   1,
	MOTOR: 2,
	CYCLE: 3,
}
