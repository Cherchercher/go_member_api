## MVP:

APIs for inserting large csv files into postgres.

Batching is achieved using prepared statement and transaction(1000 rows max per transaction)

## RUNNING ON LOCAL

1. install go
2. modify .env file to use your own crendentials
3. go run main.go

## TESTING

go run test

## TODO
1. auto migrate database: CREATE TABLE member (
id serial, first_name VARCHAR(50),last_name VARCHAR(50),phone_number integer UNIQUE, client_member_id integer UNIQUE, account_id integer)
2. finish writing tests
3. finish implemtation of GET apis
4. run time analysis for large csv files across different insert/batching methods. Currently, inserting 1000 rows takes 21730875 nano seconds (0.02 second)
5. finish api docs

# Batch Create

Insert data into postgres db from csv

**URL** : `/members/new`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
  "name": "valid file name"
}
```

**Data example**

```json
{
  "name": "member_data.csv"
}
```

## Success Response

**Code** : `200 OK`

```json
{
  "message": "success",
  "status": true
}
```

## Error Response

**Condition** : if file does not exist.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "error": "file not found."
}
```
To Be Continued 

# GET MEMBER BY CLIENT_MEMBER_ID

Get a member by client_member_id

**URL** : `/members/client_member_id/{client_member_id}`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

For a User with client member id 1234 on the local database

```json
{
  "id": 12,
  "first_name": "Joe",
  "last_name": "Bloggs",
  "phone_number": 9495221389,
  "client_member_id": 1234,
  "email": "joe25@example.com"
}
```

## ERROR RESPONSE

To Be Continued 

# GET MEMBER BY PHONE_NUMBER

Get a member by phone_number

**URL** : `/members/client_member_id/{phone_number}`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

For a User with phone number 9495112222 on the local database

```json
{
  "id": 12,
  "first_name": "Joe",
  "last_name": "Bloggs",
  "phone_number": 9495112222,
  "client_member_id": 1234,
  "email": "joe25@example.com"
}
```

## ERROR RESPONSE

To Be Continued 

# GET MEMBER BY ID

Get a member by ID

**URL** : `/members/id/{id}`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

For a User with ID 1 on the local database

```json
{
  "id": 1,
  "first_name": "Joe",
  "last_name": "Bloggs",
  "phone_number": 9495112222,
  "client_member_id": 1234,
  "email": "joe25@example.com"
}
```

## ERROR RESPONSE

To Be Continued 

# GET ALL MEMBERS UNDER ACCOUT_ID

Get all members under account id

**URL** : `/members/account_id/{account_id}`

**Method** : `GET`

## Success Response

**Code** : `200 OK`

**Content examples**

For all members with Account ID 1 on the local database

```json
[
  {
    "id": 1,
    "first_name": "Joe",
    "last_name": "Bloggs",
    "phone_number": 9495112222,
    "client_member_id": 1234,
    "email": "joe25@example.com"
  },
  {
    "id": 2,
    "first_name": "Joe",
    "last_name": "Bloggs",
    "phone_number": 9495112222,
    "client_member_id": 1234,
    "email": "joe25@example.com"
  },
  {
    "id": 3,
    "first_name": "Joe",
    "last_name": "Bloggs",
    "phone_number": 9495112222,
    "client_member_id": 1234,
    "email": "joe25@example.com"
  }
]
```
