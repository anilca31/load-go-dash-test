	package main
	
	import "github.com/Mparaiso/lodash-go"
	import "fmt"
	
	func Main(){
		type Country struct {
			Name       string
		    Population int
		}
		var total int
		err := lo.In([]Country{{"France", 1000}, {"Spain", 5000}}).
			Map(func(country Country) int { return country.Population }).
			Reduce(func(total int, count int) int { return total + count }, 0).
			Out(&total)
		
		fmt.Println(err)
		fmt.Println(total)
		
		// Output:
		// <nil>
		// 6000
	}
