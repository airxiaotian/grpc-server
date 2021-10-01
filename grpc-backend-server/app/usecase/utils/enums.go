package utils

type SortEnum int32

const (
	SortEnumNONE SortEnum = 0
	SortEnumASC  SortEnum = 1
	SortEnumDESC SortEnum = 2
)

// Enum value maps for SortEnum.
var (
	SortEnumName = map[int32]string{
		0: "NONE",
		1: "ASC",
		2: "DESC",
	}
	SortEnumValue = map[string]int32{
		"NONE": 0,
		"ASC":  1,
		"DESC": 2,
	}
)
