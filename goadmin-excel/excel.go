package goadmin_excel

import "github.com/tealeg/xlsx"

type ExportXlsx struct {
	Excel *xlsx.File
}

func NewExportXlsx() *ExportXlsx {
	return &ExportXlsx{Excel: xlsx.NewFile()}
}

func (ex *ExportXlsx) AddSheet(name string) *xlsx.Sheet {
	sheet, _ := ex.Excel.AddSheet(name)
	return sheet
}

func (ex *ExportXlsx) AddRow(sheet *xlsx.Sheet, style *xlsx.Style, options ...string) *xlsx.Row {
	row := sheet.AddRow()
	for _, option := range options {
		cell := row.AddCell()
		cell.SetString(option)
		cell.SetStyle(style)
	}
	return row
}

func (ex *ExportXlsx) MergeRow(sheet *xlsx.Sheet, style *xlsx.Style, mergeLine int, options ...string) *xlsx.Row {
	row := sheet.AddRow()
	for index, option := range options {
		cell := row.AddCell()
		cell.SetStyle(style)
		if index != 2 {
			cell.Merge(0, mergeLine)
		}
		cell.SetString(option)
	}
	return row
}

func (ex *ExportXlsx) OnlyUserRow(sheet *xlsx.Sheet, user string, style *xlsx.Style) *xlsx.Row {
	row := sheet.AddRow()

	row.AddCell().SetStyle(style)
	row.AddCell().SetStyle(style)

	cell31 := row.AddCell()
	cell31.SetString(user)
	cell31.SetStyle(style)

	row.AddCell().SetStyle(style)
	return row
}

func (ex *ExportXlsx) AddEmptyCell(row *xlsx.Row, style *xlsx.Style, count int) *xlsx.Row {
	for i := 0; i < count; i++ {
		cell := row.AddCell()
		cell.SetStyle(style)
	}
	return row
}
