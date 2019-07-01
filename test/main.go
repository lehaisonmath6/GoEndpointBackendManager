package main

import (
	backendmanager "github.com/OpenStars/GoEndpointBackendManager"
)

func main() {
	enpManager := backendmanager.NewEndPointManager("http://127.0.0.1:2379", "/openstars/services/api/zshare/zsmetadataservice")
	enpManager.LoadEndpoint()
}
