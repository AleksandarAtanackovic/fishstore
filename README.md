# Fishstore

## Introduction
This is a learning project.  
The idea is to build a web app for a fishstore that will serve as an order management and tracking solution.
The app will be built using Go as a backend and React for frontend. 
Feel free to give me advice if you stumble upon this project somehow.



## How to run the app
Currently in this stage of development, only local testing is possible without docker.  
Download source code or clone the repo.  
Position in the folder `/fishstore`.  
If you have GO installed, run `go run backend/cmd/server/main.go`  
Then test the api with some http client like `curl`, `ThunderClient` or `Postman`.  

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

