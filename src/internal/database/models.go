// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

type FileType string

const (
	FileTypeImage    FileType = "image"
	FileTypeDocument FileType = "document"
)

func (e *FileType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FileType(s)
	case string:
		*e = FileType(s)
	default:
		return fmt.Errorf("unsupported scan type for FileType: %T", src)
	}
	return nil
}

type NullFileType struct {
	FileType FileType `json:"file_type"`
	Valid    bool     `json:"valid"` // Valid is true if FileType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFileType) Scan(value interface{}) error {
	if value == nil {
		ns.FileType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FileType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFileType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FileType), nil
}

type UserType string

const (
	UserTypeAdmin UserType = "admin"
	UserTypeUser  UserType = "user"
)

func (e *UserType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserType(s)
	case string:
		*e = UserType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserType: %T", src)
	}
	return nil
}

type NullUserType struct {
	UserType UserType `json:"user_type"`
	Valid    bool     `json:"valid"` // Valid is true if UserType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserType) Scan(value interface{}) error {
	if value == nil {
		ns.UserType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserType), nil
}

type WebsiteConfigValue string

const (
	WebsiteConfigValueAllow    WebsiteConfigValue = "allow"
	WebsiteConfigValueDisallow WebsiteConfigValue = "disallow"
)

func (e *WebsiteConfigValue) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WebsiteConfigValue(s)
	case string:
		*e = WebsiteConfigValue(s)
	default:
		return fmt.Errorf("unsupported scan type for WebsiteConfigValue: %T", src)
	}
	return nil
}

type NullWebsiteConfigValue struct {
	WebsiteConfigValue WebsiteConfigValue `json:"website_config_value"`
	Valid              bool               `json:"valid"` // Valid is true if WebsiteConfigValue is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWebsiteConfigValue) Scan(value interface{}) error {
	if value == nil {
		ns.WebsiteConfigValue, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WebsiteConfigValue.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWebsiteConfigValue) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WebsiteConfigValue), nil
}

type File struct {
	ID         uuid.UUID     `json:"id"`
	Url        string        `json:"url"`
	Alt        string        `json:"alt"`
	Name       null.String   `json:"name"`
	Type       FileType      `json:"type"`
	UploadedAt time.Time     `json:"uploaded_at"`
	ProjectID  uuid.NullUUID `json:"project_id"`
}

type Project struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	ClientName   string        `json:"client_name"`
	LiveLink     null.String   `json:"live_link"`
	CodeLink     null.String   `json:"code_link"`
	StartDate    time.Time     `json:"start_date"`
	EndDate      null.Time     `json:"end_date"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	CoverImageID uuid.NullUUID `json:"cover_image_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	UserType  UserType  `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WebsiteConfiguration struct {
	ID                 uuid.UUID          `json:"id"`
	ConfigurationName  string             `json:"configuration_name"`
	ConfigurationValue WebsiteConfigValue `json:"configuration_value"`
	Description        null.String        `json:"description"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
	Active             bool               `json:"active"`
}
