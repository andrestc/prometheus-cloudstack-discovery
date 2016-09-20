package cloudstack

type ListVirtualMachinesResponse struct {
	ListVirtualMachinesResponse struct {
		VirtualMachine []VirtualMachine `json:"virtualmachine"`
	} `json:"listvirtualmachinesresponse"`
}

type VirtualMachine struct {
	Nic []NicStruct `json:"nic"`
}

type NicStruct struct {
	IpAddress string `json:"ipaddress"`
}

type ListProjectsResponse struct {
	ListProjectsResponse struct {
		Project []Project `json:"project"`
	} `json:"listprojectsresponse"`
}

type Project struct {
	Id string
}
