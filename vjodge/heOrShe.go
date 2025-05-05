package vjodge

import (
	"fmt"
	"strings"
)

func HeOrShe() {
	var name string
	fmt.Scan(&name)
	myname := map[string]int{}
	listChar := strings.Split(name, "")
	for i := 0; i < len(listChar); i++ {
		myname[listChar[i]] = 1
	}
	fmt.Println(myname)
	if len(myname)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
