package roaringcompany

func CreatedRecordId(passedId int64) uint64 {
	if passedId <= ZERO {
		return ZERO
	}
	return uint64(passedId)
}

func GetBusinessId(id uint64) uint64 {
	if id <= ONE {
		return ONE
	}
	return id
}

func GetSkipRows(rows uint64) uint64 {
	if rows <= ZERO {
		return ZERO
	}
	return rows
}

func GetTakeRows(rows uint64) uint64 {
	if rows <= ZERO {
		return FIVEHUNDRED
	}
	return rows
}

func GetId(passedId int64) uint64 {
	if passedId <= ZERO {
		return ZERO
	}
	return uint64(passedId)
}