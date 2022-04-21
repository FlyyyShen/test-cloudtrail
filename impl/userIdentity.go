package impl

import (
	"fmt"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)



var principalId_list = make([]string,0)
var principalId_map = map[string]int{}

func RemoveDuplicateElement(s []string)  []string{
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}

	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
			// 计数
			principalId_map[item] = 1
		}else {
			principalId_map[item] = principalId_map[item] + 1
		}
	}
	return result
}


func AppendprincipalIdlist(fullname string)  {
	file, err := ioutil.ReadFile(fullname)
	if err !=nil {fmt.Println(err)}

	newJson, _ := simplejson.NewJson(file)


	record,err := newJson.Get("Records").Array()
	for _, row := range record {
		if each_map, ok := row.(map[string]interface{}); ok {

			principalId := fmt.Sprintf("%v", each_map["userIdentity"].(map[string]interface{})["principalId"])

			principalId_list = append(principalId_list,principalId)

		}
	}


}

func GetprincipalId(pathname string)  {
	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n",pathname, err)

	}


	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取
		if fi.IsDir() {
			GetprincipalId(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v",fullname, err)
			}

		}
		AppendprincipalIdlist(fullname)
	}
	RemoveDuplicateElement(principalId_list)

	for k,v := range principalId_map {
		fmt.Printf("principalId: %v, 次数: %v\n",k, v)
	}
}

