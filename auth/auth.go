package auth

import (
	//Importaciones de go (vienen incluidas al instalar)
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipal/models"
)

//La parte de la estructura la implemente en models "type TokenJSON struct"

func ValidoToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("El token no es valido, no viene con tres partes")
		return false, nil, "El token no es valido, no viene con tres partes"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])

	//validamos que se pueda decodificar correctamente la info
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token: " + err.Error())
		return false, err, err.Error()
	}

	var tkj models.TokenJSON

	//Convertimos la información de userInfo decodificada a una estructura de GO "models.TokenJSON"
	err = json.Unmarshal(userInfo, &tkj)

	//Verificamos si hay un error ya que no se decodifico la estructura correctamente
	if err != nil {
		fmt.Println("No se puede decodificar en la estructura JSON models.TokenJSON: " + err.Error())
		return false, err, err.Error()
	}

	//Obtenemos la fecha actual
	ahora := time.Now()

	//Obtenemos la fecha de expiración del token de cognito
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(ahora) {
		fmt.Println("Fecha expiración token = " + tm.String())
		fmt.Println("Token Expirado !")
		return false, err, "Token Expirado !!"
	}

	return true, nil, string(tkj.Username)

}
