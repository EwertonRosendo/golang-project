package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
)

// access token
func CreateToken(userID uint64) (string, error){

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
} 

func TokenValidation(r *http.Request) error {

	stringToken := extractToken(r)
	token, err := jwt.Parse(stringToken, returnVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")

}

func ExtractUserID(r *http.Request) (uint64, error) {
	stringToken := extractToken(r)
	token, err := jwt.Parse(stringToken, returnVerificationKey)

	if err != nil {
		return 0, err
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		if is_access_token_valid := VerifyAccessToken(r, userID); is_access_token_valid != nil {
			return 0, errors.New("invalid token")	
		}
		return userID, nil
	}
	return 0, errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error){
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("wrong authentication method detected: %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func VerifyAccessToken(r *http.Request, userID uint64) error {
	cookie, err := r.Cookie(fmt.Sprintf("user_%d", userID))
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
        default:
            log.Println(err)
        }
    }
	stringAccessToken := cookie.Value
	token, err := jwt.Parse(stringAccessToken, returnVerificationKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid access token")
}

func CreateAccessTokenCookie(w http.ResponseWriter, r *http.Request, userID uint64) {
	token_name := fmt.Sprintf("user_%d", userID)
	access_token, err := CreateAcessToken(userID)
	if err != nil{
		errors.New("somewthing went wrong during the token creation")
	}

	cookie := http.Cookie{
        Name:     token_name,
        Value:    access_token,
        Path:     "/",
        MaxAge:   2450000,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    }

	http.SetCookie(w, &cookie)
}

func CreateAcessToken(userID uint64) (string, error){

	permissions_access_token := jwt.MapClaims{}
	permissions_access_token["authorized"] = true
	permissions_access_token["exp"] = time.Now().Add(time.Hour * 43200).Unix()
	permissions_access_token["user_id"] = userID

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions_access_token)

	return access_token.SignedString([]byte(config.SecretKey))
} 

func RefreshToken(r *http.Request, userID uint64) (string, error) {
	if err := TokenValidation(r); err != nil{
		
		if err := VerifyAccessToken(r, userID); err != nil {
			return "", err
		}
		token, err := CreateToken(userID)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	return extractToken(r), nil

}