package models

type Information struct {
	Os            string `json:"os"`
	KernelName    string `json:"kernelName"`
	HostName      string `json:"hostName"`
	KernelRelease string `json:"kernelRelease"`
	KernelVersion string `json:"kernelVersion"`
	Machine       string `json:"machine"`
	Processor     string `json:"processor"`
	HwPlatform    string `json:"hwPlatform"`
	UsedSpace     string `json:"usedSpace"`
	DateTime      string `json:"dateTime"`
}
