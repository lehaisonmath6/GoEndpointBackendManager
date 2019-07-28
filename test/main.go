package main

import (
	"fmt"
	"log"

	backendmanager "github.com/OpenStars/GoEndpointBackendManager"
)

func main() {
	enpManager := backendmanager.NewEndPointManager("http://127.0.0.1:2379", "/openstars/services/api/zshare/zsmetadataservice")
	enpManager.LoadEndpoint()
	err, lsEnpoints := enpManager.GetEndPoints()
	if err != nil {
		log.Println(err.Error())
	}
	CompactEndpoint := new(backendmanager.EndPoint)
	for i := 0; i < len(lsEnpoints); i++ {
		if lsEnpoints[i].Type == backendmanager.EThriftCompact {
			CompactEndpoint = lsEnpoints[i]
		}
	}
	fmt.Println("Compact host : ", CompactEndpoint.Host, " , port : ", CompactEndpoint.Port)
	log.Println(lsEnpoints)
}
