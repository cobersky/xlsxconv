package xlsxconv

import (
	"strconv"
	"strings"
	"math"
	"encoding/json"
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
func tryObjectVal(str string)interface {}{
	m:=map[string]interface {}{}
	if json.Unmarshal([]byte(str),&m)==nil{
		return m
	}else{
		return str
	}
}
func (this *Cell) GetValue(lib *Lib) interface{} {
	if this.value == nil {
		if this.IsString() {
			if index, err := strconv.Atoi(this.V); err == nil && len(lib.Strings) > index {
				v := lib.Strings[index].String()
				if strings.Index(v, "|") > 0 {
					sArr := strings.Split(v, "|")
					arr := []interface{}{}
					for i := 0; i < len(sArr); i++ {
						str := sArr[i]
						if len(str) > 0 {
							if num, err := strconv.ParseFloat(str, 64); err != nil {
								arr = append(arr, tryObjectVal(str))
							}else {
								if num == math.Floor(num) {
									arr = append(arr, int(num))
								}else {
									arr = append(arr, num)
								}
							}
						}
					}
					this.value = arr
				}else {
					this.value=tryObjectVal(v)
				}
			}else {
				this.value = this.V
			}
		}else {
			if num, err := strconv.ParseFloat(this.V, 64); err != nil {
				this.value = this.V
			}else {
				if num == math.Floor(num) {
					this.value = int(num)
				}else {
					this.value = num
				}
			}
		}
	}
	return this.value
}
