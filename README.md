# API Specification BookingApp - Golang [Echo Framework]

## Authentication

All API must use this authentication

Request :

* Header :
   * Authorization : "your jwt-token from response api login"


## Registration

* Method : POST
* Endpoint : ```/v1/registration```
* Header : 
     * Content-Type : application/json
     * Accept : application/json
* Body :

```
{
    "username": "string, unique",
    "firstname": "string",
    "lastname": "string",
    "password": "string"
}
```
Response :
```
{
    "code" : "number",
    "status" : "string",
    "data" : {
          "user_id": "number",
          "username": "string",
          "firstname": "string",
          "lastname": "string",
          "date_created": "date"
     }
}
```

## Login

* Method : POST
* Endpoint : ```/v1/login```
* Header : 
     * Content-Type : application/json
     * Accept : application/json
* Body :

```
{
    "username": "string, unique",
    "password": "string"
}
```
Response :
```
{
    "code" : "number",
    "status" : "string",
    "data" : {
         "user_id": "number",
         "username": "string",
         "firstname": "string",
         "lastname": "string",
         "token": "string",
         "date_created": "date"
    }
}
```
