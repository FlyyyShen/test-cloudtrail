package impl

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"strings"
)

func ResJsonLoadFile(fullname string)  {
	file, err := ioutil.ReadFile(fullname)
	if err !=nil {fmt.Println(err)}

	newJson, _ := simplejson.NewJson(file)


	record,err := newJson.Get("Records").Array()
	for _, row := range record {
		if each_map, ok := row.(map[string]interface{}); ok {



			result := fmt.Sprintf("%v", each_map["eventName"])
			if ok := strings.Contains(result, "Create");ok {
				//fmt.Println(result)

				data, _ := json.Marshal(each_map["responseElements"])
				fmt.Printf("create事件: %s, %s\n", result, string(data))

			}



		}
	}
}


func GetCreateResource(pathname string)  {
	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n",pathname, err)

	}


	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取
		if fi.IsDir() {
			GetCreateResource(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v",fullname, err)
			}

		}
		ResJsonLoadFile(fullname)
	}
}
