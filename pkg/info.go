package path

import (
	"strings"
)

// DefaultSeperator is back slahes
const DefaultSeperator = '\\'

//DefaultTop is the top of the tree in this case it is empty
const DefaultTop = string(DefaultSeperator) + ""

// Info represents an object's unique path on a platform
type Info struct {

	// This is the Top Item in the Tree
	Top *Info

	// The Original Value Passed in
	OriginalValue string

	// The Path that is parsed
	ParsedPath string

	// The Name
	Name string

	// Parent
	Parent *Info

	//has it been parsed?
	Parsed bool

	// what are we
	Seperator byte
}

func NewInfo(path string) *Info {
	info := MakeDefaultInfo()

	info.OriginalValue = path

	// mark as unparsed
	info.Parsed = false
	return info
}

//MakeDefaultInfo only creates a default Info object
func MakeDefaultInfo() *Info {
	// build object
	info := new(Info)

	// set an empty root
	info = info.ParseTop("")

	// set the defult seperator
	info.Seperator = DefaultSeperator

	// mark as unparsed
	info.Parsed = false

	return info
}

func (pathi *Info) ParseTop(root string) *Info {

	//fmt.Println("ParseTop: ", root)
	Root := ""
	if root == "" {
		// Default
		Root = string(pathi.Seperator)
	} else {

		if root[0] != pathi.Seperator && root[len(root)-1] != pathi.Seperator {
			Root = string(pathi.Seperator) + root + string(pathi.Seperator)
			//fmt.Println("pathi.Root = string(pathi.Seperator) + root + string(pathi.Seperator): ", pathi.Root)
		} else {

			if root[0] == pathi.Seperator && root[len(root)-1] != pathi.Seperator {
				Root = root + string(pathi.Seperator)
				//fmt.Println("root + string(pathi.Seperator): ", pathi.Root)
			} else if root[0] != pathi.Seperator && root[len(root)-1] == pathi.Seperator {
				Root = string(pathi.Seperator) + root
				//fmt.Println("pathi.Root = string(pathi.Seperator) + root: ", pathi.Root)
			} else {
				Root = root
				//fmt.Println("pathi.Root = root: ", pathi.Root)
			}
		}
	}

	if Root == "" {

	} else {

	}
	//fmt.Println("pathi.Root: ", pathi.Root)

	return pathi

}

func (pathi *Info) IsTop() bool {
	if pathi == nil {
		return false
	} else {
		lroot := strings.ToLower(pathi.Top.String())
		lpath := strings.ToLower(pathi.String())
		if len(lpath) == 1 {
			return lpath[0] == pathi.Seperator
		}
		return lpath == string(pathi.Seperator)+lroot+string(pathi.Seperator) || lpath == lroot+string(pathi.Seperator) || lpath == string(pathi.Seperator)+lroot
	}
}

func (pathi *Info) IsValid() (bool, string) {

	// has it been parsed?
	// do we have a result?
	return pathi.IsParsed()
}

func (pathi *Info) StringEquals(s string) bool {
	if s != "" {
		info := NewInfo(s)
		return pathi.Equals(info)
	}
	return false
}

func (pathi *Info) Equals(info *Info) bool {
	if info != nil {
		return strings.ToLower(pathi.String()) == strings.ToLower(info.String())
	}
	return false
}

func (pathi *Info) IsParsed() (bool, string) {
	if pathi == nil {
		return false, ""
	}
	return (pathi.ParsedPath != ""), pathi.String()
}

func (pathi *Info) String() string {
	if pathi == nil {
		return ""
	}
	return pathi.ParsedPath
}
