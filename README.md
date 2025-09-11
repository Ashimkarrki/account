# Account

## APIs

### 1. CreateAccount

Request

```
{

  "UserName": "ashimkarki",
  "Amt": 1000,
  "Password": "1234"
}

```

Response

```
{
    "Acc_id": "6"
}
```

### 2. Login

Request

```
{
	"UserName":"ashimkarki",
	"Password":"1234"
}
```

Response

```
{
    "Token": "*.*.*",
    "Msg": "Success"
}
```

### 3. UpdateAccount

Request

```
{
	"userName":"ashim",
	"Amt":300,
	"Acc_id":2
}
```

Response

```
{
    "Acc_id": "2",
    "userName": "ashim",
    "Amt": "300"
}
```

### 4. DeleteAccount

Request

```
{
    "Acc_id": 2
}
```

Response

```
{
    "msg": "Deleted"
}
```

### 5. GetAllAccount

Response

```
{
    "Accounts":
    [ {
            "Acc_id": "2",
            "userName": "pasa1",
            "Amt": "1000"
        },
        {
            "Acc_id": "3",
            "userName": "ashi",
            "Amt": "1000"
        }]
}

```

### 6. CreateTransaction

Request

```
{
	"From":2,
	"To":3,
	"Remark":"hi vello"
    "Amt": 40

}
```

Response

```
{
    "Trans_id": "4"
}
```

### 7. ReadTransaction

Request

```
{
    "Trans_id": 4
}
```

Response

```
{
    "From": "2",
    "To": "3",
    "Amt": "40",
    "Remark": "hi vello"
}
```
