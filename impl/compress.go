package impl

import (
	"log"
	"os/exec"
)

func CompressGz()  {
	cmd := "wc -l ./cloudtrail/logs/*.gz; if [[ $? == 0  ]];then gunzip ./cloudtrail/logs/*; fi"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}
