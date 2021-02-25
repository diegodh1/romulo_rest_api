package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

//DecryptPass
func decryptPass(encryptedString string) (decryptedString string) {

	key := []byte("integrappsssssss")
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

//EncryptPass func
func encryptPass(stringToEncrypt string) (encryptedString string) {
	//Since the key is in string, we need to convert decode it to bytes
	key := []byte("integrappsssssss")
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

//CreateToken create JSON WEB TOKEN
func CreateToken(userid string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
func getProfiles(userID string, db *gorm.DB) []AppUserProfile {
	profiles := []AppUserProfile{}
	if err := db.Where("user_id = ? and status = ?", userID, true).Find(&profiles).Error; err != nil {
		return profiles
	}
	return profiles
}

//LoginUser Get a user from DB
func LoginUser(userID string, password string, db *gorm.DB) Response {
	userApp := AppUser{UserID: userID}
	if err := db.Where("user_id = ?", userID).First(&userApp).Error; err != nil {
		return Response{Payload: nil, Message: "El usuario no está registrado en la base de datos", Status: 403}
	}
	switch {
	case decryptPass(userApp.Password) != password:
		return Response{Payload: nil, Message: "Contraseña incorrecta", Status: 401}
	case *userApp.Status == false:
		return Response{Payload: nil, Message: "El usuario no está activo en el sistema", Status: 403}
	default:
		token, err := CreateToken(userApp.UserID)
		if err != nil {
			return Response{Payload: nil, Message: "Error interno del servidor", Status: 500}
		}
		profiles := getProfiles(userID, db)
		var payload struct {
			User     AppUser
			Token    string
			Profiles []AppUserProfile
		}
		//profiles of the user and token
		payload.Token = token
		payload.Profiles = profiles
		payload.User = userApp
		//return
		return Response{Payload: payload, Message: "OK", Status: 200}
	}

}

//UpdateUser function
func UpdateUser(user *AppUser, profiles *[]AppUserProfile, db *gorm.DB) Response {
	switch {
	case user.UserID == "":
		return Response{Payload: nil, Message: "El ID de usuario es obligatorio", Status: 400}
	case user.Name == "":
		return Response{Payload: nil, Message: "El nombre es obligatorio", Status: 400}
	case user.Email == "" || !strings.Contains(user.Email, "@"):
		return Response{Payload: nil, Message: "El correo es obligatorio", Status: 400}
	default:
		if err := db.Model(&user).Omit("UserID", "Password", "CreactionDate").Where("user_id = ?", user.UserID).Updates(user).Error; err != nil {
			return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
		}
		for _, v := range *profiles {
			assignProfile(&v, db)
		}
		return Response{Payload: nil, Message: "Actualización Realizada!", Status: 200}
	}
}

//SearchUser struct
func SearchUser(userID string, db *gorm.DB) Response {
	var appUser AppUser
	var profiles []AppUserProfile
	db.Where("user_id = ?", userID).First(&appUser)
	db.Where("user_id = ?", userID, true).Find(&profiles)
	return Response{Payload: User{User: appUser, Profiles: profiles}, Message: "OK", Status: 200}
}

//CreateUser create a new user in the db
func CreateUser(user *AppUser, profiles *[]AppUserProfile, db *gorm.DB) Response {

	switch {
	case user.UserID == "":
		return Response{Payload: nil, Message: "El ID de usuario es obligatorio", Status: 400}
	case user.Password == "":
		return Response{Payload: nil, Message: "La contraseña no puede estar vacía", Status: 400}
	case user.Name == "":
		return Response{Payload: nil, Message: "El nombre es obligatorio", Status: 400}
	case user.Email == "" || !strings.Contains(user.Email, "@"):
		return Response{Payload: nil, Message: "El correo es obligatorio", Status: 4000}
	default:
		user.Name = strings.ToUpper(user.Name)
		user.LastName = strings.ToUpper(user.LastName)
		user.CreactionDate = time.Now()
		user.Password = encryptPass(user.Password)
		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "PRIMARY KEY") {
				return Response{Payload: nil, Message: "El registro ya existe en el sistema", Status: 400}
			}
			return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
		}
		for _, v := range *profiles {
			assignProfile(&v, db)
		}
		return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}
	}
}

//AssignProfile func
func assignProfile(profile *AppUserProfile, db *gorm.DB) {
	profile.CreationDate = time.Now()
	db.Where("user_id = ? and profile_id = ?", profile.UserID, profile.ProfileID).Save(profile)
}
