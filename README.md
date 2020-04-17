# go_generic_api
Go generic API REST

## Config
```
/config
@method GET
@header
{
    "hash":"resetHash"
}
```

## Manage Users

### New User
```
/rest/user/new
@method POST
@header
{
    "id":0,
    "token":"aaaaa"
}
@data
{
    "institution":1, // User institution
    "level": 99,     // Access level user
    "cpf": "None",   
    "name": "Test",  
    "username":"test",
    "password":"123",
    "email":"teste@teste.com"
}
```

### Login
```
/rest/login
@method POST
@data
{
    "institution":1, // User institution
    "username":"test",
    "password":"123",
}
@response
{
    "success":true,
    "institution":1,      // User institution
    "permission": 99,     // Access level user
    "id":"aksja",
    "token":"aaaaaaa"
}
```

### Lambda data
```
/lambda/new
@method POST
@header
{
    "id":0,
    "token":"aaaaa"
}
@data
{
    <Any>
}
```

### Lambda data Tagged
```
/lambda/tag/new
@method POST
@header
{
    "id":0,
    "token":"aaaaa"
}
@data
{
    "tag":"tag",
    <Any>
}
```

### Get Lambda data
```
/lambda/get/{id}
@method GET
@header
{
    "id":0,
    "token":"aaaaa"
}
@response
{
    <Any>
}
```

## Author
* João Carlos Pandolfi Santana