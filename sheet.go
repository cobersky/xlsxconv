package xlsxconv

type Sheet struct {
	Rows []*Row`xml:"sheetData>row"`
}

func (this *Sheet) ToArray(lib *Lib)(cArr,sArr []interface{}) {
	if l := len(this.Rows); l > 1 {
		head:=NewHead(this.Rows[0],lib)
		cArr=[]interface {}{}
		sArr=[]interface {}{}
		for i:=1;i<l;i++{
			c,s:=this.Rows[i].ToMap(head,lib)
			if len(c)>0{
				cArr=append(cArr,c)
			}
			if len(s)>0{
				sArr=append(sArr,s)
			}
		}
	}
	return
}
