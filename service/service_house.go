package service


func ChinaWords(str string)(strsql string)  {
	words :=[]rune(str)
	strsql = ""
	for i:=0 ;i<len(words);i++{
		strsql = strsql + "%" +string(words[i])
	}
	strsql = strsql + "%"
	return  strsql
}

