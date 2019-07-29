package main

import (
	"log"

	backendmanager "github.com/lehaisonmath6/GoEndpointBackendManager"
)

func ObChangeEndpoint(ep *backendmanager.EndPoint) {
	log.Println(ep)
	// set update new config o day
}

func main() {
	enpManager := backendmanager.NewEndPointManager("http://127.0.0.1:2379", "/dir1")
	enpManager.TestConnectEtcdServer()
	enpManager.EventChangeEndPoints(ObChangeEndpoint)
	for {

	}
	// enpManager.LoadEndpoint()
	// err, ep := enpManager.GetEndPointType(backendmanager.EThriftCompact)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// log.Println("host:", ep.Host, "; port:", ep.Port)
	// err, lsEnpoints := enpManager.GetEndPoints(typ)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// CompactEndpoint := new(backendmanager.EndPoint)
	// for i := 0; i < len(lsEnpoints); i++ {
	// 	log.Println(lsEnpoints[i])
	// 	if lsEnpoints[i].Type == backendmanager.EThriftCompact {
	// 		CompactEndpoint = lsEnpoints[i]
	// 	}
	// }
	// fmt.Println("Compact host : ", CompactEndpoint.Host, " , port : ", CompactEndpoint.Port)
}
