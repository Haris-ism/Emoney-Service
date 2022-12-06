# Emoney Service

## Techs: Golang (Gin), Postgres SQL

### Summary

This is a backend project for emoney service.

To access the system you will need JWT Authentication, to get the JWT please use the sign-in and sign-up API.
The sign-up api will hash your password inside database.

For more detail please see the API Endpoints below or you can download the postman testing collection [here](https://www.getpostman.com/collections/05767735d67230325027).

### Note:
- Bankend Port: 8888
- Postgresql port: 5432


### if you are using docker:

Required:
1. Docker

How to run this project:
1. Change your directory to root directory ./Emoney Service
2. Build and Run the images and containers by entering the following command: "docker-compose up"


### if you are not using docker:

Required:
1. Go version 1.19
2. Postgresql version 14

How to run this project:
1. Change your directory to root directory ./Emoney Service
2. Run the following command: "go run main.go"

### API Endpoints

1. Sign Up:
    - Endpoint: /signup
    - Method:POST
    - Body:    {
                    "email":"bambang@gmail.com",
                    "password":"bambang",
                    "name":"bambang"
                }

2. Sign In: 
    - Endpoint: /signin
    - Method:POST
    - Body:{
                "email":"bambang@gmail.com",
                "password":"bambang"
            }

3. Get Profile Info (Sign In Required): 
    - Endpoint: /profile
    - Method:GET

4. Get Balance Info (Sign In Required): 
    - Endpoint: /balance
    - Method:GET

5. Get Inquiry List (Sign In Required):
    - Endpoint:/inquirylist
    - Method:GET

6. Get Transaction History (Sign In Required):
    - Endpoint: /history
    - Method:GET

7. Post Inquiry Confirmation (Sign In Required):
    - Endpoint: /inquiry
    - Method:POST
    - Body:{
                "inquiry_id":4
            }

8. Top-up Account Balance (Sign In Required):
    - Endpoint: /balance
    - Method:PUT
    - Body:{
                "top_up":1000000
            }