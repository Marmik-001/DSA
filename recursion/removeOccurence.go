package recursion

import ( 
	"strings"
)

func RemoveAllOccurenceOfASubstring(str string,  subStr string) string{ 



	if strings.Contains(str, subStr){
		str = RemoveAllOccurenceOfASubstring(strings.Replace(str ,subStr,  "\000" , 1) , subStr)
	}

	return str

}