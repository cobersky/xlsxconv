package xlsxconv

import "regexp"

var cellColReg,_ =regexp.Compile(`^[a-zA-Z]+`)
var numReg,_=regexp.Compile(`^\-?\d+(\.\d+)?$`)
var fieldReg,_=regexp.Compile(`(\w+)\/([sca])`)
var sheetScanReg,_=regexp.Compile(`^xl\/worksheets\/sheet(\d+)\.xml$`)
var sheetReg,_ = regexp.Compile(`\|(\w+)`)
var wordReg,_=regexp.Compile(`[A-Za-z]+`)
