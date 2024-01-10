package setting

func GetExcelFullUrl(name string) string {
	return GetExcelPath() + name
}

func GetExcelPath() string {
	return ExportSavePath
}
