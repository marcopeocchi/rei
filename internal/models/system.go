package models

type SystemTop struct {
	CPU       string `json:"cpu"`
	CoreCount int32  `json:"coreCount"`
	Hostname  string `json:"hostname"`
	OS        string `json:"os"`
	Platform  string `json:"platform"`
	Uptime    uint64 `json:"uptime"`
	CPUTemp   string `json:"cpuTemp"`
	RAMFree   uint64 `json:"ramFree"`
}

type SystemTemp struct {
	CPUTemp string `json:"cpuTemp"`
}
