package models

type SystemTop struct {
	CPU       string `json:"cpu"`
	CoreCount int32  `json:"coreCount"`
	Hostname  string `json:"hostname"`
	OS        string `json:"os"`
	Platform  string `json:"platform"`
	Uptime    uint64 `json:"uptime"`
	CPUTemp   string `json:"cpuTemp"`
	MemFree   uint64 `json:"memFree"`
	MemTotal  uint64 `json:"memTotal"`
}

type SystemTemp struct {
	CPUTemp string `json:"cpuTemp"`
}
