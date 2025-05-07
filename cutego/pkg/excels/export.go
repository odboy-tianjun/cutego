package excels

import (
	"github.com/druidcaesa/gotool"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

// 获取title头list
func getTitle(list []map[string]string) []string {
	titleList := make([]string, 0)
	for _, item := range list {
		for _, s := range item {
			titleList = append(titleList, s)
		}
	}
	return titleList
}

// ExportExcel 导出excel
func ExportExcel(list []interface{}, title string) (error, *excelize.File) {
	// 获取标题
	headerList, expList := ExcelCreate(list[0])
	headers := getTitle(headerList)
	// 默认存在第一个工作簿是 Sheet1 首字母要大写, 否则会报错。
	// 如果想额外的创建工作簿, 可以使用, sheet2 := file.NewSheet("Sheet2"), 工作簿的名称不区分大小写。
	// 如果有多个工作簿, 可以使用 file.SetActiveSheet(index) 来指定打开文件时focus到哪个工作簿
	sheet1 := "Sheet1"
	files := excelize.NewFile()
	character := string(65 + len(headers) - 1)
	/* -------------------- 第一行大标题 -------------------- */
	// 设置行高
	err := files.SetRowHeight(sheet1, 1, 25)
	if err != nil {
		return err, nil
	}
	// 合并单元格
	err = files.MergeCell(sheet1, "A1", character+"1")
	if err != nil {
		return err, nil
	}
	// 设置单元格样式：对齐；字体, 大小；单元格边框
	styleTitle, _ := files.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true,"italic":false,"family":"Calibri","size":16,"color":"#000000"},"border":[{"type":"left","color":"#3FAD08","style":0},{"type":"top","color":"#3FAD08","style":0},{"type":"bottom","color":"#3FAD08","style":2},{"type":"right","color":"#3FAD08","style":0}]}`)
	err = files.SetCellStyle(sheet1, "A1", character+"1", styleTitle)
	if err != nil {
		return err, nil
	}

	err = files.SetCellValue(sheet1, "A1", title)
	if err != nil {
		return err, nil
	}

	/* -------------------- 字段标题 -------------------- */
	styleHeader, _ := files.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":false,"italic":false,"family":"Calibri","size":10,"color":"#000000"}}`)
	err = files.SetCellStyle(sheet1, "A2", character+"2", styleHeader)
	if err != nil {
		return err, nil
	}
	for k, v := range headers {
		err = files.SetCellValue(sheet1, string(65+k)+"2", v)
		if err != nil {
			return err, nil
		}
	}
	// 设置最后一列宽度
	err = files.SetColWidth(sheet1, "C", character, 20)
	if err != nil {
		return err, nil
	}
	// 冻结窗口：冻结第一行和第二行
	err = files.SetPanes(sheet1, `{"freeze":true,"split":false,"x_split":0,"y_split":2}`)
	if err != nil {
		return err, nil
	}
	///* -------------------- 填充行数据 -------------------- */
	line := 3
	for _, v := range list {
		var num = 0
		lineChr := strconv.Itoa(line)
		// 设置样式
		err = files.SetCellStyle(sheet1, "A"+lineChr, character+lineChr, styleHeader)
		if err != nil {
			return err, nil
		}
		// 反射获取数据和类型
		getValue := reflect.ValueOf(v)
		getType := reflect.TypeOf(v)
		n := getValue.NumField()
		for i := 0; i < n; i++ {
			val := getValue.Field(i)
			name := getType.Field(i).Name
			if !getIsTitle(name, headerList) {
				continue
			}
			err = files.SetCellValue(sheet1, string(65+num)+lineChr, getExp(name, expList, val.Interface()))
			if err != nil {
				return err, nil
			}
			num++
		}
		line++
	}
	return nil, files
}

func getIsTitle(name string, headerList []map[string]string) bool {
	flag := false
	for _, m := range headerList {
		if gotool.StrUtils.HasNotEmpty(m[name]) {
			flag = true
		}
	}
	return flag
}

func getExp(name string, expList []map[string][]map[string]string, value interface{}) interface{} {
	for _, m := range expList {
		if len(m[name]) > 0 {
			maps := m[name]
			for _, m2 := range maps {
				if gotool.StrUtils.HasNotEmpty(m2[interfaceToString(value)]) {
					value = m2[interfaceToString(value)]
				}
			}
		}
	}
	return value
}

func interfaceToString(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int:
		return strconv.Itoa(inter.(int))
	case int64:
		return strconv.FormatInt(inter.(int64), 10)
	default:
		return inter.(string)
	}
}
