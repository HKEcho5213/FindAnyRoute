package main

import (
	"findroute/utils"
	"flag"
	"fmt"
	"path/filepath"
	"strconv"
)

func main() {
	fmt.Println("FindAnyRoute.exe -p PATH/Project -l java -i True")
	fmt.Println("Author: HKEcho")
	fmt.Println("Github: https://github.com/HKEcho5213")
	fmt.Println("注: .class,jar包以及.NET项目需自行反编译")

	var project_path, language string
	var InfoTrue string

	// 使用 flag 包解析命令行参数
	flag.StringVar(&project_path, "p", "", "项目路径; eg: PATH/Project")
	flag.StringVar(&language, "l", "", "代码语言; eg: java,all ...")
	flag.StringVar(&InfoTrue, "i", "Flase", "敏感信息搜寻,搜寻时间较长，请耐心等待, 默认搜寻username,password,AK/SK等; eg: True/Flase")

	// 解析命令行参数
	flag.Parse()
	if project_path != "" {
		//处理目录文件
		arb_filespath := utils.Pathprocessing(project_path)
		resultfilename := filepath.Base(project_path) //获取项目文件名
		InfoTrue, _ := strconv.ParseBool(InfoTrue)
		fmt.Println("指定的项目路径:", project_path)
		if language == "java" {
			fmt.Println("指定的方式是:", language)
			routelists, Information := utils.JavaRoute(arb_filespath, InfoTrue)
			route := utils.RemoveDuplicates(routelists)           //去重
			utils.Save_Result(route, resultfilename+"_Route.txt") //保存路由到项目文件
			if len(Information) > 0 {
				utils.Save_Result(Information, resultfilename+"_Information.txt") //保存路由到项目文件
			}
		} else if language == "all" {
			fmt.Println("指定的方式是:", language)
			routelists_java, Information := utils.JavaRoute(arb_filespath, InfoTrue)
			routelists, Information := utils.ExtractFileRoute(arb_filespath, project_path, InfoTrue)
			// 合并切片
			All_Routelists := append(routelists, routelists_java...)   //使用 append(slice1, slice2...) 将 slice2 中的所有元素追加到 slice1 的末尾。这里的 slice2... 是 Go 的切片展开语法，将 slice2 展开为一个个独立的元素传递给 append
			routelists = utils.RemoveDuplicates(All_Routelists)        //去重
			utils.Save_Result(routelists, resultfilename+"_Route.txt") //保存路由到项目文件
			if len(Information) > 0 {
				utils.Save_Result(Information, resultfilename+"_Information.txt") //保存路由到项目文件
			}
		} else {
			fmt.Println("未指定代码语言或指定错误,请通过'-l'指定代码语言")
		}
	} else {
		fmt.Println("未指定项目路径,请通过'-p'指定项目路径")
	}

}
