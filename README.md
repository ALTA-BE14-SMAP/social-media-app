
# social-media-app

A brief description of what this project does and who it's for


# API Reference

## IP Server : 54.254.27.167


#### Status Codes

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 401 | `UNAUTHORIZED` |
| 404 | `NOT FOUND` |
| 409 | `CONFLICT` |
| 500 | `INTERNAL SERVER ERROR` |

### User Register

```http
  POST /register
```

##### Form/JSON
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
###### 201 Created

```javascript
{
	"message": "success add data"
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

###### 409 Conflict

```javascript
{
	"message": "email/username sudah terdaftar"
}
```

###### 500 Internal Server Error

```javascript
{
	"message": "data tidak bisa diolah"
}
```
### User Login

```http
  POST /login
```
##### with email
##### Form/JSON
| Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `email` | **Required**. Must email format |
| `password`      | `string` | **Required** |

```javascript
{
  "email"    : "amr@alterra.id",
  "password" : "rizal123"
}
```

##### with username
##### Form/JSON
| Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required**. Alphanumeric |
| `password`      | `string` | **Required** |

```javascript
{
  "username" : "amr",
  "password" : "rizal123"
}
```

##### Responses
###### 200 OK

```javascript
{
	"message": "success login",
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzM4NTk2NTIsInVzZXJJRCI6MX0.A4KwWP8vn_VMOufnLPX6hF_VLCHtI1WvY-Za8sDQhuE"
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

###### 401 Unauthorized

```javascript
{
	"message": "password tidak sesuai"
}
```
###### 404 Not Found

```javascript
{
	"message": "email/username belum terdaftar"
}
```
###### 500 Internal Server Error

```javascript
{
	"message": "data tidak bisa diolah"
}
```




### User Update

```http
  PUT /users 
```
##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Multipart Form
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

##### Responses
###### 200 OK

```javascript
{
	"message": "berhasil update profil"
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

###### 400 Bad Request

```javascript
{
	"message": "format phone_number salah"
}
```
###### 400 Bad Request

```javascript
{
	"message": "format input file type tidak diizinkan"
}
```
###### 400 Bad Request

```javascript
{
	"message": "format input file size tidak diizinkan, size melebihi 1 MB"
}
```

###### 401 Unauthorized

```javascript
{
	"message": "invalid or expired jwt"
}
```
###### 409 Conflict

```javascript
{
	"message": "email/username sudah terdaftar"
}
```
###### 500 Internal Server Error

```javascript
{
	"message": "data tidak bisa diolah"
}
```

## User Profile

```http
  GET /myprofile 
```

##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |


##### Responses
###### 200 OK

```javascript
{
	"data": {
		"id": 4,
		"name": "Rizal4",
		"email": "zaki@gmail.com",
		"username": "amrzaki",
		"photo": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
		"date_of_birth": "",
		"phone_number": "08123022342",
		"about_me": "who am i"
	},
	"message": "berhasil lihat profil"
}
```

###### 401 Unauthorized

```javascript
{
	"message": "invalid or expired jwt"
}
```

###### 500 Internal Server Error

```javascript
{
	"message": "data tidak bisa diolah"
}
```
## User Get All

```http
  GET /users 
```

##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Responses
###### 200 OK

```javascript
{
	"data": [
		{
			"id": 1,
			"username": "amrzaki1",
			"photo": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg"
		},
		{
			"id": 2,
			"username": "amrzaki2",
			"photo": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872558.jpg"
		},
		{
			"id": 3,
			"username": "amrzaki3",
			"photo": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872643.jpg"
		},
		{
			"id": 4,
			"username": "amrzaki",
			"photo": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png"
		},
		{
			"id": 10,
			"username": "amr",
			"photo": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673870507.png"
		}
	],
	"message": "berhasil lihat profil"
}
```

###### 401 Unauthorized

```javascript
{
	"message": "invalid or expired jwt"
}
```

###### 500 Internal Server Error

```javascript
{
	"message": "data tidak bisa diolah"
}
```