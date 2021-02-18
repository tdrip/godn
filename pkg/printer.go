package path

import "fmt"

//PrintDetails this prints the path info to fmt
func PrintDetails(info *Info) {

	if info == nil {
		fmt.Println("Info is nil")
	} else {

		fmt.Println("")
		fmt.Printf("Top: %s", info.Top)
		fmt.Println("")
		fmt.Printf("IsTop: %v", info.IsTop())
		fmt.Println("")
		fmt.Printf("OriginalValue: %s", info.OriginalValue)
		fmt.Println("")
		fmt.Printf("ParsedPath: %s", info.ParsedPath)
		fmt.Println("")
		fmt.Printf("Name: %s", info.Name)
		fmt.Println("")
		fmt.Printf("Parent: %s", info.Parent)
		fmt.Println("")
		fmt.Printf("Parsed: %t", info.Parsed)
		fmt.Println("")
		fmt.Printf("Seperator: %c", info.Seperator)
		fmt.Println("")
		fmt.Printf("String: %s", info)
		fmt.Println("")
	}
}
