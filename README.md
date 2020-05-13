# Goldie Blocking Chain

A RESTful representation demonstrating the underlying principles of how Blockchain conceptually works, using an oversimplified implementation of it in Go. 

## Usage

`go run main.go`

## Endpoints

### Sending 

`POST` `/api/v1/blockchain/send`

**Request**

JSON Payload:

```
{
	"from": "adam",
	"to": "eve",
	"amount": 75
}
```

**Response**

```
{
    "message": "adam successfully sent 75 golden nuggets to eve"
}
```

**Errors**

If there is an insufficient amount of gold in the account when sending:

```
{
    "status_code": 500,
    "error_message": "adam does not have enough funds. Current balance: 25"
}
```

### Listing

`GET` `/api/v1/blockchain/list`

**Response**

```
{
    "blocks": [
        {
            "hash": "AA50MfQmw4JusuR0bijxKkvYsFOnMilFagTc6PaEyGI=",
            "transactions": [
                {
                    "id": "SosXal+dlbOR9aghh/CHd4xC/AASddHDVWSjC6VZEtM=",
                    "inputs": [
                        {
                            "input_id": "",
                            "output": -1,
                            "signature": "This is a reference to the genesis block"
                        }
                    ],
                    "outputs": [
                        {
                            "value": 100,
                            "address": "adam"
                        }
                    ]
                }
            ],
            "previous_hash": "",
            "nonce": 12477
        },
        {
            "hash": "AALeW5pkAUpo6/rbgUnt16M1H0tpUsbUTKPqikHicBg=",
            "transactions": [
                {
                    "id": "i3OJIXB8ONB9VYOtbmZKc1nub+H/vIs3O+JJGvBU0bA=",
                    "inputs": [
                        {
                            "input_id": "SosXal+dlbOR9aghh/CHd4xC/AASddHDVWSjC6VZEtM=",
                            "output": 0,
                            "signature": "adam"
                        }
                    ],
                    "outputs": [
                        {
                            "value": 75,
                            "address": "eve"
                        },
                        {
                            "value": 25,
                            "address": "adam"
                        }
                    ]
                }
            ],
            "previous_hash": "AA50MfQmw4JusuR0bijxKkvYsFOnMilFagTc6PaEyGI=",
            "nonce": 1242
        }
    ]
}
```

### Show Balance

`GET` `/api/v1/blockchain/balance/:addr`

**Request**

`addr` is an identifier which is usually a public key, but names are used here as a simple means of demonstration. The first block is signed by `adam`.  So to show the balance of `adam`:

`GET` `/api/v1/blockchain/balance/adam`

**Response**

```
{
    "message": "Balance of adam: 25 golden nuggets"
}
```


