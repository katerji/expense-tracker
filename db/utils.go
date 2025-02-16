package db

func fromLastInsertIDtoUint32(id int64, err error) (uint32, bool) {
	if err != nil {
		return 0, false
	}

	return uint32(id), true
}
