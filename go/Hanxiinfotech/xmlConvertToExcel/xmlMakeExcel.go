package xmlConvertToExcel

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

//参数管理
type FileDispose struct {
	//传入xml文件路径
	FilePath string
	//保存文件名
	ExportFileName string
	// 文件处理格式
	F_type string
	//需要处理的节点
	All_nodes string
}

type (
	xml_form struct {
		//带结构标签，反引号来包围字符串
		//XMLName           xml.Name         `xml:"OutLookBarItemList"`
		OutLookBarItemStr []OutLookBarItem `xml:"OutLookBarItem"`
	}
	Children struct {
		//XMLName  string `xml:string`
		TextName       string `xml:"Text,attr"`
		CodeName       string `xml:"Code,attr"`
		UrlName        string `xml:"Url,attr"`
		ParentCodeName string `xml:"ParentCode,attr"`
		HasChildren    string `xml:"HasChildren,attr"`
	}
	OutLookBarItem struct {
		//XMLName  string `xml:string`
		TextName    string     `xml:"Text,attr"`
		CodeName    string     `xml:"Code,attr"`
		ChildrenStr []Children `xml:"Children"`
		HasChildren string     `xml:"HasChildren,attr"`
	}
)

type (
	json_children struct {
		Id       string          `json:"Id"`
		Text     string          `json:"Text"`
		Icon     string          `json:"icon"`
		View     string          `json:"View"`
		Leaf     bool            `json:"Leaf"`
		Children []json_children `json:"Children"`
	}
	json_form struct {
		Expanded bool            `json:"expanded"`
		Children []json_children `json:"children"`
	}
)

//xml对外接口
type Xml_former interface {
	WriteCsv(name string)
	WriteXlsx(name string)
}

//json对外接口
type Json_former interface {
	WriteCsv(name string)
	WriteXlsx(name string)
}

/*
提取xml数据
*/
func (f *FileDispose) XmlDispose() Xml_former {
	//打开指定得xml文档
	xmlFile, err := os.Open(f.FilePath)
	//检查是否有错误
	if err != nil {
		//如果打开文档发生错位直接退出
		fmt.Println("Error opening XML file!")
		panic(err)
	}
	//关闭文档  注意defer
	defer xmlFile.Close()
	//从xml文档中读取数据
	xmlData, err := ioutil.ReadAll(xmlFile)
	//错误判断
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		panic(err)
	}
	var data xml_form
	//将xml数据封装到结构体
	xml.Unmarshal(xmlData, &data)
	ob := data.OutLookBarItemStr
	if len(ob) < 1 {
		//err = errors.New("未检测到根节点<OutLookBarItem>")
		panic("未检测到根节点<OutLookBarItem>")
	}
	if len(ob[0].ChildrenStr) < 1 {
		//err = errors.New("未检测到子节点<Children>")
		panic("未检测到子节点<Children>")
	}
	return &data
}

/*
json数据处理
*/
func (f *FileDispose) JsonDispose() Json_former {
	data, err := ioutil.ReadFile(f.FilePath)
	if err != nil {
		panic(err)
	}
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
	var json_data json_form
	unmarshal_err := json.Unmarshal(data, &json_data)
	if unmarshal_err != nil {
		panic(unmarshal_err)
	}
	ob := json_data.Children
	if len(ob) < 1 {
		//err = errors.New("未检测到根节点<OutLookBarItem>")
		panic("未检测到根节点<Children>")
	}
	if len(ob[0].Children) < 1 {
		//err = errors.New("未检测到子节点<Children>")
		panic("未检测到子节点<Children>")
	}
	return &json_data
}

/*
保存至csv文件
*/
func (p *xml_form) WriteCsv(exportFileName string) {
	f, err := os.Create(exportFileName + ".csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.Write([]string{
		"菜单编码", "菜单名称", "路径", "链接地址", "上级菜单标识", "层级", "是否叶子节点", "排序", "打开方式", "对齐情况", "页面状态", "宽度", "高度", "高度占比", "宽度占比"})
	var e_i int = 1
	for _, e := range p.OutLookBarItemStr {
		//标记
		var i int = 1
		w.Write([]string{e.CodeName, e.TextName, "", "", "root", "1", e.HasChildren, strconv.Itoa(e_i), "0", "0", "0", "1200", "550", "PX", "PX"})
		for _, ch := range e.ChildrenStr {
			w.Write([]string{ch.CodeName, ch.TextName, ch.ParentCodeName, ch.UrlName, ch.ParentCodeName, "2", ch.HasChildren, strconv.Itoa(i), "0", "0", "0", "1200", "550", "PX", "PX"})
			i += 1
		}
		e_i += 1
	}

	w.Flush()
}

/*
保存至csv文件
*/
func (j *json_form) WriteCsv(exportFileName string) {
	f, err := os.Create(exportFileName + ".csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.Write([]string{
		"菜单编码", "菜单名称", "路径", "链接地址", "上级菜单标识", "层级", "是否叶子节点", "排序", "打开方式", "对齐情况", "页面状态", "宽度", "高度", "高度占比", "宽度占比"})
	var e_i int = 1
	for _, e := range j.Children {
		//标记
		var i int = 1
		leaf := strconv.FormatBool(e.Leaf)
		w.Write([]string{e.Id, e.Text, "", e.View, "root", "1", "=text(" + leaf + ",0)", "" + strconv.Itoa(e_i), "0", "0", "0", "1200", "550", "PX", "PX"})
		for _, ch := range e.Children {
			leaf = strconv.FormatBool(ch.Leaf)
			w.Write([]string{ch.Id, ch.Text, e.Id, ch.View, e.Id, "2", leaf, strconv.Itoa(i), "0", "0", "0", "1200", "550", "PX", "PX"})
			i += 1
		}
		e_i += 1
	}

	w.Flush()
}

/*
保存至xls文件
*/
func (p *xml_form) WriteXlsx(exportFileName string) {
	ch := make(chan []interface{})
	go func() {
		for _, v := range p.OutLookBarItemStr {
			ch <- []interface{}{
				excelize.Cell{Value: v.CodeName},
				excelize.Cell{Value: v.TextName},
				excelize.Cell{Value: ""},
				excelize.Cell{Value: ""},
				excelize.Cell{Value: "root"},
				excelize.Cell{Value: "1"},
				excelize.Cell{Value: "false"},
				excelize.Cell{Value: "1"},
				excelize.Cell{Value: "0"},
				excelize.Cell{Value: "0"},
				excelize.Cell{Value: "0"},
				excelize.Cell{Value: "1200"},
				excelize.Cell{Value: "550"},
				excelize.Cell{Value: "PX"},
				excelize.Cell{Value: "PX"},
			}
			for _, children := range v.ChildrenStr {
				ch <- []interface{}{
					excelize.Cell{Value: children.CodeName},
					excelize.Cell{Value: children.TextName},
					excelize.Cell{Value: children.ParentCodeName},
					excelize.Cell{Value: children.UrlName},
					excelize.Cell{Value: children.ParentCodeName},
					excelize.Cell{Value: "2"},
					excelize.Cell{Value: v.HasChildren},
					excelize.Cell{Value: "1"},
					excelize.Cell{Value: "0"},
					excelize.Cell{Value: "0"},
					excelize.Cell{Value: "0"},
					excelize.Cell{Value: "1200"},
					excelize.Cell{Value: "550"},
					excelize.Cell{Value: "PX"},
					excelize.Cell{Value: "PX"},
				}
			}
		}
		close(ch)
	}()
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		panic(err)
	}
	header := []string{
		"菜单编码", "菜单名称", "路径", "链接地址", "上级菜单标识", "层级", "是否叶子节点", "排序", "打开方式", "对齐情况", "页面状态", "宽度", "高度", "高度占比", "宽度占比"}
	header_data := make([]interface{}, len(header))
	for i, e := range header {
		header_data[i] = e
	}
	streamWriter.SetRow("A1", header_data)

	var i int = 1
	for val := range ch {
		i += 1
		a, _ := excelize.CoordinatesToCellName(1, i)
		streamWriter.SetRow(a, val)
	}
	if err := streamWriter.Flush(); err != nil {
		panic(err)
	}
	if err := file.SaveAs(exportFileName + ".xlsx"); err != nil {
		panic(err)
	}
}

/*
保存至xls文件
*/
func (j *json_form) WriteXlsx(exportFileName string) {
	ch := make(chan []interface{})
	go func() {
		for _, v := range j.Children {
			leaf := strconv.FormatBool(v.Leaf)
			ch <- []interface{}{
				excelize.Cell{Value: v.Id},
				excelize.Cell{Value: v.Text},
				excelize.Cell{Value: ""},
				excelize.Cell{Value: ""},
				excelize.Cell{Value: "root"},
				excelize.Cell{Value: "1"},
				excelize.Cell{Value: leaf},
				excelize.Cell{Value: "1"},
				excelize.Cell{Value: "0"},
				excelize.Cell{Value: "0"},
				excelize.Cell{Value: "0"},
				excelize.Cell{Value: "1200"},
				excelize.Cell{Value: "550"},
				excelize.Cell{Value: "PX"},
				excelize.Cell{Value: "PX"},
			}
			for _, children := range v.Children {
				leaf = strconv.FormatBool(children.Leaf)
				ch <- []interface{}{
					excelize.Cell{Value: children.Id},
					excelize.Cell{Value: children.Text},
					excelize.Cell{Value: v.Id},
					excelize.Cell{Value: children.View},
					excelize.Cell{Value: v.Id},
					excelize.Cell{Value: "2"},
					excelize.Cell{Value: leaf},
					excelize.Cell{Value: "1"},
					excelize.Cell{Value: "0"},
					excelize.Cell{Value: "0"},
					excelize.Cell{Value: "0"},
					excelize.Cell{Value: "1200"},
					excelize.Cell{Value: "550"},
					excelize.Cell{Value: "PX"},
					excelize.Cell{Value: "PX"},
				}
			}
		}
		close(ch)
	}()
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		panic(err)
	}
	header := []string{
		"菜单编码", "菜单名称", "路径", "链接地址", "上级菜单标识", "层级", "是否叶子节点", "排序", "打开方式", "对齐情况", "页面状态", "宽度", "高度", "高度占比", "宽度占比"}
	header_data := make([]interface{}, len(header))
	for i, e := range header {
		header_data[i] = e
	}
	streamWriter.SetRow("A1", header_data)

	var i int = 1
	for val := range ch {
		i += 1
		a, _ := excelize.CoordinatesToCellName(1, i)
		streamWriter.SetRow(a, val)
	}
	if err := streamWriter.Flush(); err != nil {
		panic(err)
	}
	if err := file.SaveAs(exportFileName + ".xlsx"); err != nil {
		panic(err)
	}
}
