# Test Routes

## Ping

`http://localhost:8081/ping`

##  for users without authorization(guests)
## Upcoming Events List
`returns list of events that not deleted & not ended yet`
`GET`
`http://localhost:8081/g/events/upcoming/{page}`

## All Events List
`returns list of all events, including deleted , not ended yet, ended`
`GET`
`http://localhost:8081/g/events/all/{page}`

## Event (by ID)
`returns event by id, dont return deleted event, but returns not ended yet & ended`
`GET`
`http://localhost:8081/g/events/{id}`

##  for users with authorization Admin (ROLE_ADMIN =1)
## Upcoming Events List
`returns list of events that not deleted & not ended yet`
`GET`
`http://localhost:8081/a/events/upcoming/{page}`

## All Events List
`returns list of all events, including deleted , not ended yet, ended`
`GET`
`http://localhost:8081/a/events/all/{page}`

## Event (by ID)
`returns event by id, dont return deleted event, but returns not ended yet & ended`
`GET`
`http://localhost:8081/a/events/{id}`

## Add Event
`POST`
`http://localhost:8081/a/events`
```json
{
  "Title": "ярмарки",
  "ShortDescription": "",
  "Description": "Деснянський район",
  "Longitude": 50.5135884069817,
  "Latitude": 30.614860362771463,
  "Images": "imageSource",
  "Preview": "somePreview",
  "Date": "2022-06-22T00:00:00Z",
  "IsEnded": false

}
```
## Login
`POST`
`http://localhost:8081/a/user/login`
```json
{
  "email":"admin@admin.com",
  "password":"password"
}
```
## Create user
####  using "role_id": 1 for  create user with ROLE_ADMIN
####  using "role_id": 2 for  create user with ROLE_MODERATOR
`POST`
`http://localhost:8081/a/user/login`
```json
{
  "name":"My name8",
  "email":"admin8@admin.com",
  "password":"password",
  "role_id": 1
}
```
##  For users with authorization Moderator (ROLE_MODERATOR =2)
## Upcoming Events List
`returns list of events that not deleted & not ended yet`
`GET`
`http://localhost:8081/m/events/upcoming/{page}`

## All Events List
`returns list of all events, including deleted , not ended yet, ended`
`GET`
`http://localhost:8081/m/events/all/{page}`

## Event (by ID)
`returns event by id, dont return deleted event, but returns not ended yet & ended`
`GET`
`http://localhost:8081/m/events/{id}`

## Add Event
`POST`
`http://localhost:8081/m/events`
```json
{
  "Title": "ярмарки",
  "ShortDescription": "",
  "Description": "Деснянський район",
  "Longitude": 50.5135884069817,
  "Latitude": 30.614860362771463,
  "Images": "imageSource",
  "Preview": "somePreview",
  "Date": "2022-06-22T00:00:00Z",
  "IsEnded": false

}
```
## Login
`POST`
`http://localhost:8081/m/user/login`
```json
{
  "email":"moderator@moderator.com",
  "password":"password"
}
```
```responce in body & in Headers -key Authorization with value of JWT token (type Bearer)
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX3JvbGUiOjEsImV4cCI6MTY1ODg3NDU0Mn0.OY6LHqGdFu6Dg2K4FWxfC2liuRCd3seASLNRNEHtvOg"
}
```
## Using Bearer JWT token  for Moderator or Admin requests