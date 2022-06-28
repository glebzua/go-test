# Test Routes

## Ping

`http://localhost:8081/ping`

## Upcoming Events List
`returns list of events that not deleted & not ended yet`
`http://localhost:8081/v1/events/upcoming/{page}`

## all Events List
`returns list of all events, including deleted , not ended yet, ended`
`http://localhost:8081/v1/events/all/{page}`

## Event (by ID)
`returns event by id, dont return deleted event, but returns not ended yet & ended`
`http://localhost:8081/v1/events/{id}`


## Add Event
`POST`
`http://localhost:8081/v1/events`
```json
{
  "Title": "ярмарки",
  "ShortDescription": "",
  "Description": "Деснянський район",
  "Longitude": 50.5135884069817,
  "Latitude": 30.614860362771463,
  "Images": "imageSource",
  "Preview": "somePpreviev",
  "Date": "2022-06-22T00:00:00Z",
  "IsEnded": true

}
```

