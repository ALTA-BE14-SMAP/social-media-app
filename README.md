
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
| `image`      | `file` | **Optional**. file must image |

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
## User Deactive

```http
  DELETE /users 
```

##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Responses
###### 200 OK
```javascript
{
	"message": "berhasil menonaktifkan akun"
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

### Content Add

```http
  POST /contents
```

##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Form/JSON
| Field | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `content` | `string` | **Required** |
| `image` | `image` | **Optional**. must image |


##### Responses
###### 201 Created
```javascript
{
	"message": "posting content berhasil"
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
	"message": "kesalahan input"
}
```
###### 500 Internal Server Error
```javascript
{
	"message": "data tidak bisa diolah"
}
```

### Content GetAll
```http
 GET /contents
```

##### Responses
###### 200 OK
```javascript
{
    "data": [
        {
            "id": 2,
            "content": "halo gaes",
            "image": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674014577.jpg",
            "userid": 1,
            "name": "devangga"
        },
        {
            "id": 3,
            "content": "spongebob",
            "image": "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1674014592.jpg",
            "userid": 1,
            "name": "devangga"
        }
    ],
    "message": "berhasil menampilkan content"
}
```
###### 404 Not Found
```javascript
{
	"message": " content tidak ditemukan"
}
```
###### 500 Internal Server Error
```javascript
{
	"message": "data tidak bisa diolah"
}
```

### Content Update
```http
  PUT /contents
``` 
##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Multipart Form
| Field | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `content` | `string` | **Required** |
| `image` | `image` | **Optional**. must image |

##### Responses
###### 200 OK
```javascript
{
	"message": "berhasil update content"
}
```
###### 400 Bad Request
```javascript
{
	"message": " Error binding data "
}
```
###### 400 Bad Request
```javascript
{
	"message": "kesalahan input"
}
```
###### 400 Bad Request
```javascript
{
	"message": " format input file tidak dapat dibuka "
}
```
###### 400 Bad Request
```javascript
{
	"message": " format input file type tidak diizinkan "
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

### Content Delete
```http
  DELETE /contents
```

| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Responses
###### 200 OK
```javascript
{
    "message": "berhasil delete content"
}
```
###### 400 Bad Request
```javascript
{
	"message": "Kesalahan input"
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

## Comment Add

```http
  POST /comments/{idPost}
```

##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Parameters Path
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `{idPost}`      | `numuric` | **Required**. Your id post/content |

##### Form/JSON
| Field | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `text` | `string` | **Required**. text for your comment|

##### Responses
###### 201 Created

```javascript
{
	"message": "success add comment"
}
```


###### 400 Bad Request

```javascript
{
	"message": "field required wajib diisi"
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


## List Comments  (by id post/content)

```http
  GET /comments/{idPost}
```

##### Parameters Path
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `{idPost}`      | `numuric` | **Required**. Your id post/content |

##### Responses
###### 200 OK

```javascript
{
	"data": [
		{
			"id": 1,
			"text": "Happy Second Birthday nak kicik❤️Terimakasih sudah menjadikanku ibu, selalu setia jadi teman ibu👶",
			"created_at": "2023-01-18 18:26:30.472 +0700 WIB",
			"commentator": "Budi Sukses"
		},
		{
			"id": 2,
			"text": "You can add any url params you want there and it will show you the url preview above it.",
			"created_at": "2023-01-18 18:30:55.665 +0700 WIB",
			"commentator": "Budi Sukses"
		},
		{
			"id": 3,
			"text": "Sy jg pengen punya app kayak gini",
			"created_at": "2023-01-18 18:33:24.388 +0700 WIB",
			"commentator": "Budi Sukses"
		},
		{
			"id": 4,
			"text": "Tutorial yg paling banyak dicari tentang golang menurut saya adalah cara Deploy ke production di hosting yg murah untuk produktion,",
			"created_at": "2023-01-18 18:51:15.698 +0700 WIB",
			"commentator": "Budi Sukses"
		}
	],
	"message": "Berhasil melihat list comments"
}
```

###### 500 Internal Server Error

```javascript
{
	"message": "data tidak bisa diolah"
}
```
## Comment Delete

```http
  DELETE /comments/{idPost}/{idComment}
```

##### Authorization JWT
| Authentication | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `bearer token`      | `string` | **Required**. Your token key |

##### Parameters Path
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `{idPost}`      | `numuric` | **Required**. Your id post/content |
| `{idComment}`      | `numuric` | **Required**. Your id comment |

##### Responses
###### 200 OK
```javascript
{
	"message": "Berhasil delete comment"
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