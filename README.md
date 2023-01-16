# social-media-app

A brief description of what this project does and who it's for


# API Reference

## IP Server : 54.254.27.167

#### User Register

```http
  POST /register
```

| Field | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. Full name user |
| `username` | `alphanum` | **Required**. Alphanumeric |
| `email` | `email` | **Required**. Must email format |
| `password` | `string` | **Required**. |

```javascript
{
  "name"     : "Budi Sukses",
  "username" : "amr",
  "email"    : "amr@alterra.id",
  "password" : "rizal123"
}
```
##### Responses
###### 200 OK

```javascript
{
	"message": "success add data"
}
```

###### 400 Bad Request

```javascript
{
	"message": "email/username sudah terdaftar"
}
```

###### 400 Bad Request

```javascript
{
	"message": "field required wajib diisi"
}
```

###### 400 Bad Request

```javascript
{
	"message": "format email salah"
}
```

###### 400 Bad Request

```javascript
{
	"message": "format username salah"
}
```

#### User Login

```http
  POST /login
```
##### with email
| Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `email` | **Required**. Must email format |
| `password`      | `string` | **Required** |

##### with username
| Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required**. Alphanumeric |
| `password`      | `string` | **Required** |


#### User Update

```http
  PUT /users 
```
##### Authorization
| Header | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your API key |

| Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Optional**. |
| `username`      | `string` | **Optional**. Alphanumeric |
| `email`      | `email` | **Optional**.  Must email format|
| `date_of_birth`      | `string` | **Optional**.  |
| `phone_number`      | `numeric` | **Optional**.  Numeric|
| `about_me`      | `string` | **Optional**.  |
| `Password`      | `string` | **Optional**.  |
| `file`      | `file` | **Optional**. file must image |


