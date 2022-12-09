package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Token struct {
	UserName string `json:"username"`
	VimTime  string `json:"vim_time"`
	jwt.RegisteredClaims
}

func TestUseJwt() {

	mySigningKey := []byte("XingChenWq")

	tk := Token{
		UserName: "user1",
		VimTime:  "1670510765",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "wq",
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 8)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		println(ss)
		log.Printf("err: %s \n", err)
		return
	}
	println(ss)

	t, err := jwt.ParseWithClaims(ss, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Println("十点半", err)
	}
	if c, ok := t.Claims.(*Token); ok && t.Valid {
		fmt.Printf("%#v\n", c)
		println(c.UserName)
	} else {
		fmt.Println("十点半就是", err)
	}
}
