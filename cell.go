package xlsxconv

import "strconv"

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
				this.value = lib.Strings[index]
			}else {
				this.value = this.V
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
