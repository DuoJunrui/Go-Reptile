package model

type Profile struct {
	Data `json:"data"`
}

type Data struct {
	Id         int    `json:"memberID"`
	Name       string `json:"nickname"`
	Gender     string `json:"genderString"`
	Age        int    `json:"age"`
	Height     string `json:"heightString"`
	Weight     string
	Income     string `json:"salaryString"`
	Marriage   string `json:"marriageString"`
	Education  string `json:"educationString"`
	Occupation string
	Jiguan     string
	WorkCity   string `json:"workProvinceCityString"`
	Xinzuo     string
	House      string
	Car        string
}
