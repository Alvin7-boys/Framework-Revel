package models

import (
	"Tugas/app/util"
	"fmt"
)

type User struct {
	ID       int
	Name     string
	Age      int
	Address  string
	Password string
	Email    string
}

// GetAllUsers mengambil seluruh data user dari database
func GetAllUsers() ([]User, error) {
	// panggil InitDB() untuk membuat koneksi ke database
	db, err := util.InitDB()
	if err != nil {
		// handling jika terdapat error saat koneksi ke database
		return nil, err
	}
	defer db.Close()

	// lakukan operasi yang diinginkan dengan koneksi ke database
	// contoh: query ke database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		// handling jika terdapat error saat menjalankan query
		return nil, err
	}
	defer rows.Close()

	// buat slice untuk menampung hasil query
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Password, &user.Email)
		if err != nil {
			// handling jika terdapat error saat memindai hasil query
			return nil, err
		}
		users = append(users, user)
	}

	// kirim data ke controller
	return users, nil
}
func DeleteUserByName(name string) error {
	// buka koneksi database
	db, err := util.InitDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// query untuk mengambil ID user berdasarkan nama
	query := "SELECT ID FROM users WHERE name = ?"

	// eksekusi query
	var userID int
	err = db.QueryRow(query, name).Scan(&userID)
	if err != nil {
		return err
	}

	// query untuk menghapus transaksi berdasarkan userID
	query = "DELETE FROM tabel_transactions WHERE userID = ?"

	// eksekusi query
	_, err = db.Exec(query, userID)
	fmt.Println(userID)
	if err != nil {
		return err
	}

	// query untuk menghapus user berdasarkan nama
	query = "DELETE FROM users WHERE name = ?"

	// eksekusi query
	_, err = db.Exec(query, name)
	if err != nil {
		return err
	}

	return nil
}
func UpdateUserByID(id int, name string, age int, address string, password string, email string) (User, error) {
	// buka koneksi database
	db, err := util.InitDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	// mengambil user lama
	oldUser, err := GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	// cek nilai parameter, jika nil maka gunakan nilai lama
	if name == "" {
		name = oldUser.Name
	}
	if age == 0 {
		age = oldUser.Age
	}
	if address == "" {
		address = oldUser.Address
	}
	if password == "" {
		password = oldUser.Password
	}
	if email == "" {
		email = oldUser.Email
	}

	// query untuk mengupdate user berdasarkan ID
	query := "UPDATE users SET name = ?, age = ?, address = ?, passwordd = ?, email = ? WHERE ID = ?"

	// eksekusi query
	_, err = db.Exec(query, name, age, address, password, email, id)
	if err != nil {
		return User{}, err
	}
	newUser, err := GetUserByID(id)
	return newUser, nil
}
func GetUserByID(id int) (User, error) {
	// buka koneksi database
	db, err := util.InitDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	// query untuk mengambil user berdasarkan ID
	query := "SELECT * FROM users WHERE ID = ?"

	// eksekusi query
	var user User
	err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Password, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
func CreateUser(name string, age int, address string, password string, email string, id int) (User, error) {
	// buka koneksi database
	db, err := util.InitDB()
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	// query untuk memasukkan data user
	query := "INSERT INTO users (ID, Name, Age, Address, Passwordd, Email) VALUES (?, ?, ?, ?, ?, ?)"

	// eksekusi query
	_, err = db.Exec(query, id, name, age, address, password, email)
	if err != nil {
		return User{}, err
	}

	newUser, err := GetUserByID(id)
	return newUser, nil
}
func ValidateUser(email string, password string) error {
	// panggil InitDB() untuk membuat koneksi ke database
	db, err := util.InitDB()
	if err != nil {
		// handling jika terdapat error saat koneksi ke database
		return err
	}
	defer db.Close()

	// lakukan operasi yang diinginkan dengan koneksi ke database
	// contoh: query ke database
	query := "SELECT ID FROM users WHERE email = ?"

	// eksekusi query
	var userID int
	err = db.QueryRow(query, email).Scan(&userID)
	if err != nil {
		return err
	}

	// kirim data ke controller
	return nil
}
func Users(id int) ([]User, error) {
	// panggil InitDB() untuk membuat koneksi ke database
	db, err := util.InitDB()
	if err != nil {
		// handling jika terdapat error saat koneksi ke database
		return nil, err
	}
	defer db.Close()

	// lakukan operasi yang diinginkan dengan koneksi ke database
	// contoh: query ke database
	rows, err := db.Query("SELECT * FROM users = ? ", id)
	if err != nil {
		// handling jika terdapat error saat menjalankan query
		return nil, err
	}
	defer rows.Close()

	// buat slice untuk menampung hasil query
	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Password, &user.Email)
		if err != nil {
			// handling jika terdapat error saat memindai hasil query
			return nil, err
		}
		users = append(users, user)
	}

	// kirim data ke controller
	return users, nil
}
