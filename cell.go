package xlsxconv

import (
	"strconv"
	"strings"
)

type Cell struct {
	R     string`xml:"r,attr"`
	T     string `xml:"t,attr"`
	V     string `xml:"v"`
	value interface{}
}

func (this Cell) IsString() bool {
	return this.T == "s"
}
func (this *Cell) GetValue(lib *Lib) interface{} {
	if this.value == nil {
		if this.IsString() {
			if index, err := strconv.Atoi(this.V); err == nil && len(lib.Strings) > index {
				v:=lib.Strings[index]
				if strings.Index(v, "|") > 0 {
					sArr:=strings.Split(v, "|")
					arr:=[]interface {}{}
					for i:=0;i<len(sArr);i++{
						str:=sArr[i]
						if num, err := strconv.ParseFloat(str, 64); err != nil {
							arr=append(arr,str)
						}else{
							arr=append(arr,num)
						}
					}
					this.value = arr
				}else {
					this.value = v
				}
			}else {
				this.value=this.V
			}
		}else {
			if num, err := strconv.ParseFloat(this.V, 64); err != nil {
				this.value = this.V
			}else {
				this.value = num
			}
		}
	}
	return this.value
}
