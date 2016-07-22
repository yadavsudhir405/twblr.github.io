package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	results:=[]string{}
	for superhero,skills := range data{
		if contains(skills,interest){
			results=append(results,superhero)
		}
	}
	return results
}

func contains(src []string, elem string) bool {
	found:=false
	for _,value:= range src{
		if elem==value {
			found=true;
		}
	}
	return found
}
