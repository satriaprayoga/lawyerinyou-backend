package models

import "time"

//LoginForm :
type LoginForm struct {
	Account  string `json:"account" valid:"Required"`
	UserType string `json:"user_type" valid:"Required"`
	Password string `json:"pwd" valid:"Required"`
	FcmToken string `json:"fcm_token" valid:"Required"`
}

// RegisterForm :
type RegisterForm struct {
	Name        string    `json:"name" valid:"Required"`
	BirthOfDate time.Time `json:"birth_of_date"`
	ResetPasswd
}

// ForgotForm :
type ForgotForm struct {
	Account  string `json:"account" valid:"Required"`
	UserType string `json:"user_type" valid:"Required"`
}

// ResetPasswd :
type ResetPasswd struct {
	Account       string `json:"account" valid:"Required"`
	UserType      string `json:"user_type" valid:"Required"`
	Passwd        string `json:"pwd" valid:"Required"`
	ConfirmPasswd string `json:"confirm_pwd" valid:"Required"`
}

type VerifyForm struct {
	Account    string `json:"account" valid:"Required"`
	UserType   string `json:"user_type" valid:"Required"`
	VerifyCode string `json:"verify_code" valid:"Required"`
	FcmToken   string `json:"fcm_token,omitempty"`
}

type DataLogin struct {
	UserID   int       `json:"user_id" db:"user_id"`
	Password string    `json:"pwd" db:"pwd"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Telp     string    `json:"telp" db:"telp"`
	JoinDate time.Time `json:"join_date" db:"join_date"`
	UserType string    `json:"user_type" db:"user_type"`
	//FileID   sql.NullInt64  `json:"file_id" db:"file_id"`
	//FileName sql.NullString `json:"file_name" db:"file_name"`
	//FilePath sql.NullString `json:"file_path" db:"file_path"`
}
