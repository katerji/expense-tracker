package expense

type merchant struct {
	ID           uint32
	Name         string
	merchantType *merchantType
}

type merchantType struct {
	ID   uint32
	Type string
}

type createMerchantInput struct {
	Name   string
	TypeID uint32
}
