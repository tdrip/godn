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
	Top string

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

func NewInfo(Path string) *Info {
	return NewInfoCustom(DefaultTop, DefaultSeperator, strings.TrimSpace(Path))
}

func NewInfoCustomTop(Top string, Path string) *Info {
	return NewInfoCustom(Top, DefaultSeperator, strings.TrimSpace(Path))
}

func NewInfoCustomSeperator(Seperator byte, Path string) *Info {
	return NewInfoCustom(DefaultTop, Seperator, strings.TrimSpace(Path))
}

func NewInfoCustom(Top string, Seperator byte, Path string) *Info {

	info := MakeDefaultInfo()

	// store the orginal value passed in
	info.OriginalValue = Path

	// set the seperator
	info.Seperator = Seperator

	// set the top of the path
	info = info.parseTop(Top)

	// Parse the Path
	info.ParsedPath = info.parsePath(Path, info.Seperator, info.Top)

	//build the name
	info.Name = info.buildName()

	// Mark as Parse
	info.Parsed = true

	return info
}

//MakeDefaultInfo only creates a default Info object
func MakeDefaultInfo() *Info {
	// build object
	info := new(Info)

	// set an empty top
	info = info.parseTop(DefaultTop)

	// set the defult seperator
	info.Seperator = DefaultSeperator

	// mark as unparsed
	info.Parsed = false

	return info
}

func (pathi *Info) buildName() string {

	if len(pathi.ParsedPath) > 0 {

		r := []rune(pathi.ParsedPath)

		pathi.Name = ""
		pathi.Parent = nil

		//ParentPath := ""
		PTop := pathi.Top
		//fmt.Println("Name: ", pathi.Name)
		//fmt.Println("Parent: ", pathi.Parent)
		foundfirstseperator := false

		// fix later
		ppath := strings.Replace(pathi.Parent.String(), string(pathi.Seperator), "", -1)
		ptop := strings.Replace(PTop, string(pathi.Seperator), "", -1)

		//fmt.Println("Parent Top: ", ptop)
		//fmt.Println("Parent Path: ", ppath)

		if strings.ToLower(ppath) == strings.ToLower(ptop) {
			pathi.Name = ppath
			pathi.Parent = nil
		} else {
			// walk in reverse over the Path
			for i := len(r) - 1; i >= 0; i-- {

				if foundfirstseperator {

					// the first
					pathi.Name = string(r[i+2 : len(r)])
					pathi.Parent = NewInfoCustom(ptop, pathi.Seperator, string(r[0:i+1]))

					//fmt.Println("i : %d ", i);
					//fmt.Println("len : %d ",len(r));
					break
				}

				// we found
				if r[i] == rune(pathi.Seperator) {
					foundfirstseperator = true
				}
			}
		}
		//fmt.Println("Name: ", pathi.Name)
		//fmt.Println("Parent: ", pathi.Parent)
	}

	return ""
}

// clean the path
func (pathi *Info) parsePath(Path string, Seperator byte, Top string) string {

	//fmt.Println("Top: ", Top)
	//fmt.Println("Path: ", Path)

	if len(Path) > 0 {

		// clear up any white spaces
		Path = strings.TrimSpace(Path)

		Path = strings.Replace(Path, string(Seperator)+string(Seperator), string(Seperator), -1)

		// it's not null, has one character and that character is a seperator
		// \\ or //
		if len(Path) > 0 && len(Path) == 1 && Path[0] == Seperator {
			Path = Top
		}

		// chop off the end slash if it is provided
		if strings.HasSuffix(Path, string(Seperator)) {
			Path = Path[:(len(Path) - 1)]
			//fmt.Println("Path  = Path[:(len(Path) - 1)]: ", Path)
		}

		// make sure that the start has a slash
		if len(Path) > 0 && Path[0] != Seperator {
			Path = string(Seperator) + Path
			//fmt.Println("Path = string(Seperator) + Path: ", Path)
		}

		// Is the Path toped?
		if !strings.HasPrefix(strings.ToLower(Path), strings.ToLower(Top)) {
			Path = pathi.addTop(Path, Seperator, Top)
		}

		// fix double slashes being provided
		// clean up if we were passed a Path with extra slashes (JSON)
		return Path
	} else {
		Path = pathi.addTop(Path, Seperator, Top)
	}

	return Path
}

func (pathi *Info) addTop(Path string, Seperator byte, Top string) string {

	pdn := strings.Replace(Path, string(Seperator), "", -1)
	ptop := strings.Replace(Top, string(Seperator), "", -1)

	//fmt.Println("Top: ", ptop)
	//fmt.Println("Path: ", pdn)

	if strings.ToLower(pdn) == strings.ToLower(ptop) {
		return Path
	} else {

		// two checks on the top
		tophasfirstslash := (Top[0] == Seperator)
		tophasendslash := (Top[len(Top)-1] == Seperator)

		// split the Top
		r := []rune(Top)

		// we have two seperators
		// |foo|, \loo\ etc
		if tophasendslash && tophasfirstslash {
			Path = string(r[0:len(Top)-1]) + Path
			//fmt.Println("Path =  string(r[0:len(Root) - 1]) + Path: ", Path)
		} else {

			if tophasfirstslash && !tophasendslash {
				// we have no end seperators
				// \foo, \loo etc
				// has seperator

				Path = string(Seperator) + string(r[0:len(Top)-1]) + Path
				//fmt.Println("Path =  string(Seperator) + string(r[0:len(Top) - 1]) + Path: ", Path)
			} else if !tophasfirstslash && tophasendslash {

				// we have no start seperators
				// \foo, \loo etc
				// has seperator
				Path = string(r[0:len(Top)-1]) + Path
				//fmt.Println("Path =  string(r[0:len(Top) - 1]) + Path: ", Path)
			} else {
				Path = string(Seperator) + Top + Path
				//fmt.Println("Path =  string(Seperator) + Top + Path: ", Path)
			}
		}

		return Path
	}
}

///
func (pathi *Info) parseTop(top string) *Info {

	//fmt.Println("ParseTop: ", top)
	Top := ""
	if len(top) == 0 {
		// Default
		Top = string(pathi.Seperator)
	} else {

		if top[0] != pathi.Seperator && top[len(top)-1] != pathi.Seperator {
			Top = string(pathi.Seperator) + top + string(pathi.Seperator)
			//fmt.Println("pathi.Root = string(pathi.Seperator) + top + string(pathi.Seperator): ", pathi.Root)
		} else {

			if top[0] == pathi.Seperator && top[len(top)-1] != pathi.Seperator {
				Top = top + string(pathi.Seperator)
				//fmt.Println("top + string(pathi.Seperator): ", pathi.Root)
			} else if top[0] != pathi.Seperator && top[len(top)-1] == pathi.Seperator {
				Top = string(pathi.Seperator) + top
				//fmt.Println("pathi.Root = string(pathi.Seperator) + top: ", pathi.Root)
			} else {
				Top = top
				//fmt.Println("pathi.Root = top: ", pathi.Root)
			}
		}
	}

	if Top == "" {

	} else {

	}
	//fmt.Println("pathi.Top: ", pathi.Top)

	return pathi

}

func (pathi *Info) IsTop() bool {
	if pathi == nil {
		return false
	} else {
		ltop := strings.ToLower(pathi.Top)
		lpath := strings.ToLower(pathi.String())
		if len(lpath) == 1 {
			return lpath[0] == pathi.Seperator
		}
		return lpath == string(pathi.Seperator)+ltop+string(pathi.Seperator) || lpath == ltop+string(pathi.Seperator) || lpath == string(pathi.Seperator)+ltop
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
