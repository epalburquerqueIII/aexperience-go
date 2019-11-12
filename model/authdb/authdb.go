// this acts as a simple in memory db to handle users and refresh tokens.
// replace these functions with actual calls to your db

package authdb

import (
	"errors"
	"log"

	"../../controller"
	"../../model"
	"../../util"
	"golang.org/x/crypto/bcrypt"
)

// create a database of users
// the map key is the uuid
var users = map[string]model.User{}

// create a database of refresh tokens
// map key is the jti (json token identifier)
// the val doesn't represent anything but could be used to hold "valid", "revoked", etc.
var refreshTokens map[string]string

// InitDB inicializa mapa
func InitDB() {
	refreshTokens = make(map[string]string)
}

// StoreUser password is hashed before getting here
func StoreUser(email string, password string, role int, username string, userid int) (uuid string, err error) {
	uuid, err = util.GenerateRandomString(32)
	if err != nil {
		return "", err
	}

	// check to make sure our uuid is unique
	u := model.User{}
	for u != users[uuid] {
		uuid, err = util.GenerateRandomString(32)
		if err != nil {
			return "", err
		}
	}

	// generate the bcrypt password hash
	passwordHash, hashErr := generateBcryptHash(password)
	if hashErr != nil {
		err = hashErr
		return
	}

	users[uuid] = model.User{email, passwordHash, username, role, userid}

	return uuid, err
}

// DeleteUser Anula usuario
func DeleteUser(uuid string) {
	delete(users, uuid)
}

// FetchUserByID Anula usuario
func FetchUserByID(uuid string) (model.User, error) {
	u := users[uuid]
	blankUser := model.User{}

	if blankUser != u {
		// found the user
		return u, nil
	} else {
		return u, errors.New("User not found that matches given uuid")
	}
}

// FetchUserByEmail returns the user and the userId or an error if not found
func FetchUserByEmail(email string) (model.User, string, error) {
	// so of course this is dumb, but it's just an example
	// your db will be much faster!

	located, userid, username, k, userRol := controller.GetUsuario(email)
	if located {
		return model.User{email, k, username, userRol, userid}, k, nil
	}

	for k, v := range users {
		if v.Email == email {
			return v, k, nil
		}
	}

	return model.User{}, "", errors.New("User not found that matches given email")
}

// StoreRefreshToken Refresh Token
func StoreRefreshToken() (jti string, err error) {
	jti, err = util.GenerateRandomString(32)
	if err != nil {
		return jti, err
	}

	// check to make sure our jti is unique
	for refreshTokens[jti] != "" {
		jti, err = util.GenerateRandomString(32)
		if err != nil {
			return jti, err
		}
	}

	refreshTokens[jti] = "valid"

	return jti, err
}

// DeleteRefreshToken Delete Token
func DeleteRefreshToken(jti string) {
	delete(refreshTokens, jti)
}

// CheckRefreshToken Check Token
func CheckRefreshToken(jti string) bool {
	return refreshTokens[jti] != ""
}

// LogUserIn Login
func LogUserIn(email string, password string) (model.User, string, error) {
	user, uuid, userErr := FetchUserByEmail(email)
	log.Println(user, uuid, userErr)
	if userErr != nil {
		return model.User{}, "", userErr
	}

	return user, uuid, checkPasswordAgainstHash(user.PasswordHash, password)
}

func generateBcryptHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash[:]), err
}

func checkPasswordAgainstHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
