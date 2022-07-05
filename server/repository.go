package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	Database     = "mysql"
	DbUsername   = "root"
	DbPassword   = "Ajay1234#"
	DbHost       = "127.0.0.1"
	DbPort       = "3306"
	DatabaseName = "secret_vault_db"
	DataSource   = DbUsername + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DatabaseName
)

func insertUser(user *User) (*User, error) {
	db, err := sql.Open(Database, DataSource)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO User (uid, email, username, first_name, last_name, password) "+
		"VALUES (UUID_TO_BIN(UUID()), '%s', '%s', '%s', '%s', '%s')",
		user.Email, user.Username, user.FirstName, user.LastName, user.Password)
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func findUserByUsername(username string) (*DbUser, error) {
	db, err := sql.Open(Database, DataSource)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := fmt.Sprintf(
		"SELECT BIN_TO_UUID(uid), email, username, first_name, last_name, password, created_at FROM User"+
			" WHERE username = '%s'", username)
	row := db.QueryRow(query)
	if err != nil {
		return nil, err
	}
	var userData DbUser
	err = row.Scan(
		&userData.UID, &userData.Email, &userData.Username, &userData.FirstName, &userData.LastName,
		&userData.Password, &userData.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &userData, nil
}

func getCredentialsFromDb(uid, key string) (*Credentials, error) {
	db, err := sql.Open(Database, DataSource)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(
		"SELECT cred_key, "+
			"cred_value "+
			"FROM KV_Store "+
			"WHERE "+
			"BIN_TO_UUID(user_id) = '%s' "+
			"AND "+
			"cred_key = '%s' "+
			"AND "+
			"destroyed IS NOT TRUE",
		uid, key,
	)
	res := db.QueryRow(query)
	var queryRes Credentials
	err = res.Scan(&queryRes.Key, &queryRes.Value)
	if err != nil {
		return nil, err
	}

	return &queryRes, nil
}

func insertCredentialsInDb(uid, key, value string) error {
	db, err := sql.Open(Database, DataSource)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"INSERT INTO KV_Store (id, user_id, cred_key, cred_value) "+
			"VALUES ("+
			"UUID_TO_BIN(UUID()),"+
			"UUID_TO_BIN('%s'),"+
			"'%s',"+
			"'%s'"+
			");", uid, key, value,
	)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
