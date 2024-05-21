package bd

import (
	//Importaciones de go (vienen incluidas al instalar)
	"database/sql"
	"fmt"
	"os"

	//importaciones externas (descargadas)
	_ "github.com/go-sql-driver/mysql"

	//importaciones personalizadas (creadas desde cero)
	"github.com/PedroAntonioKira/ecommerceEscomPrincipal/models"
	"github.com/PedroAntonioKira/ecommerceEscomPrincipal/secretsmanager"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretsmanager.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	//Verificamos que todo este bien y no se produzca un error

	//Verificamos que nos conectamos a la bd
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()

	//Verificamos que podamos enviar ping (haya conección) a la bd
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión Exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "ecommerceEscom" //Nombre que aparece en HeidiSQL.
	//dsName
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser,
		authToken,
		dbEndpoint,
		dbName)

	fmt.Println(dsn)

	return dsn
}
