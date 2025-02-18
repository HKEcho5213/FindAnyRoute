package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func JavaRoute(javafiles []string, InfoTrue bool) ([]string, []string) {

	//定义用于匹配类级和方法级的 @RequestMapping 注解的正则表达式
	//class_pattern := regexp.MustCompile(`@RequestMapping\s*\(\s*{\s*"([^"]+)"\s*}\s*\)`)
	//@RestController @RequestMapping({"/3.0/"})
	RestController_pattern1 := regexp.MustCompile(`@RestController\s+@RequestMapping\s*\(\s*{\s*"([^"]*)"\s*}\s*\)`)
	//@RestController@RequestMapping("/api")
	RestController_pattern2 := regexp.MustCompile(`@RestController\s*[\r\n]*@RequestMapping\(\s*"([^"]+)"\s*\)`)

	//@RequestMapping( value = {"/internal/v1/machine/addModel"}, method = {RequestMethod.*} )
	RequestMapping_value_pattern := regexp.MustCompile(`@RequestMapping\s*\(\s*value\s*=\s*{\s*"([^"]*)"\s*}\s*,\s*method\s*=\s*{\s*([^}]*)\s*}\s*\)`)

	//@RequestMapping({"/3.0/systemMaintenanceService/downloadApply"})
	Request_Mapping_pattern1 := regexp.MustCompile(`@RequestMapping\s*\(\s*{\s*"([^"]*)"\s*}\s*\)`)
	//@RequestMapping("/3.0/systemMaintenanceService/updateLicense")
	Request_Mapping_pattern2 := regexp.MustCompile(`@RequestMapping\s*\(\s*"([^"]+)"\s*\)`)

	//@GetMapping("/demo")
	Get_Mapping_pattern1 := regexp.MustCompile(`@GetMapping\s*\(\s*"([^"]+)"\s*\)`)
	//@GetMapping( 	value = {"/messageService/addNotice"}, 	method = {RequestMethod.GET} )
	Get_Mapping_pattern2 := regexp.MustCompile(`@GetMapping\s*\(\s*value\s*=\s*{\s*"([^"]*)"\s*}\s*,\s*method\s*=\s*{\s*([^}]*)\s*}\s*\)`)

	//@PostMapping("/create")
	Post_Mapping_pattern1 := regexp.MustCompile(`@PostMapping\s*\(\s*"([^"]*)"\s*\)`)
	//@PostMapping(value = {"/3.0/systemMaintenanceService/updateLicense"},method = {RequestMethod.POST})
	Post_Mapping_pattern2 := regexp.MustCompile(`@PostMapping\s*\(\s*value\s*=\s*{\s*"([^"]*)"\s*}\s*,\s*method\s*=\s*{\s*([^}]*)\s*}\s*\)`)

	var route_lists []string // 创建一个空切片,动态添加元素

	for i := range javafiles {
		//获取文件名（从路径中提取）,判断文件名是否以 ".java" 结尾
		if strings.HasSuffix(filepath.Base(javafiles[i]), ".java") {
			content, err := ioutil.ReadFile(javafiles[i])
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(string(content))
			// 提取@RestController @RequestMapping({"/3.0/"})类级别的路径
			//class_match := class_pattern.FindAllStringSubmatch(string(content), -1) //classMatch的类型是 [][]string，使用 FindAllStringSubmatch 进行多次匹配时，返回的是一个二维切片
			class_match1 := RestController_pattern1.FindStringSubmatch(string(content))
			class_match2 := RestController_pattern2.FindStringSubmatch(string(content))
			// 获取匹配到的路径，如果没有匹配到，返回空字符串
			//fmt.Println(reflect.TypeOf(class_match))
			var basePath string
			if len(class_match1) > 1 {
				basePath = class_match1[1]
			} else if len(class_match2) > 1 {
				basePath = class_match2[1]
			} else {
				basePath = ""
			}

			//fmt.Printf("Base path: %s\n", basePath)
			// 使用正则表达式匹配@RequestMapping( value = {"/internal/v1/machine/addModel"}, method = {RequestMethod.*} ) 注解的值，并将结果拼接为完整的接口路径
			RequestMapping_value_matches := RequestMapping_value_pattern.FindAllStringSubmatch(string(content), -1)
			fmt.Println("RequestMapping_value_method注解正则匹配中......")
			//fmt.Println(RequestMapping_value_matches)
			for _, requestmapping_value_matche := range RequestMapping_value_matches {
				api_path := basePath + requestmapping_value_matche[1]
				api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
				methods := strings.Split(requestmapping_value_matche[2], ".")[1]
				fmt.Println("【" + methods + "】 " + api_path)
				route_lists = append(route_lists, api_path)
			}

			// 使用正则表达式匹配 @RequestMapping({"/3.0/systemMaintenanceService/downloadApply"}) 注解的值，并将结果拼接为完整的接口路径
			Request_Mapping_matches1 := Request_Mapping_pattern1.FindAllStringSubmatch(string(content), -1)
			Request_Mapping_matches2 := Request_Mapping_pattern2.FindAllStringSubmatch(string(content), -1)
			fmt.Println("@RequestMapping注解正则匹配中......")
			if len(Request_Mapping_matches1) > 1 {
				for _, request_mapping_matche1 := range Request_Mapping_matches1 {
					if request_mapping_matche1[1] != basePath { //避免@RestController @RequestMapping再次匹配后输入结果
						api_path := basePath + request_mapping_matche1[1]
						api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
						fmt.Println("【 API_PATH 】 " + api_path)
						route_lists = append(route_lists, api_path)
					}
				}
			} else if len(Request_Mapping_matches2) > 1 {
				for _, request_mapping_matche2 := range Request_Mapping_matches2 {
					if request_mapping_matche2[1] != basePath { //避免@RestController @RequestMapping再次匹配后输入结果
						api_path := basePath + request_mapping_matche2[1]
						api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
						fmt.Println("【 API_PATH 】 " + api_path)
						route_lists = append(route_lists, api_path)
					}
				}
			}

			// 使用正则表达式匹配 @GetMapping("/demo") 注解的值，并将结果拼接为完整的接口路径
			Get_Mapping_matches1 := Get_Mapping_pattern1.FindAllStringSubmatch(string(content), -1)
			Get_Mapping_matches2 := Get_Mapping_pattern2.FindAllStringSubmatch(string(content), -1)
			fmt.Println("@GetMapping注解正则匹配中......")
			//fmt.Println("@GetMapping注解正则匹配中......")
			if len(Get_Mapping_matches1) > 1 {
				for _, get_mapping_matche1 := range Get_Mapping_matches1 {
					api_path := basePath + get_mapping_matche1[1]
					api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
					fmt.Println("【 GET 】 " + api_path)
					route_lists = append(route_lists, api_path)
				}
			} else if len(Get_Mapping_matches2) > 1 {
				for _, get_mapping_matche2 := range Get_Mapping_matches2 {
					api_path := basePath + get_mapping_matche2[1]
					api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
					fmt.Println("【 GET 】 " + api_path)
					route_lists = append(route_lists, api_path)
				}
			}

			// 使用正则表达式匹配 @PostMapping 注解的值，并将结果拼接为完整的接口路径
			post_mapping_matches1 := Post_Mapping_pattern1.FindAllStringSubmatch(string(content), -1)
			post_mapping_matches2 := Post_Mapping_pattern2.FindAllStringSubmatch(string(content), -1)
			fmt.Println("@PostMapping注解正则匹配中......")
			if len(post_mapping_matches1) > 1 {
				for _, post_mapping_matche1 := range post_mapping_matches1 {
					api_path := basePath + post_mapping_matche1[1]
					api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
					fmt.Println("【 POST 】 " + api_path)
					route_lists = append(route_lists, api_path)
				}
			} else if len(post_mapping_matches2) > 1 {
				for _, post_mapping_matche2 := range post_mapping_matches2 {
					api_path := basePath + post_mapping_matche2[1]
					api_path = strings.ReplaceAll(api_path, "//", "/") //替换路径中的//为/
					fmt.Println("【 POST 】 " + api_path)
					route_lists = append(route_lists, api_path)
				}
			}
			if InfoTrue {
				Informationlist = Sensitive_Information(string(content))
			}
		}
	}
	return route_lists, Informationlist
}
