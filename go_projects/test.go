package main

import (
	"package/example"
)

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
	path := "./example_file/t.ini"
	file := example.New(path)
	file.ReadFileV1()
	file.ReadFileV2()
	file.ReadFileV3()
	file.ReadBytes()
	file.ReadString()
	file.Os_ReadFileContent()
}
