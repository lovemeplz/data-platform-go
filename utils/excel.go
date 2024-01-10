package utils

import (
	"fmt"
	"github.com/lovemeplz/data-platform-go/pkg/setting"
	"github.com/xuri/excelize/v2"
)

func GenerateExcel() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件

	path := setting.GetExcelFullUrl("test.xlsx")
	fmt.Println("path:::", path)
	if err := f.SaveAs("/runtime/export/test.xlsx"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("succeed")
}
