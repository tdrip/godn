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
	return NewInfoCustom(string(Seperator)+"", Seperator, strings.TrimSpace(Path))
}

func NewInfoCustom(Top string, Seperator byte, Path string) *Info {

	info := MakeDefaultInfo()

	// store the orginal value passed in
	info.OriginalValue = Path

	// set the seperator
	info.Seperator = Seperator

	// set the top of the path
	info.Top = info.parseTop(Top)

	// Parse the Path
	info = info.parsePath(Path, info.Seperator, info.Top)

	//build the name
	info = info.buildName()

	// Mark as Parse
	info.Parsed = true

	return info
}

//MakeDefaultInfo only creates a default Info object
func MakeDefaultInfo() *Info {
	// build object
	info := new(Info)

	// set an empty top
	info.Top = info.parseTop(DefaultTop)

	// set the defult seperator
	info.Seperator = DefaultSeperator

	// mark as unparsed
	info.Parsed = false

	return info
}

func (pathi *Info) buildName() *Info {

	if len(pathi.ParsedPath) > 0 {
		r := []rune(pathi.ParsedPath)

		pathi.Name = ""
		pathi.Parent = nil

		//fmt.Println("Name: ", pathi.Name)
		//fmt.Println("Parent: ", pathi.Parent)

		foundfirstseperator := false

		//ppath := strings.Replace(pathi.ParsedPath, string(pathi.Seperator), "", -1)
		//ptop := strings.Replace(pathi.Top, string(pathi.Seperator), "", -1)
		ppath := strings.TrimSuffix(pathi.ParsedPath, string(pathi.Seperator))
		ptop := strings.TrimSuffix(pathi.Top, string(pathi.Seperator))

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
					parent := string(r[0 : i+1])
					//fmt.Printf("parent : %s \n", parent)

					pathi.Parent = NewInfoCustom(ptop, pathi.Seperator, parent)

					//fmt.Printf("i : %d \n", i)
					//fmt.Printf("len : %d \n", len(r))
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

	return pathi
}

// clean the path
func (pathi *Info) parsePath(Path string, Seperator byte, Top string) *Info {

	//fmt.Println("Top: ", Top)
	//fmt.Println("Path: ", Path)

	if len(Path) > 0 {

		// clear up any white spaces
		path := strings.TrimSpace(Path)
		//fmt.Printf("Path = Trimmed: %s \n", path)

		// remove any double seperators with single seperators
		// for example \\ or //
		path = strings.Replace(path, string(Seperator)+string(Seperator), string(Seperator), -1)
		//fmt.Printf("Path = Replace %c%c: %s \n", Seperator, Seperator, path)

		// it's not null, has one character and that character is a seperator
		// for example: \ or /
		if len(path) == 1 && path[0] == Seperator {
			path = Top
		}

		// chop off the end seperator if it is provided
		if len(path) > 1 && strings.HasSuffix(path, string(Seperator)) {
			path = path[:(len(path) - 1)]
			//fmt.Printf("Path = Path[:(len(Path) - 1)]: %s \n", path)
		}

		// make sure that the start has a slash
		if len(path) > 0 && path[0] != Seperator {
			path = string(Seperator) + path
			//fmt.Printf("Path = string(Seperator) + Path: %s \n", path)
		}

		// Is the Path toped?
		if !strings.HasPrefix(strings.ToLower(path), strings.ToLower(Top)) {
			path = pathi.addTop(path, Seperator, Top)
		}

		// fix double slashes being provided
		// clean up if we were passed a Path with extra slashes (JSON)
		pathi.ParsedPath = pathi.addTop(path, Seperator, Top)
	} else {
		pathi.ParsedPath = pathi.addTop(Path, Seperator, Top)
	}

	//fmt.Printf("pathi.ParsedPath: %s \n", pathi.ParsedPath)

	return pathi
}

func (pathi *Info) addTop(Path string, Seperator byte, Top string) string {

	pdn := strings.Replace(Path, string(Seperator), "", -1)
	ptop := strings.Replace(Top, string(Seperator), "", -1)

	////fmt.Println("Top: ", ptop)
	////fmt.Println("Path: ", pdn)

	if strings.ToLower(pdn) == strings.ToLower(ptop) {
		return Path
	}

	// two checks on the top
	tophasfirstslash := (Top[0] == Seperator)
	tophasendslash := (Top[len(Top)-1] == Seperator)

	// split the Top
	r := []rune(Top)

	// we have two seperators
	// |foo|, \loo\ etc
	if tophasendslash && tophasfirstslash {
		Path = string(r[0:len(Top)-1]) + Path
		//fmt.Printf("Path =  string(r[0:len(Root) - 1]) + Path: %s \n", Path)
	} else {

		if tophasfirstslash && !tophasendslash {
			// we have no end seperators
			// \foo, \loo etc
			// has seperator

			Path = string(Seperator) + string(r[0:len(Top)-1]) + Path
			//fmt.Printf("Path =  string(Seperator) + string(r[0:len(Top) - 1]) + Path: %s \n", Path)
		} else if !tophasfirstslash && tophasendslash {

			// we have no start seperators
			// \foo, \loo etc
			// has seperator
			Path = string(r[0:len(Top)-1]) + Path
			//fmt.Printf("Path =  string(r[0:len(Top) - 1]) + Path: %s \n", Path)
		} else {
			Path = string(Seperator) + Top + Path
			//fmt.Printf("Path =  string(Seperator) + Top + Path: %s \n", Path)
		}
	}

	return Path

}

///
func (pathi *Info) parseTop(top string) string {

	//fmt.Printf("ParseTop: %s \n", top)
	Top := ""
	if len(top) == 0 {
		// Default
		Top = string(pathi.Seperator)
	} else {

		if top[0] != pathi.Seperator && top[len(top)-1] != pathi.Seperator {
			Top = string(pathi.Seperator) + top + string(pathi.Seperator)
			//fmt.Printf("pathi.Root = string(pathi.Seperator) + top + string(pathi.Seperator): %s \n", pathi.Top)
		} else {

			if top[0] == pathi.Seperator && top[len(top)-1] != pathi.Seperator {
				Top = top + string(pathi.Seperator)
				//fmt.Printf("top + string(pathi.Seperator): %s \n", pathi.Top)
			} else if top[0] != pathi.Seperator && top[len(top)-1] == pathi.Seperator {
				Top = string(pathi.Seperator) + top
				//fmt.Printf("pathi.Root = string(pathi.Seperator) + top: %s \n", pathi.Top)
			} else {
				Top = top
				//fmt.Printf("pathi.Root = top: %s \n", pathi.Top)
			}
		}
	}

	//fmt.Printf("pathi.Top: %s \n", pathi.Top)

	return Top
}

func (pathi *Info) IsTop() bool {
	if pathi == nil {
		return false
	}

	parsed, path := pathi.IsParsed()

	if parsed && pathi.Parent == nil {
		return true
	}

	ltop := strings.ToLower(pathi.Top)
	lpath := strings.ToLower(path)
	//fmt.Printf("lpath: %s \n", lpath)
	//fmt.Printf("ltop: %s \n", ltop)
	if len(lpath) == 1 {

		return lpath[0] == pathi.Seperator
	}
	return lpath == string(pathi.Seperator)+ltop+string(pathi.Seperator) || lpath == ltop+string(pathi.Seperator) || lpath == string(pathi.Seperator)+ltop

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
	return (len(pathi.ParsedPath) > 0), pathi.String()
}

func (pathi *Info) String() string {
	if pathi == nil {
		return ""
	}
	return pathi.ParsedPath
}
