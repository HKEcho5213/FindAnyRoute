package utils

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func Pathprocessing(project_path string) []string {
	var arb_filespath []string // 创建一个空切片,动态添加元素

	// 使用 Walk 函数递归遍历目录
	err := filepath.Walk(project_path, func(path string, info fs.FileInfo, err error) error {
		//处理遍历过程中的错误
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		// 检查当前路径是否包含 "maven" 文件夹
		if info.IsDir() && strings.Contains(path, "maven") { //排除"maven"目录下文件
			// 跳过 maven 目录及其子目录
			return filepath.SkipDir
		}
		// 如果是文件，则输出文件名,只处理文件，忽略目录
		if !info.IsDir() {
			// 获取文件的绝对路径
			arb_files, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			// 将绝对路径加入到结果切片中
			arb_filespath = append(arb_filespath, arb_files)
		}
		return nil
	})
	if err != nil {
		fmt.Println("解析项目路径出错:", err)
	}
	return arb_filespath
}
