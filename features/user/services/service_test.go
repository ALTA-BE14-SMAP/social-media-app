package services

import (
	"errors"
	"log"
	"mime/multipart"
	"os"
	"social-media-app/features/user"
	"social-media-app/helper"
	"social-media-app/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)
	password := "be1422"
	hash := helper.HashPassword(password)
	t.Run("Berhasil registrasi", func(t *testing.T) {
		inputData := user.Core{
			Name:     "jerry",
			Email:    "jerr@alterra.id",
			Username: "jerrypas77",
			Password: password,
		}

		resData := user.Core{
			ID:       uint(1),
			Name:     "jerry",
			Email:    "jerr@alterra.id",
			Username: "jerrypas77",
			Password: hash,
		}
		inputData.Password = hash
		repo.On("Register", inputData).Return(resData, nil).Once()
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Name, res.Name)
		assert.Equal(t, resData.Email, res.Email)
		assert.Equal(t, resData.Username, res.Username)
		repo.AssertExpectations(t)
	})

	t.Run("email/username sudah terdaftar", func(t *testing.T) {
		inputData := user.Core{
			Name:     "jerry",
			Email:    "jerr@alterra.id",
			Username: "jerrypas77",
			Password: password,
		}
		inputData.Password = hash
		repo.On("Register", inputData).Return(user.Core{}, errors.New("data is duplicated")).Once()
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "sudah terdaftar")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah pada server", func(t *testing.T) {
		inputData := user.Core{
			Name:     "jerry",
			Email:    "jerr@alterra.id",
			Username: "jerrypas77",
			Password: password,
		}
		inputData.Password = hash
		repo.On("Register", inputData).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "masalah pada server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("field required wajib diisi", func(t *testing.T) {
		inputData := user.Core{
			Name:     "jerry",
			Username: "jerrypas77",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wajib diisi")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("format email salah", func(t *testing.T) {
		inputData := user.Core{
			Name:     "jerry",
			Email:    "jerralterra.id",
			Username: "jerrypas77",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("format username salah", func(t *testing.T) {
		inputData := user.Core{
			Name:     "jerry",
			Email:    "jerr@alterra.id",
			Username: "jerrypas77 panda",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t) // mock data
	password := "be1422"
	t.Run("Berhasil login", func(t *testing.T) {
		// input dan respond untuk mock data
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		// res dari data akan mengembalik password yang sudah di hash
		hashed := helper.HashPassword(password)
		resData := user.Core{ID: uint(1), Password: hashed}

		repo.On("Login", inputData).Return(resData, nil).Once() // simulasi method login pada layer data

		srv := New(repo)
		token, res, err := srv.Login(inputData)
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("email/username belum terdaftar", func(t *testing.T) {
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		repo.On("Login", inputData).Return(user.Core{}, errors.New("record not found")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "belum terdaftar")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah pada server", func(t *testing.T) {
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		repo.On("Login", inputData).Return(user.Core{}, errors.New("login query error :")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "masalah pada server")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Salah password", func(t *testing.T) {
		// input dan respond untuk mock data
		inputData := user.Core{
			Email:    "jerr@alterra.id",
			Password: password,
		}
		// res dari data akan mengembalik password yang sudah di hash
		hashed := helper.HashPassword("asdasdasdad")
		resData := user.Core{ID: uint(1), Password: hashed}

		repo.On("Login", inputData).Return(resData, nil).Once() // simulasi method login pada layer data

		srv := New(repo)
		token, res, err := srv.Login(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password tidak sesuai")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("email/username belum terdaftar", func(t *testing.T) {
		inputData := user.Core{
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		token, res, err := srv.Login(inputData)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "belum terdaftar")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("format email salah", func(t *testing.T) {
		inputData := user.Core{
			Email:    "jerralterra.id",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		token, res, err := srv.Login(inputData)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})
	t.Run("format username salah", func(t *testing.T) {
		inputData := user.Core{
			Username: "jerrypas77 panda",
			Password: password,
		}
		srv := New(repo)
		inputData.Password = password
		token, res, err := srv.Login(inputData)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}

		repo.On("Profile", uint(1)).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		repo.On("Profile", uint(4)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah di server", func(t *testing.T) {
		repo.On("Profile", uint(4)).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func getFileHeader(file *os.File) (*multipart.FileHeader, error) {
	// get file size
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// create *multipart.FileHeader
	return &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size:     fileStat.Size(),
	}, nil
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("Berhasil update user tanpa image", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		resData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		repo.On("Update", uint(4), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, pToken, nil)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Name, res.Name)
		assert.Equal(t, resData.Email, res.Email)
		assert.Equal(t, resData.PhoneNumber, res.PhoneNumber)
		repo.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(repo)
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		_, token := helper.GenerateJWT(1)

		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		repo.On("Update", uint(4), inputData).Return(user.Core{}, errors.New("record not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Masalah di server", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		repo.On("Update", uint(4), inputData).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("format email salah", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zakigmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("format username salah", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki 77",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("format phone number salah", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342a",
			AboutMe:     "who am i",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("user tidak ditemukan", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		srv := New(repo)
		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, token, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("format input file tidak dapat dibuka", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}

		file, err := os.Open("../../../mocks/IMG_0225.jpg")
		if err != nil {
			log.Println(err)
		}
		image, _ := getFileHeader(file)

		srv := New(repo)
		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, pToken, image)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "format")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("email/username sudah terdaftar", func(t *testing.T) {
		inputData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		resData := user.Core{
			ID:          4,
			Name:        "Rizal4",
			Email:       "zaki@gmail.com",
			Username:    "amrzaki",
			Photo:       "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			PhoneNumber: "08123022342",
			AboutMe:     "who am i",
		}
		repo.On("Update", uint(4), inputData).Return(resData, errors.New("Duplicate email or password")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(inputData, pToken, nil)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "sudah terdaftar")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}
func TestListUsers(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("Berhasil lihat get users", func(t *testing.T) {
		resData := []user.Core{
			{
				ID:       1,
				Username: "amrzaki1",
				Photo:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872495.jpeg",
			},
			{
				ID:       2,
				Username: "amrzaki2",
				Photo:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872558.jpg",
			},
			{
				ID:       3,
				Username: "amrzaki3",
				Photo:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673872643.jpg",
			},
			{
				ID:       4,
				Username: "amrzaki",
				Photo:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673863241.png",
			},
			{
				ID:       10,
				Username: "amr",
				Photo:    "https://mediasosial.s3.ap-southeast-1.amazonaws.com/images/profile/1673870507.png",
			},
		}
		repo.On("ListUsers").Return(resData, nil).Once()

		srv := New(repo)
		res, err := srv.ListUsers()
		assert.Nil(t, err)
		assert.Equal(t, resData[0].ID, res[0].ID)
		assert.Equal(t, resData[0].Name, res[0].Name)
		assert.Equal(t, resData[0].Email, res[0].Email)
		assert.Equal(t, resData[0].PhoneNumber, res[0].PhoneNumber)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		resData := []user.Core{}
		repo.On("ListUsers").Return(resData, errors.New("record not found")).Once()

		srv := New(repo)
		res, err := srv.ListUsers()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, res, []user.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {
		resData := []user.Core{}
		repo.On("ListUsers").Return(resData, errors.New("query error")).Once()

		srv := New(repo)
		res, err := srv.ListUsers()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res, []user.Core{})
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserData(t)
	t.Run("User tidak ditemukan", func(t *testing.T) {
		_, token := helper.GenerateJWT(0)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Deactive(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("User tidak ditemukan diquery", func(t *testing.T) {

		repo.On("Deactive", uint(1)).Return(errors.New("record not found")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Deactive(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		repo.AssertExpectations(t)
	})

	t.Run("Terjadi kesalahan pada server", func(t *testing.T) {

		repo.On("Deactive", uint(1)).Return(errors.New("query error")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Deactive(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

	t.Run("Berhasil menonaktifkan akun", func(t *testing.T) {
		repo.On("Deactive", uint(1)).Return(nil).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.Deactive(pToken)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
