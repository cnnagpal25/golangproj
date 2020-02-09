package testlib

var myLang map[string]string

func init() {
	myLang = make(map[string]string)
	myLang["chirag"] = "golang"
	myLang["reema"] = ".net"
}

/*
GetLang : For getting language according to name
*/
func GetLang(key string) string {
	return myLang[key]
}
