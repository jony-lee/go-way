package main

type FormatFunc func(string, ...interface{}) (string, error)
func format(f FormatFunc,s string,a ...interface{})(string,error){
	return f(s,a...)
}
