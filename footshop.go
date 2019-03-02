package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/tidwall/gjson"
)

const (
	productid = "Aic-hWgBz42mve0cIOaP"
	sizeid    = "e2e8cbc4-20a7-11e9-9206-02420a000220"
)

type payload struct {
	ID              string `json:"id"`
	SizerunID       string `json:"sizerunId"`
	Account         string `json:"account"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Birthday        string `json:"birthday"`
	DeliveryAddress struct {
		Country     string `json:"country"`
		State       string `json:"state"`
		County      string `json:"county"`
		City        string `json:"city"`
		Street      string `json:"street"`
		HouseNumber string `json:"houseNumber"`
		Additional  string `json:"additional"`
		PostalCode  string `json:"postalCode"`
	} `json:"deliveryAddress"`
	Consents  []string `json:"consents"`
	CardToken string   `json:"cardToken"`
	CardLast4 string   `json:"cardLast4"`
}

type card struct {
	Number        string `json:"number"`
	ExpiryMonth   string `json:"expiryMonth"`
	ExpiryYear    string `json:"expiryYear"`
	Cvv           string `json:"cvv"`
	RequestSource string `json:"requestSource"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func loadjson(filename string) gjson.Result {

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	filebytes, err := ioutil.ReadAll(file)
	check(err)
	return gjson.Parse(string(filebytes))
}

func getcarddata(num, mo, yr, cvv string) (string, string) {
	client := &http.Client{}
	c := card{}

	c.Number = num
	c.ExpiryMonth = mo
	c.ExpiryYear = yr
	c.Cvv = cvv
	c.RequestSource = "JS"

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(c)
	data, err := json.Marshal(c)
	check(err)

	req, err := http.NewRequest("POST", "https://api2.checkout.com/v2/tokens/card", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	useragent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	req.Header.Add("origin", "https://js.checkout.com")
	req.Header.Add("AUTHORIZATION", "pk_76be6fbf-2cbb-4b4a-bd3a-4865039ef187")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", useragent)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	responseString := string(bodyBytes)
	// fmt.Println(responseString)
	cardata := gjson.GetMany(responseString, "id", "card.last4")
	return cardata[0].String(), cardata[1].String()
}

func register(data map[string]gjson.Result) {
	client := &http.Client{}
	p := payload{}

	p.ID = "null"
	p.SizerunID = sizeid
	p.Account = "New Customer"

	p.Email = data["email"].String()
	p.Phone = data["phone"].String()
	p.Gender = data["gender"].String()
	p.FirstName = data["firstname"].String()
	p.LastName = data["lastname"].String()
	p.Birthday = data["birthday"].String()
	p.DeliveryAddress.Country = data["country"].String()
	p.DeliveryAddress.State = ""
	p.DeliveryAddress.County = ""
	p.DeliveryAddress.City = data["city"].String()
	p.DeliveryAddress.Street = data["street"].String()
	p.DeliveryAddress.HouseNumber = data["housenumber"].String()
	p.DeliveryAddress.Additional = ""
	p.DeliveryAddress.PostalCode = data["zip"].String()
	p.Consents = []string{"privacy-policy-101"}

	token, last4 := getcarddata(data["cardnumber"].String(), data["expmonth"].String(), data["expyear"].String(), data["cvv"].String())
	p.CardToken = token
	p.CardLast4 = last4

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(p)
	body, err := json.Marshal(p)
	check(err)

	req, err := http.NewRequest("POST", "https://releases.footshop.com/api/registrations/create/"+productid, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	useragent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("authority", "releases.footshop.com")
	req.Header.Add("referer", "https://releases.footshop.com/register/"+productid+"/Unisex/"+sizeid)
	req.Header.Add("user-agent", useragent)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	responseString := string(bodyBytes)

	fmt.Println("raw data start ----------------")
	fmt.Println(responseString)
	fmt.Println("raw data end ------------------")
	threeDSecURL := gjson.Get(responseString, "secure3DRedirectUrl").String()
	err = exec.Command("open", threeDSecURL).Start()
	check(err)
	fmt.Println("DONE")
}

func main() {
	fmt.Println("Footshop")
	data := loadjson("people.json")
	for _, d := range data.Array() {
		fmt.Println("Registering as", d.Map()["firstname"].String(), d.Map()["lastname"].String())
		register(d.Map())
	}
}
