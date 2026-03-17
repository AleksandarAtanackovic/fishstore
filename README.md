# Fishstore

## Introduction
This is a learning project.  
The idea is to build a web app for a fishstore that will serve as an order management and tracking solution.
The app will be built using Go as a backend and React for frontend. 
Feel free to give me advice if you stumble upon this project somehow.



## How to run the app
Currently in this stage of development, only local testing is possible with or without Docker.  
### !!IMPORTANT!!
Since i am working on connecting the app to the database, if you don't have postgres installed locally, the app won't start. You will also need to set the password for the postgres user to be `password`, or change it in the source code.

### Without Docker
Download source code or clone the repo.  
Position in the folder `/fishstore`.  
If you have GO installed, run `go run backend/cmd/server/main.go`  

### Docker
For now, until a public version of the docker image is published, clone the source code repo or download it, position at the root folder /fishstore and run  
`docker build -t fishstore:0.0.1 .`  
This will build the app image with a tag fishstore:0.0.1.  
Usually you would run `docker run -p 9090:9090 fishstore:0.0.1` to start the app, however if you have postgres installed locally like i do right now, you should run `docker run --network=host fishstore:0.0.1` so the app can reach the database from inside a docker container.


Whether you started the app with or without docker now you can test the api with some http client like `curl`, `ThunderClient` or `Postman`.  

Example of a `POST` call with curl: 
```
curl -X POST "http://localhost:9090/api/v1/orders" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "1",
    "customer": {
      "id": "c1",
      "name": "Zika",
      "surname": "Zikic",
      "phone_number": "0601012456"
    },
    "fish_type": "saran",
    "order_type": "fry",
    "prepared": false,
    "completed": false
  }'
```
Example of a response:
```
{
  "id": "1",
  "customer": {
    "id": "c1",
    "name": "Zika",
    "surname": "Zikic",
    "phone_number": "0601012456"
  },
  "created_at": "2025-11-26T11:43:02.032788969+01:00",
  "fish_type": "saran",
  "order_type": "fry",
  "prepared": false,
  "completed": false
}
```

## Api specification

### Available endpoints

The 1st version of api is located on the path `/api/v1`.  
You can call:  
`GET /orders` or `GET /order/:id` to get all the orders or the order with a specific id.  
`GET /orders/unfinished` to get all uncompleted orders.  
`POST /orders` to create an order.  
`PATCH /orders/:id` to toggle the `prepared` value true/false.  

### Data validation

Some data validation will be done during an API call.  
* `fish_type` must be one of the following values: `saran`, `pastrmka` or `oslic`.  
* `order_type` must be one of the following values: `fry`, `clean`, or `fresh`.  

