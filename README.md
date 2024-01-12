
# API Documentation

## 1 Authentikasi Management

### 1.1 Insert User Data

Request :
- Method : POST
- URL : `{{local}}:3636/api/tesbedbo/v1/users/insert`
- Body (form-data) :
    - username : string, required
    - password : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 11,
        "username": "bela",
        "password": "$2a$04$r8ZnYATkORLx12y0/hUHz.L6ZX0pkm1yWmXSyhViLHZrqwQS4.btq",
        "created_at": "2024-01-12T13:42:19.003+07:00",
        "updated_at": "2024-01-12T13:42:19.003+07:00"
    }
}
```

### 1.2 Login

Request :
- Method : POST
- URL : `{{local}}:3636/api/tesbedbo/v1/auth/login`
- Body (form-data) :
    - username : string, required
    - password : string, required
- Response :

```json 
{
    "meta": {
        "message": "Login success.",
        "code": 200
    },
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc2MjM4MjksInVzZXIiOnsiaWQiOjl9fQ.uaXT_ityjLBUjsiKH7ay8xdKNnWbx8_hOYoEGqo2PPc",
        "expires": "2024-02-11T10:57:09.345711417+07:00"
    }
}
```

### 1.3 Get Login Data

Request :
- Method : GET
- URL : `{{local}}:3636/api/tesbedbo/v1/logindata/getlogindata`
- Header : 
    - Authorization : string
- Response :

```json 
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 9,
        "username": "adit",
        "created_at": "2024-01-12T10:13:57+07:00",
        "updated_at": "2024-01-12T10:57:09+07:00"
    }
}
```

### 1.2 Logout

Request :
- Method : POST
- URL : `{{local}}:3636/api/tesbedbo/v1/auth/logout`
- Header : 
    - Authorization : string
- Response :

```json 
{
    "meta": {
        "message": "Logout success.",
        "code": 200
    },
    "data": null
}
```

## 2 Customer Management

### 2.1 Get Detail

Request :
- Method : GET
- URL : `{{local}}:3636/api/tesbedbo/v1/customers/getdetail/9`
- Header :
    - Authorization : string
- Response :
    
```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 9,
        "user_id": 8,
        "name": "Sinta",
        "email": "sinta@gmail.com",
        "phone": "3232",
        "street": "jl.kebaikan",
        "zip_code": 3351,
        "city": "Malang",
        "country": "Indonesia",
        "created_at": "2024-01-12T08:13:22+07:00",
        "updated_at": "2024-01-12T08:13:22+07:00"
    }
}
```

### 2.2 Get All With Paginate & Search

Request :
- Method : GET
- URL : `{{local}}:3636/api/tesbedbo/v1/customers/getwithpaginateandsearch
- Header :
    - Authorization : string
- Params :
    - limit 
    - page
    - sort
    - order
    - user_id
    - name
    - email
    - phone
    - street
    - zip_code
    - city
    - country
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 1,
        "total_filtered": 1
    },
    "data": [
        {
            "id": 9,
            "user_id": 8,
            "name": "Sinta",
            "email": "sinta@gmail.com",
            "phone": "3232",
            "street": "jl.kebaikan",
            "zip_code": 3351,
            "city": "Malang",
            "country": "Indonesia",
            "created_at": "2024-01-12T08:13:22+07:00",
            "updated_at": "2024-01-12T08:13:22+07:00"
        }
    ]
}
```

### 2.3 Insert

Request :
- Method : POST
- URL : `{{local}}:3636/api/tesbedbo/v1/customers/insert`
- Header :
    - Authorization : string
- Body (form-data) :
    - user_id : int, required
    - name : string, required
    - email : string, required
    - phone : string, required
    - street : string, required
    - zip_code : int, required
    - city : string, required
    - country : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 13,
        "user_id": 10,
        "name": "Roy Suryo",
        "email": "roy@gmail.com",
        "phone": "2222",
        "street": "jl.kebaikan",
        "zip_code": 3351,
        "city": "Malang",
        "country": "Indonesia",
        "created_at": "2024-01-12T13:25:58.473+07:00",
        "updated_at": "2024-01-12T13:25:58.473+07:00"
    }
}
```

### 2.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/api/tesbedbo/v1/customers/update/13`
- Header :
    - Authorization : string
- Body (form-data) :
    - user_id : int
    - name : string
    - email : string
    - phone : string
    - street : string
    - zip_code : int
    - city : string
    - country : string
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 13,
        "user_id": null,
        "name": "",
        "email": "",
        "phone": "08113600888",
        "street": "",
        "zip_code": 0,
        "city": "",
        "country": "",
        "created_at": null,
        "updated_at": "2024-01-12T13:26:56.815+07:00"
    }
}
```

### 2.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/api/tesbedbo/v1/customers/delete/13`
- Header :
    - Authorization : string
- Response : 

```json
{
    "meta": {
        "message": "Successfully deleted data.",
        "code": 200
    },
    "data": null
}
```

## 3 Order Management

### 3.1 Get Detail

Request :
- Method : GET
- URL : `{{local}}:3636/api/tesbedbo/v1/orders/getdetail/2`
- Header :
    - Authorization : string
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 2,
        "customer_id": 9,
        "total_amount": 1500,
        "status": "success",
        "payment_type": "debit card",
        "created_at": "2024-01-12T09:48:43+07:00",
        "updated_at": "2024-01-12T09:48:43+07:00"
    }
}
```

### 3.2 Get All With Paginate & Search

Request :
- Method : GET
- URL : `{{local}}:3636/api/tesbedbo/v1/orders/getwithpaginateandsearch
- Header :
    - Authorization : string
- Params :
    - limit 
    - page
    - sort
    - order
    - customer_id
    - total_amount
    - status
    - payment_type
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 1,
        "total_filtered": 1
    },
    "data": [
        {
            "id": 2,
            "customer_id": 9,
            "total_amount": 1500,
            "status": "success",
            "payment_type": "debit card",
            "created_at": "2024-01-12T09:48:43+07:00",
            "updated_at": "2024-01-12T09:48:43+07:00"
        }
    ]
}
```

### 3.3 Insert

Request :
- Method : POST
- URL : `{{local}}:3636/api/tesbedbo/v1/orders/insert`
- Header :
    - Authorization : string
- Body (form-data) :
    - customer_id : int, required
    - total_amount : int, required
    - status : string, required
    - payment_type : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 3,
        "customer_id": 9,
        "total_amount": 1500,
        "status": "success",
        "payment_type": "debit card",
        "created_at": "2024-01-12T13:29:04.973+07:00",
        "updated_at": "2024-01-12T13:29:04.973+07:00"
    }
}
```

### 3.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/api/tesbedbo/v1/orders/update/3`
- Header :
    - Authorization : string
- Body (form-data) :
    - customer_id : int
    - total_amount : int
    - status : string
    - payment_type : string
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 3,
        "customer_id": 9,
        "total_amount": 3500,
        "status": "success",
        "payment_type": "debit card",
        "created_at": "2024-01-12T13:29:04.973+07:00",
        "updated_at": "2024-01-12T13:29:04.973+07:00"
    }
}
```

### 3.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/api/tesbedbo/v1/orders/delete/3`
- Header :
    - Authorization : string
- Response :

```json
{
    "meta": {
        "message": "Successfully deleted data.",
        "code": 200
    },
    "data": null
}
```


