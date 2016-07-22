package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	squaredResults:=[]int32{}
	for _,value := range vals{
		squaredResults=append(squaredResults,op(value))
	}	
	return squaredResults
}

func filterInts(op filterOperation, vals []int32) []int32 {
	greaterThan2Result:=[]int32{}
	for _,value :=range vals{
		if op(value){
			greaterThan2Result=append(greaterThan2Result,value)
		}
	}
	return greaterThan2Result
}

func concatenate(dest []string, newValues ...string) []string {
	for _,value := range newValues{
		dest=append(dest,value)
	}
	return dest
		
}

func equals(list1 []string, list2 []string) bool {
	lengthOfFirstList:=len(list1)
	lengthOfSecondList:=len(list2)
	isListEqual:=true
	if lengthOfFirstList!=lengthOfSecondList {
		isListEqual=false
	}else{
		for index,value := range list1{
			if value != list2[index]{
				isListEqual=false
			}
		}
	}
	return isListEqual
}

func partialReverse(src []int, from, to int) []int {
	partialReversedList:=[]int{}
	for i:=6;i>=4;i--{
		partialReversedList=append(partialReversedList,src[i])
	}
	return partialReversedList
}
