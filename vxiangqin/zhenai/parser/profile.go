package parser

import (
	"Go-Reptile/multitask/engine"
	"Go-Reptile/multitask/fetcher"
	"Go-Reptile/multitask/model"
	"encoding/json"
	"regexp"
)

var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([0-9]+kg)</div>`)
var jiguanRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:(.*?)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^>]*>*座[^<]+)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+房)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+车)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:.*?<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

func ParseProfile(contents []byte, id, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	/*用正则表达式解析html代码*/
	profile.Data.Weight = extractString(contents, weightRe)
	profile.Data.Jiguan = extractString(contents, jiguanRe)
	profile.Data.Xinzuo = extractString(contents, xinzuoRe)
	profile.Data.House = extractString(contents, houseRe)
	profile.Data.Car = extractString(contents, carRe)
	profile.Data.Occupation = extractString(contents, occupationRe)

	/*解析JSON*/
	profileUrl := `https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=` + id + `&_=1649045597876&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2Fd9d80d68-4b98-4f71-abd1-a36cc12ac6ed%2F0%2F0%2F1980405550&data=eyJ2IjoiSEJOTWgrMTM3Ym9qMkJQL0NaZS93Zz09Iiwib3MiOiJ3ZWIiLCJpdCI6MzI3MiwicyI6NCwiZCI6IntcInBhcnRuZXJcIjpcInpoZW5haVwiLFwiYXBwX25hbWVcIjpcInBjUHJpbnRcIixcInRva2VuX2lkXCI6XCJ6aGVuYWktMTY0OTA0NTE1NzY3MC05ZTY2MDU5NjcyMWMxXCIsXCJhXCI6XCJkSjZhUks1cHlId1Bvd2liclpIN0F1aEtjUTVDUC9xNFpDWStOUW5CTW1ESUI2Rm9INGxTcmJRaXBFV0xmR3ZGeTFrUGE5NjNXR3V2c2poMzdSNi8vQlBSdEhKdHVVMGhJM3dXeE5oUVhFYVoxQ0dOTkM1VmlidUViMjl2UENkWU9pLzZWYUxIVVVXQ0pldWlNRkhra01Wc0VPcU9nS0pkVk9PS2tPWS84bEFKbFNVVTFEd216U2xkQkNSUUZHOEJiaC9LZGNqa1F2bHJUMXY2dDVJMlExb0JrYlZublNCSVwiLFwiYlwiOlwiUUpYdVpMZWptdURVSjlLVnpSRkFWcHR2eko3RXNFQ09CUUNtK2U0UHVJTUtrMTBLTUVCK1VhaEw4UEw5ZXp4RVZKUE8rOTVXNkhxNGVjdDZyaU9GVjhuSnZSMjloVG5jWnNUb0Z1NDlBMkVSS1B5UThIenBsQm5lWDRLM0Y4bVJwMWJnczVkYmFvenZiTWxySE5tNWx3cUlrVHQ4UlI5VXZHMjNLNXRpZ1lqVExlWWhMS0EySEhFMUtZWWJVYUNFT0Z6N3o4UlMyVVZhbVNZNmFRNmpwVWdSL3dQR25Rb09cIixcImNcIjpcIkZpOWVDTXRsU04xT1JnSjJJNXI2SThBTGJ1ZFFnNVcycmRTa2RoR25sdmZ2VUp5VnhFRmdxNjczN2t0b0Q5ZC84T2IyN0ZCTGVpK1k4TXNRTGQ3VXptSVJKV1F3S0hkYXpkY2FuSFViVTNsSk0rL01uMXhLNkxXOC85VnUyUGtsXCIsXCJkXCI6XCJsL1MrZUdFb0ZNN2w0VjJzNVpzQnpZVUdTWGVGdnBKVjlDQ04zVHRFTTBlQXUwdlQyVlQwZXpFMHg1TGd1OExVUWpKaExUeGdPRDA4a2I3YTd5aHF6NUtxQjB3RlhQRjVDZ2MydGROYW5KemFvUUpQNTZpTUFndjhWcnRWK253TkZiVjJVUXFSelhScjBMRmtGeHhzcXFtMExrK0pWNVhMWUYxa3RyaDFCdC9kdkJzOWxST1ZJdVdRdmdsK2Y5RzlqUnBCekNzQ2EvL2J5WGkvYXIrT005UTNWRTlrU24rb2ZscG5YRi9kMDB1azh1Rkk0ODE5S3NKK2dsQjZBMlpDQlQwb0NHa1VxL0ozVis5TWJuWDZwZHlzZzk4T3VNcHRUZjVkaW1SNGRmZ2h1UTJnZ285RTFTdUhUenI1VHhocFpDazQrWEVCbjYvUm5PV3dZM0xtYkpMSGJIZXIzM0FGdnMxdGNkOXN3K0V2RWo5Q3RwMmQ2NitpWGJJNFh6bFBpZm8rMTd3Y21TcWVyWURKMXVLWVJSSTdwV3ZGVTIvb1JEbXY0Nko4YjMvYmxHK3dPVUFka1FjMUxBTmx2c05ZVnNiZy94NFhhKzJQQVZpVUs5SXlocURXcDFZeW5CSUxiSUZDK2pHWnI4RW04RVRBenhXeFQ5MElnR211ZnVYSktUTTNwTkR6QzlEakxVbzE0eVVZTWNBdjVadnF6M01OWHdrbis3SVptYURxaTJ6NzNLWWJSUT09XCIsXCJnXCI6XCJxU1pnV2EzbXNmT2gwN05NV3FyMnZDdXp6a0NqdWUxdTFEeWRXNnhuZW14dUJzbWNrNlhBRzBha0cyNDZhWkZJczFRSDQyME93VGxianhBclBkYXJXN1d2OFNrdXJlb2NMUUZEaHRIYkhkWTNsZ3pjWUdCNmNqdXpLbWMwS0RQWHVDdndiUVlzUXFHeXZsaEhtN0ZrR0V6OURUV3JFQW9XYmM0U1l2dUU2eTF2U1F6MW5ZQ1h3NmRnRzNEUHpvT2xYS1N2ekZnWFAvdUx6c2k5R3VWUHZNbWhMalNvc0g0QzEvbEhyeXJTNE9zYlRiQWdmNStXTmRhWkRJQnh2Unk1SlZiZHgrM3Y5SGpLT3l0YUVtTUlsdGw5Rm9MYkhkNG83WG51OHRUYUlrVGpSVUcyTlRlSUl3V2ttVGcxaHBFczFvbkxNVHJiaW5TMXBJK29WT2hXWlVZUUx0Z0gxMktiNWVZQjRoVGpKUW9TbG5Wbk5uS0hkSWNwTitjRmEyenE1VTRSZ0YxaDlHTjlkdGFKNDZHLzNPdlN4VWp4ZHZMdjExOS9QSWFhdkNMKzZhK2FieUxUNkE9PVwiLFwiZlwiOlwiait5UW1LcERiRzVLaEs5TzdGTDNpZz09XCIsXCJ1XCI6XCJXZTkzYWE2NTkwYzAtODAwNi00MDk5LTg2MTMtYzhlOGMzYjJkMzI1XCIsXCJlXCI6XCJaZEU1MVRLb3FjU3VYRXNkM1VQbTFRNmFpZ21MWFBnbm9CSzlmZ0tQZjV2clZwWFJuWCtJMjB6eXJSUTlMRDdmUWl4VXlTK3R5L1JPbWVFd05uajBNazE0Q3k3S1BVbnNROGs4N3JvSW1mTT1cIixcInZcIjpcIkhCTk1oKzEzN2JvajJCUC9DWmUvd2c9PVwiLFwiaWRmXCI6XCIxNjQ5MDQ1MTU4OTY5LTE3MzE4NjU1NTUyXCIsXCJ3XCI6XCI0Z2FGTFNCTklGMkNNcjh5OWlCdlZxWnA5TndLSFZCKzhCOUE5WDF0QllzPVwiLFwiY3RcIjpcInkxRG1aU0o0SzdZPVwifSJ9&MmEwMD=5yeSATtABHr0ISR_gVE6pTvhUebXdBP25IHRWVFGX6_SuZYu2tfDmiRh6WeCF7g.toczjkGyI6vIpNLFzChilKSryMcMm31cf0MfzepeAQ3HBZey3U16fBJ2hNSmjWcI5UGgAbr_NEQ95T8HDJSc5OzTF68_8tojrTy49u2.KaIZHWDhgryfxN.E8caG1voBVIs.UlXHa_qwI0rDO.ZjKMiHrU5DhHWZlse7qwJJRXv22n7oCm3_nwvOxKtL5oDgYV2iV8oB01rbG9hakndNFKA`
	txt, err := fetcher.Fetch(profileUrl)
	if err != nil || txt == nil {
		return engine.NilParser(nil)
	}

	err = json.Unmarshal(txt, &profile)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", profile.Data.Name)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
