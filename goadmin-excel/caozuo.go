package goadmin_excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func ExportExcel() *xlsx.File {
	border := xlsx.Border{
		Left:        "thin",
		LeftColor:   "5B9BD5",
		Right:       "thin",
		RightColor:  "5B9BD5",
		Top:         "thin",
		TopColor:    "5B9BD5",
		Bottom:      "thin",
		BottomColor: "5B9BD5",
	}

	titleStyle := xlsx.NewStyle()
	titleStyle.Alignment.Horizontal = "Center"
	titleStyle.Alignment.Vertical = "Center"
	titleStyle.Border = border
	titleStyle.Font.Bold = true

	bgStyle := xlsx.NewStyle()
	bgStyle.Border = border
	bgStyle.Fill = *xlsx.NewFill("solid", "DDEBF7", "DDEBF7")

	bdStyle := xlsx.NewStyle()
	bdStyle.Border = border
	bdStyle.Alignment.Horizontal = "left"
	bdStyle.Alignment.Vertical = "top"
	bdStyle.ApplyFill = true

	footerStyle := xlsx.NewStyle()
	footerStyle.Alignment.Horizontal = "Center"
	footerStyle.Alignment.Vertical = "Center"
	footerStyle.Border = border

	exportXlsx := NewExportXlsx()
	sheet := exportXlsx.AddSheet("sheet")

	titleRow := sheet.AddRow()
	titleCell := titleRow.AddCell()
	titleCell.Merge(3, 0)
	titleCell.SetString("大标题")
	titleCell.SetStyle(titleStyle)
	exportXlsx.AddEmptyCell(titleRow, titleStyle, 3)

	exportXlsx.AddRow(sheet, bgStyle, "标题一", "标题二", "标题三", "占比")

	exportXlsx.MergeRow(sheet, bdStyle, 1, "选项一", "xxx", "张三", "28.99%")
	exportXlsx.OnlyUserRow(sheet, "李四", bdStyle)

	exportXlsx.MergeRow(sheet, bdStyle, 2, "选项二", "xxx", "xxx", "28.99%")
	exportXlsx.OnlyUserRow(sheet, "xxx", bdStyle)
	exportXlsx.OnlyUserRow(sheet, "xxx", bdStyle)

	exportXlsx.MergeRow(sheet, bdStyle, 1, "选项三", "xxx", "bbb", "28.99%")
	exportXlsx.OnlyUserRow(sheet, "bbb", bdStyle)

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetString("表脚")
	cell.SetStyle(footerStyle)
	cell.Merge(3, 0)
	exportXlsx.AddEmptyCell(row, footerStyle, 3)

	return exportXlsx.Excel
}

func main() {
	file := ExportExcel()
	err := file.Save("test.xlsx")
	if err != nil {
		fmt.Println("failed")
		return
	}
	fmt.Println("successful")
}
