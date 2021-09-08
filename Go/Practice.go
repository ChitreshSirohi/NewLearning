package main

import "fmt"

func main()  {
 fmt.Println("HelloWorld")
  x := [...]int{1,2,3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
	fmt.Println(x==y)
	var zSlice = make([]int,5)
	var zSliceWithCap = make([]int,0,10)
	fmt.Println("zSlice" , len(zSlice))
	fmt.Println("zSlice",append(zSlice,2),append(zSliceWithCap,2))
	fmt.Println("zSlice cap",cap(zSlice),cap(zSliceWithCap))
	fmt.Println(x[1:])
	ax := []int{1, 2, 3, 4}
	num := copy(ax[:3], ax[1:])
	fmt.Println(ax, num)
	var s string = "Hello there"
	var s2 string = s[4:7]
	fmt.Println(s2)
	fmt.Println(s[:7])

	totWins := map[string]int{
		"Oracs":1,
		"Sharks":2,

	}
	fmt.Println(len(totWins),totWins)
	totWins["Oracs"]++
	fmt.Println(totWins["Oracs"])
	i, ok := totWins["Orac"]
	fmt.Println(i,ok)

	if ok {
		fmt.Println("Its true")
	}else {
		fmt.Println("Its false")
	}


	for k,v := range "totWins" {
		fmt.Println(k,string(v))
	}
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

}
