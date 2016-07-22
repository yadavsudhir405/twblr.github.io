package control_structures

func fizzBuzz(i int32) string {
	result:="Invalied"
	if i%2==0{
		result="2"
	}else if i%3==0{
		result="Fizz"
	}else if i%5==0{
		result="Buzz"
	}else if i%15==0{
		result="FizzBuzz"
	}else if i%999==0{
		result="Fizz"
	}else{
		result="Invalid"
	}
	return result
}
