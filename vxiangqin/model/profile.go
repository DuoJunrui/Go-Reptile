package model

type Profile struct {
	UserId               string `json:"user_id"`
	Name                 string `json:"name"`
	Sex                  string `json:"sex"`
	Age                  string `json:"age"`
	Marriage             string `json:"marriage"`
	Height               string `json:"height"`
	Weight               string `json:"weight"`
	Income               string `json:"income"`
	Education            string `json:"education"`
	Occupation           string `json:"occupation"`
	House                string `json:"house"`
	Car                  string `json:"car"`
	ExpectedMarriageDate string `json:"expected_marriage_date"`
}
