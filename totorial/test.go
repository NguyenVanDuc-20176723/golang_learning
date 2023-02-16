package main

import (
	"encoding/json"
	"fmt"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

type SearchOpts struct {
	Name string
	Age  int
}

type Paginator struct {
	Limit  int
	Marker string
}

func (p *Paginator) show() {
	fmt.Println(p.Limit)
}

type ListOpts struct {
	SearchOpts
	Paginator
}

func main() {
	paramMap := map[string]interface{}{
		"age": 23, "name": "Duc",
	}
	jsonData, _ := json.Marshal(paramMap)
	data := ListOpts{SearchOpts{Name: "duc"}, Paginator{Limit: 1}}
	var structData ListOpts
	json.Unmarshal(jsonData, &structData)
	fmt.Println(structData, data)
	data.show()
	servers.ListOpts{}
}
