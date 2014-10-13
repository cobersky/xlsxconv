package xlsxconv

import (
	"archive/zip"
	"io"
	"errors"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

var cm = map[string]interface{}{}
var sm = map[string]interface{}{}

type WorkBook struct {
	SheetsInfo []SheetInfo`xml:"sheets>sheet"`
	Sheets map[string]*Sheet`xml:"-"`
	Lib *Lib`xml:"-"`
}

type Lib struct {
	Strings []LibString`xml:"si"`
}
type LibString struct {
	T string`xml:"t"`
	RT []string`xml:"r>t"`
}
func (this LibString)String()string{
	if this.T!=""{
		return this.T
	}
	return strings.Join(this.RT,"")
}
type SheetInfo struct {
	Name string`xml:"name,attr"`
	Id   string`xml:"sheetId,attr"`
}

func (this *WorkBook) ToMap() {
	for k, v := range this.Sheets {
		c, s := v.ToArray(this.Lib)
		cm[k] = c
		sm[k] = s
	}
}

func Parse(filePath string) {
	sx := simpleParse(filePath)
	workbook := &WorkBook{}
	workbook.Lib = &Lib{}
	workbook.Sheets = map[string]*Sheet{}
	parseXML(sx.WorkBook, workbook)
	parseXML(sx.SharedStrings, workbook.Lib)
	for k, v := range workbook.SheetsInfo {
		if matches := sheetReg.FindStringSubmatch(v.Name); len(matches) > 0 {
			s := &Sheet{}
			err:=parseXML(sx.Sheets[strconv.Itoa(k+1)], s)
			if err!=nil{
				fmt.Println(err)
			}
			workbook.Sheets[matches[1]] = s
		}
	}
	workbook.ToMap()
}
func GetClientMap() map[string]interface{} {
	return cm
}
func GetServerMap() map[string]interface{} {
	return sm
}

type SimpleXlsx struct {
	WorkBook *zip.File
	SharedStrings *zip.File
	Sheets map[string]*zip.File
}

const (
	workbookPath = "xl/workbook.xml"
	libPath      = "xl/sharedStrings.xml"
)

func simpleParse(filePath string) *SimpleXlsx {
	if reader, err := zip.OpenReader(filePath); err != nil {
		return nil
	}else {
		x := &SimpleXlsx{}
		x.Sheets = map[string]*zip.File{}
		for _, f := range reader.File {
			if f.Name == workbookPath {
				x.WorkBook = f
			}else if f.Name == libPath {
				x.SharedStrings = f
			}else {
				if matches := sheetScanReg.FindStringSubmatch(f.Name); len(matches) > 0 {
					x.Sheets[matches[1]] = f
				}
			}
		}
		return x
	}
}
func parseXML(file *zip.File, v interface{}) error {
	if b := getXMLData(file); b != nil {
		return xml.Unmarshal(b, v)
	}
	return errors.New("unknown error!")
}

func getXMLData(file *zip.File) []byte {
	if file != nil {
		if reader, err := file.Open(); err != nil {
			return nil
		}else {
			b := make([]byte, file.UncompressedSize64)
			if _, err := io.ReadFull(reader, b); err != nil {
				return nil
			}else {
				return b
			}
		}
	}
	return nil
}
