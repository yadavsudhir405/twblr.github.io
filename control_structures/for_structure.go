package control_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	return []int32{}
}

func filterInts(op filterOperation, vals []int32) []int32 {
	return []int32{3}
}
