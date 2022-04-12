package example

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type F_Info struct {
	fileName string
}

type Filer interface {
	ReadFileV1()
	ReadFileV2()
	ReadFileV3()
	ReadBytes()
	ReadBytes()
	ReadString()
	Os_ReadFileContent()
}

func New(file_name string) Filer {
	file := F_Info{
		fileName: file_name,
	}
	return &file
}

/*
	一、指定文件名读取
	1.使用os.ReadFile
*/
func (f *F_Info) ReadFileV1() {
	content, err := os.ReadFile(f.fileName)
	if err != nil {
		panic(f)
	}
	fmt.Fprintf(os.Stdout, "%s\n-----------------------------------", content)
}

/*
	2.使用ioutil.ReadFile
*/
func (f *F_Info) ReadFileV2() {
	content, err := ioutil.ReadFile(f.fileName)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "\n%s\n-----------------------------------", content)
}

/*
	二、先创建句柄再读取
*/
func (f *F_Info) ReadFileV3() {
	file, err := os.Open(f.fileName)
	// file, err := os.OpenFile(f.fileName, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Fprintf(os.Stdout, "\n%s\n-----------------------------------", content)
}

/*
	三、每次只读取一行
	bufio.ReadLine()
	bufio.ReadBytes(':raw-latex:`\n`')
	bufio.ReadString(':raw-latex:`\n`')
*/
func (f *F_Info) ReadBytes() {
	fi, err := os.Open(f.fileName)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "\n%s", line)
	}
}

func (f *F_Info) ReadString() {
	fi, err := os.Open(f.fileName)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(string(line))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "\n%s", line)
	}
}

/*
	每次按字节读取(这里取的是1024字节)
*/
func (f *F_Info) Os_ReadFileContent() {
	fi, err := os.Open(f.fileName)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)
	buf := make([]byte, 1024)
	for {
		content, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "\n%s,%v", buf[:content], content)
	}
}
