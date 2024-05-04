### Payhere API
#### Installation
```go
go get -u github.com/vannleonheart/payhere-api-go
```
#### Config
```go
payhereConfig := payhere.Config {
    BaseUrl: "{payhere_api_url}",
    Token: "{your_payhere_token}"
}
```
Payhere Base Url for production
```text
https://api.payhere.id
```
#### Create Client
```go
payhereClient := payhere.New(&payhereConfig)
```
#### Get Balance
```go
result, err := payhereClient.GetBalance()

if err != nil {
    // handle error
}

fmt.Println(result.Balance)
```
#### Create Disbursement
```go
testAmount := 10000

result, err := payhereClient.CreateDisbursement(payhere.CreateDisbursementRequest{
    MerchantTrxId:          "{your_transaction_id}",
    RecipientName:          "{recipient_name}",
    RecipientBankCode:      "{indonesian_bank_code}",
    RecipientAccountNumber: "{recipient_bank_account_number}",
    Amount:                 testAmount,
    Description:            "{optional_description}",
    CallbackUrl:            "{optional_your_callback_url}",
})

if err != nil {
    // handle error
}

fmt.Println(result.Uuid)
```
#### Get Disbursement Detail By UUID
```go
uuid := "{transaction_uuid}"

result, err := payhereClient.InquiryDisbursementByUuid(uuid)

if err != nil {
    // handle error
}

fmt.Println(result.MerchantTrxId)
```
#### Get Disbursement Detail By Merchant Transaction Id
```go
merchantTrxId := "{your_transaction_id}"

result, err := payhereClient.InquiryDisbursementByMerchantTrx(merchantTrxId)

if err != nil {
    // handle error
}

fmt.Println(result.Uuid)
```