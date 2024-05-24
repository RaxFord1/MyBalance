package personal

import "encoding/json"

func getCache() *ClientInfoStruct {
	a := []byte(`{}`)

	result := &ClientInfoStruct{}
	if err := json.Unmarshal(a, result); err != nil {
		return nil
	}

	return result
}

//"clientId": "removed",
//"name": "removed",
//"webHookUrl": "",
//"permissions": "removed",
//"accounts": [
//{
//"id": "removed",
//"sendId": "",
//"currencyCode": 980,
//"cashbackType": "UAH",
//"balance": 0,
//"creditLimit": 0,
//"maskedPan": [
//"removed"
//],
//"type": "removed",
//"iban": "removed"
//},
//{
//"id": "removed",
//"sendId": "3Uqk8jcpXb",
//"currencyCode": 840,
//"cashbackType": "UAH",
//"balance": 0,
//"creditLimit": 0,
//"maskedPan": [
//"removed"
//],
//"type": "black",
//"iban": "removed"
//},
//{
//"id": "removed",
//"sendId": "removed",
//"currencyCode": 980,
//"cashbackType": "UAH",
//"balance": 0,
//"creditLimit": 0,
//"maskedPan": [
//"removed",
//"removed"
//],
//"type": "black",
//"iban": "removed"
//},
//{
//"id": "removed",
//"sendId": "removed",
//"currencyCode": 980,
//"cashbackType": "UAH",
//"balance": 0,
//"creditLimit": 0,
//"maskedPan": [
//"removed"
//],
//"type": "white",
//"iban": "removed"
//},
//{
//"id": "removed",
//"sendId": "",
//"currencyCode": 980,
//"balance": 0,
//"creditLimit": 0,
//"maskedPan": [],
//"type": "fop",
//"iban": "removed"
//}
//]
