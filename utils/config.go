package utils

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func RemoveDuplicates(routelists []string) []string {
	//针对切片去重函数
	// 创建一个 map 用于记录字符串是否已出现
	unique := make(map[string]bool)
	var result []string
	// 遍历原始切片
	for _, str := range routelists {
		// 如果 map 中没有这个字符串，则添加到结果切片，并标记为已出现
		if !unique[str] {
			unique[str] = true
			result = append(result, str)
		}
	}
	return result
}

func Save_Result(result []string, filename string) {
	// 打开文件，如果文件不存在则创建，如果文件存在则覆盖
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	content := strings.Join(result, "\n") //将字符串切片合并为一个带换行符的单一字符串,一次性写入
	// 写入文件
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	// 成功写入
	fmt.Println("File written successfully")
}

// 判断字符串是否符合 UTF-8 编码
func isValidUTF8(s string) bool {
	return utf8.Valid([]byte(s))
}

// 判断字符串中是否包含空格
func containsSpace(s string) bool {
	return strings.Contains(s, " ")
}

// 判断文件名是否以指定的目录前缀开头
func startsWithDirPrefix(filename string) bool {
	// 需要检查的目录前缀列表
	prefixes := []string{"project://", "java://", "abc://", "META-INF/", "file://", "/.ssh/", "/!", "/dd/MM/", "/+8/", "/SSH/", "/FTP/", "/Telnet/", "/bin/", "/boot/", "/CloudrResetPwdAgent/", "/dev/", "/etc/", "/home/", "/lib/", "/lib32/", "/lib64/", "/libx32/", "/lost+found/", "/media/", "/mnt/", "/opt/", "/proc/", "/root/", "/run/", "/sbin/", "/snap/", "/srv/", "/sys/", "/tmp/", "/usr/", "/var/"}
	// 遍历前缀列表
	for _, prefix := range prefixes {
		if strings.HasPrefix(filename, prefix) {
			return true
		}
	}
	return false
}

// 替换URL路径中的'和"字符为空
func sanitizeURLPath(urlPath string) string {
	// 替换单引号和双引号为空
	urlPath = strings.ReplaceAll(urlPath, "'", "")
	urlPath = strings.ReplaceAll(urlPath, "\"", "")
	urlPath = strings.ReplaceAll(urlPath, "!", "")
	urlPath = strings.ReplaceAll(urlPath, ":", "")
	urlPath = strings.ReplaceAll(urlPath, ":id", "1")
	urlPath = strings.ReplaceAll(urlPath, ":type", "1")
	urlPath = strings.ReplaceAll(urlPath, ":index", "1")
	urlPath = strings.ReplaceAll(urlPath, "./", "")
	urlPath = strings.ReplaceAll(urlPath, "......", "")
	urlPath = strings.ReplaceAll(urlPath, "\\", "")
	urlPath = strings.ReplaceAll(urlPath, "%0", "")
	return urlPath
}

// 清除字符串中的控制字符
func removeControlChars(s string) string {
	result := ""
	for _, r := range s {
		if !unicode.IsControl(r) {
			result += string(r)
		}
	}
	return result
}
