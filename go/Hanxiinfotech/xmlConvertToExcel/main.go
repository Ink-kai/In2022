package xmlConvertToExcel

import (
	"flag"
	"fmt"
	"path"
)

//主函数
func Run() {
	var f FileDispose
	flag.StringVar(&f.ExportFileName, "name", "temp", "文件保存名")
	flag.StringVar(&f.FilePath, "f", "G:\\test\\OutLookMenu.xml", "文件路径（xml）。例：G:\\test\\OutLookMenu.xml")
	flag.StringVar(&f.F_type, "type", "1", "文件处理格式。默认1-xls	2-csv")
	flag.Parse()
	/*
		异常捕捉
	*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic错误：\t%v", r)
		}
	}()
	if extName := path.Ext(f.FilePath); extName == ".xml" {
		//1.提取xml数据
		var xmldataer = f.XmlDispose()
		//2.数据导出
		switch f.F_type {
		case "2":
			//写入csv
			xmldataer.WriteCsv(f.ExportFileName)
		case "1":
			//写入xlsx
			xmldataer.WriteXlsx(f.ExportFileName)
		case "3":
			//写入xls
		}
	} else {
		//1.提取xml数据
		var xmldataer = f.JsonDispose()
		//2.数据导出
		switch f.F_type {
		case "2":
			//写入csv
			xmldataer.WriteCsv(f.ExportFileName)
		case "1":
			//写入xlsx
			xmldataer.WriteXlsx(f.ExportFileName)
		case "3":
			//写入xls
		}
		panic("只支持xml文件。json导出 数据格式需自行调整")
	}

}
