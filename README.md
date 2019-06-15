# Footshop Releases Bot

Automated raffle script made to easily enter raffles on releases.footshop.com
If you don't have Go language installed follow [this](https://golang.org/doc/install) link.


// *THIS IS A POOR HEADLESS VERSION, IM NO LONGER SUPPORTING A DEVELOPMENT OF ANY BOT RELATED TO FOOTSHOP. HERE'S A QUICK DEMONSTRATION OF REQUESTS SO YOU CAN CODE IT YOURSELF EVEN THO ITS USELESS AS YOU AREN'T GOING TO COOK ANYTHING*


## Send a POST request to get a card_id_token needed to send a valid checkout post request.
*Send POST to this URL: *https://api2.checkout.com/v2/tokens/card* with following header details.*
```
{
	'Accept': 'application/json, text/javascript, */*; q=0.01',
	'Referer': 'https://js.checkout.com/frames/?v=1.0.16&publicKey=pk_76be6fbf-2cbb-4b4a-bd3a-4865039ef187&localisation=EN-GB&theme=standard',
	'Origin': 'https://js.checkout.com',
	'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36',
	'AUTHORIZATION': 'pk_76be6fbf-2cbb-4b4a-bd3a-4865039ef187',
	'Content-Type': 'application/json'
}
```
*BODY part should be formatted in application/json format*
```
{
	"number": "1488148814881488",
	"expiryMonth": "14",
	"expiryYear": "2088",
	"cvv": "148",
	"requestSource": "JS"
}
```
*Now we have a valid card_id token so we can send a POST request to get a successful raffle signup.*

## Send a POST request to Footshop Releases's API to get a successful signup

*Send POST to this URL: *https://releases.footshop.com/api/registrations/create/VARIANT* with following header details.*
^Example: https://releases.footshop.com/api/registrations/create/-AXT6WoB8Bg_k2wdvyvE (Yeezy 700s Black)

Variant is the ID of product you want to sign-up.

Header details:
```
{
    'origin': 'https://releases.footshop.com',
    'accept-language': 'en-GB,en-US;q=0.9,en;q=0.8',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36',
    'content-type': 'application/json;charset=UTF-8',
    'accept': 'application/json, text/plain, */*',
    'cache-control': 'no-cache',
    'authority': 'releases.footshop.com',
    'referer': 'https://releases.footshop.com/register/' + task['variant'] + '/Unisex/' + sizeID
}
```
Body part, make sure your data is in correct format
sizerunID is a size you're going for, cardToken is the token we got in previous step.
```
{
	"id": null,
	"sizerunId": "d7ade886-7e1f-11e9-995c-0242ac13000a",
	"account": "New Customer",
	"email": "your@mail.cz",
	"phone": "+420999999999",
	"gender": "Mr",
	"firstName": "Samuel",
	"lastName": "Sasasasa",
	"birthday": "1997-01-01",
	"deliveryAddress": {
		"country": "CZ",
		"state": "CZ",
		"county": "",
		"city": "Prague",
		"street": "Sartoriova",
		"houseNumber": "1488",
		"additional": "",
		"postalCode": "16900"
	},
	"consents": ["privacy-policy-101"],
	"cardToken": "CARD_TOKEN_WE_GOT_IN_PREVIOUS_STEP",
	"cardLast4": "XXXX"
	
}
```

So that's it, now you got a valid raffle signup. Back to the original Bot and it's files.

## Installation
Tested on macOS & Ubuntu 18.04 Desktop

1. Clone this repo to your machine.
```
git clone https://github.com/samoinsecure/footshopReleases-raffle-bot
```
2. Navigate to your cloned repository
```
cd $HOME/footshopReleases-raffle-bot
```
3. Edit your people.json file with your Billing details and save the file (⌘+S)

3.1. Edit your footshop.go file with your selected SizeID and ProductID (which could be found in the product url) and save the file (⌘+S)

4. Make the file executable by running this command
```
go build footshop.go
```
5. Run the script by typing the following command to your terminal.
```
./footshop
```
or
```
go run footshop.go
```
Enjoy!


## Contact me
*Feel free to DM me on Twitter or Discord if you need any help*.

* [Twitter](https://twitter.com/samoinsecure)
* [Mail](mailto:github@samuelmikula.com)

*DISCLAIMER*
*Your sensitive informations are stored locally on your machine.*

