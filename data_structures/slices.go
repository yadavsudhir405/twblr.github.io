package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	result:=[]int32{}
	for _,value := range vals{
		result=append(result,mapOperation(value))
	}	
	return result
}

func filterInts(op filterOperation, vals []int32) []int32 {
	return []int32{}

}

func concatenate(dest []string, newValues ...string) []string {
	return nil
	
		
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
