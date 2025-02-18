package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

// 黑名单
// 黑名单正则表达式
var ExRouteRegexs = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\.(exp|vue|js|svg|jpg|png|css|jpeg|gif|mp3|wav|flac|mp4|mov|avi|mkv|bmp|tiff|pdf|ppt|odt|ods|odp|ps|eps|ai|psd|indd|fla|swf)$`),
	regexp.MustCompile(`^text\/(vnd.|coffeescript|x-coffeescript|x-gss|x-jade|css|plain|javascript|xml|asp|html|asa|h323|x-component|webviewhtml|x-styl|x-scss|x-handlebars-template|x-pug|x-sh|x-sass|typescript|ecmascript|x-less)$`),
	regexp.MustCompile(`^image\/(gif|jpeg|png|tiff|x-icon)$`),
	regexp.MustCompile(`^application\/(vnd.|vnd.wap.mms-message|java|x-sh|typescript|x-json|ecmascript|x-javascript|javascript|xhtml+xml|xml|atom+xml|json|pdf|msword|octet-stream|x-www-form-urlencoded|x-001|x-301|x-906|x-a11|postscript|x-anv|vnd.adobe.workflow|x-bmp|x-bot|x-c4t|x-c90|x-cals|vnd.ms-pki.seccat|x-netcdf|x-cdr|x-cel|x-x509-ca-cert|x-cgm|x-cit|x-cmp|x-cmx|x-cot|pkix-crl|x-dib|x-msdownload|x-drw|x-ebx|x-emf|x-epi|x-ps|fractals|x-frm|x-g4|x-gbr|x-gl2|x-gp4|x-hgl|x-hmr|x-hpgl|x-hpl|mac-binhex40|x-hrf|hta|x-icb|x-ico|x-iff|x-iphone|x-img|x-internet-signup)$`),
	regexp.MustCompile(`^multipart/form-data`),
	regexp.MustCompile(`^audio\/(x-mei-aac|aiff|basic|L8|x-wav)$`),
	regexp.MustCompile(`^video\/(x-ms-asf|avi)$`),
	regexp.MustCompile(`^drawing/907`),
	regexp.MustCompile(`^type/format`),
	regexp.MustCompile(`^java/.*`),
	regexp.MustCompile(`^message/rfc822`),
	regexp.MustCompile(`^script/x-vue`),
	regexp.MustCompile(`^M/D/yy`),
	regexp.MustCompile(`^https?://`),
	regexp.MustCompile(`@`),
	regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`),
	regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$`),
	regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\-(\d+)$`),
	regexp.MustCompile(`=`),
	regexp.MustCompile(`）`),
	regexp.MustCompile(`\)`),
	regexp.MustCompile(`#`),
	regexp.MustCompile(`^validation(/.*)?$`),
}

var route_lists []string // 创建一个空切片,动态添加元素

// 将切片转换为 map
func sliceToMap(slice []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range slice {
		m[v] = struct{}{} // struct{}{} 占用零内存
	}
	return m
}

// 检查路径是否匹配黑名单中的任何正则表达式
func isBlacklisted(path string) bool {
	for _, blacklistRegex := range ExRouteRegexs {
		if blacklistRegex.MatchString(path) {
			return true
		}
	}
	return false
}

func AbsolutePathRoute(source_path string, project_path string) {
	//绝对路径路由
	// 你可以在这个切片中添加所有需要检查的后缀
	extensions := []string{".jsp", ".jspx", ".properties", ".asp", ".aspx", ".ashx", ".ascx", ".asmx", ".php", ".xml", ".html", ".cshtml", ".vbhtml"}
	// 遍历所有后缀，检查路径是否以某个后缀结尾
	for _, ext := range extensions {
		if strings.HasSuffix(source_path, ext) {
			relativePath, err := filepath.Rel(project_path, source_path) //获取相对路径
			// 在路径前面加上"/"
			relativePath = "/" + relativePath
			// 替换反斜杠为正斜杠
			relativePath = strings.Replace(relativePath, "\\", "/", -1)
			if err != nil {
				log.Fatal(err)
			}
			if len(relativePath) > 6 {
				m_route_lists := sliceToMap(route_lists)              // 将切片转换为 map
				if _, exists := m_route_lists[relativePath]; exists { //判断切片中是否已经存在路径，决定是否存入切片
				} else {
					relativePath = strings.ReplaceAll(relativePath, "//", "/") //替换路径中的//为/
					fmt.Println("JsRouteScan匹配到的路径:", relativePath)
					route_lists = append(route_lists, relativePath)
				}
			}
		}
	}
}

func Unexpected_information(fileContent string) {
	// 正则表达式
	re := regexp.MustCompile(`[\"|'](/[0-9a-zA-Z.]+(?:/[\\w,\\?,-,\\.,_]*?)+)[\"|']`)
	// 查找所有匹配的内容
	matches := re.FindAllStringSubmatch(fileContent, -1)
	m_route_lists := sliceToMap(route_lists) // 将切片转换为 map
	// 输出匹配结果
	for _, match := range matches {
		if len(match[1]) > 6 {
			if isValidUTF8(match[1]) { //判断是否为乱码
				if !isBlacklisted(match[1]) {
					// 在路径前面加上"/",并将"//"替换为"/"
					relativePath := "/" + match[1]
					relativePath = strings.ReplaceAll(relativePath, "webapps", "")
					relativePath = strings.ReplaceAll(relativePath, "//", "/")
					relativePath = sanitizeURLPath(relativePath) // 替换URL路径中的'和"字符为空
					// 移除控制字符后的字符串
					relativePath = removeControlChars(relativePath)
					if containsSpace(relativePath) == false { //判断路径中是否存在空格
						if startsWithDirPrefix(relativePath) == false { //判断文件中是否以linux目录开头，为True则直接跳过
							if !isBlacklisted(relativePath) { //去除黑名单
								if _, exists := m_route_lists[relativePath]; exists { //判断切片中是否已经存在路径，决定是否存入切片
								} else {
									relativePath = strings.ReplaceAll(relativePath, "//", "/") //替换路径中的//为/
									fmt.Println("Unexpected_information匹配到的路径:", relativePath)
									route_lists = append(route_lists, relativePath)
								}
							}
						}
					}
				}
			}
		}
	}
}

func JsRouteScan(fileContent string) {
	re := regexp.MustCompile(`.{10}["'` + "`" + `]([a-zA-Z0-9/=_{}\.\?&!-]+/[a-zA-Z0-9/=_{}\.\?&!-]+(\.jspx|\.jsp|\.html|\.php|\.do|\.aspx|\.action|\.json)*)["'` + "`" + `].{160}`)
	matches := re.FindAllStringSubmatch(fileContent, -1)
	m_route_lists := sliceToMap(route_lists) // 将切片转换为 map
	for _, match := range matches {
		// match[1] 是我们感兴趣的捕获组（路径）
		if len(match[1]) > 6 {
			if isValidUTF8(match[1]) { //判断是否为乱码
				if !isBlacklisted(match[1]) { //去除黑名单
					relativePath := "/" + match[1]
					relativePath = strings.ReplaceAll(relativePath, "//", "/")
					relativePath = sanitizeURLPath(relativePath) // 替换URL路径中的'和"字符为空
					// 移除控制字符后的字符串
					relativePath = removeControlChars(relativePath)
					if containsSpace(relativePath) == false { //判断路径中是否存在空格
						if startsWithDirPrefix(relativePath) == false { //判断文件中是否以linux目录开头，为True则直接跳过
							if !isBlacklisted(relativePath) { //去除黑名单
								if _, exists := m_route_lists[relativePath]; exists { //判断切片中是否已经存在路径，决定是否存入切片
								} else {
									relativePath = strings.ReplaceAll(relativePath, "//", "/") //替换路径中的//为/
									fmt.Println("JsRouteScan匹配到的路径:", relativePath)
									route_lists = append(route_lists, relativePath)
								}
							}
						}
					}
				}
			}
		}
	}
}

func FilterJs(fileContent string) {
	re := regexp.MustCompile(`["']([a-zA-Z0-9/=_{}\?&!:\.-]+/[a-zA-Z0-9/=_{}\?&!:\.-]+(\.jspx|\.jsp|\.html|\.php|\.do|\.aspx|\.action|\.json)*)["']`)
	matches := re.FindAllStringSubmatch(fileContent, -1)
	m_route_lists := sliceToMap(route_lists) // 将切片转换为 map
	for _, match := range matches {
		// match[1] 是我们感兴趣的捕获组（路径）
		if len(match[1]) > 6 {
			if isValidUTF8(match[1]) { //判断是否为乱码
				if !isBlacklisted(match[1]) { //去除黑名单
					relativePath := "/" + match[1]
					relativePath = strings.ReplaceAll(relativePath, "//", "/")
					relativePath = sanitizeURLPath(relativePath) // 替换URL路径中的'和"字符为空
					// 移除控制字符后的字符串
					relativePath = removeControlChars(relativePath)
					if containsSpace(relativePath) == false { //判断路径中是否存在空格
						if startsWithDirPrefix(relativePath) == false { //判断文件中是否以linux目录开头，为True则直接跳过
							if !isBlacklisted(relativePath) { //去除黑名单
								if _, exists := m_route_lists[relativePath]; exists { //判断切片中是否已经存在路径，决定是否存入切片
								} else {
									relativePath = strings.ReplaceAll(relativePath, "//", "/") //替换路径中的//为/
									fmt.Println("FilterJs匹配到的路径:", relativePath)
									route_lists = append(route_lists, relativePath)
								}
							}
						}
					}
				}
			}
		}
	}
}

func HaE(fileContent string) {
	re := regexp.MustCompile(`(?:"|')(((?:[a-zA-Z]{1,10}://|//)[^"'/]{1,}\.[a-zA-Z]{2,}[^"']{0,})|((?:/|\.\./|\./)[^"'><,;|*()(%%$^/\\\[\]][^"'><,;|()]{1,})|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{1,}\.(?:[a-zA-Z]{1,4}|action)(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{3,}(?:[\?|#][^"|']{0,}|))|([a-zA-Z0-9_\-]{1,}\.(?:\w)(?:[\?|#][^"|']{0,}|)))(?:"|')`)
	matches := re.FindAllStringSubmatch(fileContent, -1)
	m_route_lists := sliceToMap(route_lists) // 将切片转换为 map
	for _, match := range matches {
		// match[1] 是我们感兴趣的捕获组（路径）
		if len(match[1]) > 6 {
			if isValidUTF8(match[1]) { //判断是否为乱码
				if !isBlacklisted(match[1]) { //去除黑名单
					relativePath := "/" + match[1]                             //在路径前加上/
					relativePath = strings.ReplaceAll(relativePath, "//", "/") //替换路径中的//为/
					relativePath = sanitizeURLPath(relativePath)               // 替换URL路径中的'和"字符为空
					// 移除控制字符后的字符串
					relativePath = removeControlChars(relativePath)
					if containsSpace(relativePath) == false { //判断路径中是否存在空格
						if startsWithDirPrefix(relativePath) == false { //判断文件中是否以linux目录开头，为True则直接跳过
							if !isBlacklisted(relativePath) { //去除黑名单
								if _, exists := m_route_lists[relativePath]; exists { //判断切片中是否已经存在路径，决定是否存入切片
								} else {
									relativePath = strings.ReplaceAll(relativePath, "//", "/") //替换路径中的//为/
									fmt.Println("HaE匹配到的路径:", relativePath)
									route_lists = append(route_lists, relativePath)
								}
							}
						}
					}
				}
			}
		}
	}
}

func ExtractFileRoute(javafiles []string, project_path string, InfoTrue bool) ([]string, []string) {
	for i := range javafiles {
		//获取绝对路径路由
		AbsolutePathRoute(javafiles[i], project_path)
		content, err := ioutil.ReadFile(javafiles[i]) //读取单个文件内容
		if err != nil {
			log.Fatal(err)
		}
		Unexpected_information(string(content))
		JsRouteScan(string(content))
		FilterJs(string(content))
		HaE(string(content))
		if InfoTrue {
			Informationlist = Sensitive_Information(string(content))
		}
	}
	return route_lists, Informationlist
}
