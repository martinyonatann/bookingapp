# API Specification BookingApp - Golang [Echo Framework]


![alt text] (https://github.com/martinyonathann/bookingapp/blob/master/flowchart.png)


## Authentication

All API must use this authentication

Request :

* Header :
   * Authorization : "your jwt-token from response api login"

#### API Specification BookingApp - Golang [Echo Framework]

## Authentication

All API must use this authentication

Request :

* Header :
   * Authorization : "your jwt-token from response api login"


## USER API
### Registration

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

## Profile

* Method : GET
* Endpoint : ```/v1/profile```
* Header : 
     * Content-Type : application/json
     * Accept : application/json
     * Authorization : "string"

Response :
```
{
    "code" : "number",
    "status" : "string",
    "data" : {
       "user_id": "int",
       "username": "string",
       "firstname": "string",
       "lastname": "string",
       "token": "string",
       "date_created": "date"
   }
}
```
### Registration

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

### Login

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

### Profile

* Method : GET
* Endpoint : ```/v1/profile```

Response :
```
{
    "code" : "number",
    "status" : "string",
    "data" : {
       "user_id": "int",
       "username": "string",
       "firstname": "string",
       "lastname": "string",
       "token": "string",
       "date_created": "date"
   }
}
```

## Hotel API
### Get All Hotel

* Method : GET
* Endpoint : ```/v1/hotel```

Response :
```
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "hotel_id": "number",
      "hotel_name": "string",
      "hotel_address": "string",
      "city": "string",
      "state": "string",
      "zipCode": "number",
      "country": "string",
      "price": "decimal"
    }
  ]
}
```
### Delete Hotel by Id

* Method : DELETE
* Endpoint : ```/v1/hotel/{id}```

Response :
```
{
    "rc": "number",
    "message": "string",
    "detail": "string",
    "data": null
}
```
### Create Hotel

* Method : POST
* Endpoint : ``` /v1/hotel ```

Body Request :
```
{
    "hotel_name": "string",
    "hotel_address": "string",
    "city": "string",
    "state": "string",
    "zipCode": "number",
    "country": "string",
    "price": "decimal"
}

```
Response :
```
{
    "rc": "number",
    "message": "string",
    "detail": "string",
    "data": [
        {
            "hotel_id": "number",
            "hotel_name": "string",
            "hotel_address": "string",
            "city": "string",
            "state": "string",
            "zipCode": "number",
            "country": "string",
            "price": "decimal"
        }
    ]
}

```
### Update Hotel

* Method : POST
* Endpoint : ``` /v1/hotel/update ```

Body Request :
```
{
    "hotel_id": "number",
    "hotel_name": "string",
    "hotel_address": "string",
    "city": "string",
    "state": "string",
    "zipCode": "number",
    "country": "string",
    "price": "decimal"
}

```
Response :
```
{
    "rc": "number",
    "message": "string",
    "detail": "string",
    "data": [
        {
            "hotel_id": "number",
            "hotel_name": "string",
            "hotel_address": "string",
            "city": "string",
            "state": "string",
            "zipCode": "number",
            "country": "string",
            "price": "decimal"
        }
    ]
}

```