# Taxi cars

Web services for car management.

- Framework: [Gin Framework](https://github.com/gin-gonic/gin)
- ORM: [GORM](https://gorm.io/index.html)


## Get all cars
```sh
curl -X GET 'localhost:8080/v1/cars'
```

## Add new car
```sh
curl -X POST 'localhost:8080/v1/cars' -H 'Content-Type: application/json' --data '{"regPlate": "A001","brand": "Toyota","color": "Black", "year": 2022, "isReady": false, "posX": 2, "posY": 3}'
```

## Get car by id
```sh
curl -X GET 'localhost:8080/v1/cars/1'
```

## Update car by id
```sh
curl -X PUT 'localhost:8080/v1/cars/1' -H 'Content-Type: application/json' --data '{"regPlate": "B002","brand": "Honda","color": "Grey", "year": 2021, "isReady": true, "posX":3, "posY":4}'
```

## Find nearest cars by client's coordinates (distance: 5)
```sh
curl -X GET 'localhost:8080/v1/cars/nearest?x=3&y=3'
```

## Reserve car by id
```sh
curl -X PUT 'localhost:8080/v1/cars/1/reserve'
```

## Delete car by id
```sh
curl -X DELETE 'localhost:8080/v1/cars/1'
```