package entity

type Injection struct {
	Version   string `json:"version"`
	BuildTime string `json:"buildTime"`
	Commit    string `json:"commit"`
	GoVersion string `json:"goVersion"`
}
