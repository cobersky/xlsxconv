package xlsxconv

type Head map[string]Field

func NewHead(row *Row, lib *Lib) *Head {
	h := Head{}
	for _, v := range row.Cells {
		str, ok := v.GetValue(lib).(string)
		if ok{
			if matches := fieldReg.FindStringSubmatch(str); len(matches) > 0 {
				f := Field{}
				f.Name = matches[1]
				switch matches[2]{
				case "a":
					f.ExportClient = true
					f.ExportServer = true
				case "s":
					f.ExportServer = true
				case "c":
					f.ExportClient = true
				}
				h[wordReg.FindString(v.R)] = f
			}
		}
	}
	return &h
}

type Field struct {
	ExportClient bool
	ExportServer bool
	Name         string
}
