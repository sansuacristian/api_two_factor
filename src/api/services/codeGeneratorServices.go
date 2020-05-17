package services

import (
	"crypto/rand"
	"fmt"
)

//CodeGeneratorServices sirve para exportar el metodo que contiene a la capa superior o otros archivos dentro del codigo
type CodeGeneratorServices interface {
	Code() (string, error)
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
	token, err := generateRandomString(c.lenCode)
	if err != nil {
		return "", err
	}
	fmt.Println(token)

	return token, nil
}
