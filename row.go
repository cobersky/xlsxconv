package xlsxconv

type Row struct {
	R     string`xml:"r,attr"`
	Cells []*Cell`xml:"c"`
}

func (this *Row) ToMap(head *Head, lib *Lib) (c, s map[string]interface{}) {
	c = map[string]interface{}{}
	s = map[string]interface{}{}
	rl := len(this.R)
	for _, v := range this.Cells {
		k := v.R[:len(v.R)-rl]
		if f, ok := (*head)[k]; ok {
			if f.ExportClient {
				c[f.Name] = v.GetValue(lib)
			}
			if f.ExportServer {
				s[f.Name] = v.GetValue(lib)
			}
		}
	}
	return
}
