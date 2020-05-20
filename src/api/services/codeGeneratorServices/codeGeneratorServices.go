package codeGeneratorServices

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/personal/api_two_factor/src/api/models"
)

var (
	token   string
	copy    string
	err     error
	copyErr error
)

//CodeGeneratorServices sirve para exportar el metodo que contiene a la capa superior o otros archivos dentro del codigo
type CodeGeneratorServices interface {
	Code() (string, error)
	Validation(userToken models.Token) (string, error)
}

//codeGeneratorServices genera la dependencia y se usa dentro del mismo codigo
type codeGeneratorServices struct {
	lenCode int
}

//NewCodeServices constructor de este servicio
func NewCodeServices() CodeGeneratorServices {
	return &codeGeneratorServices{lenCode: 32}
}

//GenerateRandomBytes generar una secuencia de bits aleatorios
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString define el contenido del codigo (string) final
func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	bytes, err := generateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		//la siguiente linea utiliza el valor de b que proviene de bytes para obtener un valor de la constante letters
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

//Code genera el codigo
func (c *codeGeneratorServices) Code() (string, error) {
	// Example: la extensión del codigo sera de 32 bytes
	token, err = generateRandomString(c.lenCode)
	if err != nil {
		return "", err
	}
	errSendEmail := sendEmail(token)
	if errSendEmail != nil {
		return "", errSendEmail
	}

	return token, nil
}

func (c *codeGeneratorServices) Validation(userToken models.Token) (string, error) {

	if token != userToken.Token {
		return "", errors.New(" El token no coincide por favor intente de nuevo")
	}
	return "Token validado exitosamente continue con su compra", nil
}

//SendEmail envia el token al correo electronico
func sendEmail(code string) error {
	from := "notificacionesapideseguridad@gmail.com"
	to := "cristian.sanchezs@unicafam.edu.co"

	// servidor SMTP
	host := "smtp.gmail.com"
	//Autenticación de la cuenta que envia (from)
	auth := smtp.PlainAuth("", from, "estoeslapruebadelapi2020", host)
	message := fmt.Sprintf("Su codigo de autenticación es &s", code)

	if err := smtp.SendMail(host+":25", auth, from, []string{to}, []byte(message)); err != nil {
		fmt.Println("Error SendMail: ", err)
		return err
	}
	fmt.Sprintln("Correo enviado")
	return nil

}
