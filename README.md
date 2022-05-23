# in2022

记录 2022 年代码

## 03.19

1.初始化仓库，添加.gitignore 文件  
2.go 和 vue 项目

## 03.20

~~1.搭建 vue+ant 后台管理系统~~  
~~2.go 写二分查找算法，开始学算法~~

## 04.01

1.梳理 GO 工作目录  
2.学习 GO 读取文件的方式

## 04.04

#### 新增 Blog 项目

-   项目工作目录初始化,
-   User 表建立

## 04.12

用 go 写文件替换的程序

## 05.11

## 新增 linux 目录

centos 系统安装的初始化脚本

## 05.15

#### 调整项目结构
~~go 目录新增 MySQLStore 项目。~~
~~初步完成文件二进制上传接口，待完善。~~
curl -X POST -F "files=@文件物理路径" -H "Content-Type: multipart/form-data" http://localhost:8080/api/uploadFile
例：curl -X POST -F "files=@D:\\waljgjl.jpg" -F "files=@D:\\walwgh.jpg" -H "Content-Type: multipart/form-data" http://localhost:8080/api/uploadFile

## 05.23

+ 完善go编写的附件上传口
+ 开始学习GO单元测试