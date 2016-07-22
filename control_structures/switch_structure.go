package control_structures

const add string = "+"
const subtract string = "-"
const multiply string = "*"
const divide = "/"

func calculate(op string, arg1, arg2 float32) float32 {
	var result float32 = 0.0
	switch {
		case op==add:
			result=arg1+arg2
		case op==subtract:
			result=arg1-arg2
		case op==multiply:
			result=arg1*arg2
		case op==divide:
			result=arg1/arg2
	}
	return result
}