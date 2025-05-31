package user

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		err = fmt.Errorf("error on generate hash password %d", err)
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func GenerateUUID() (uuid.UUID, error) {
	generate, err := uuid.NewRandom()
	if err != nil {
		return generate, err
	}

	return generate, err
}

func MapFormToStruct(form map[string][]string, dst interface{}) error {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to struct, got %T", dst)
	}

	dstVal = dstVal.Elem()
	dstType := dstVal.Type()

	for i := 0; i < dstType.NumField(); i++ {
		field := dstType.Field(i)
		formTag := field.Tag.Get("form")
		if values, ok := form[formTag]; ok && len(values) > 0 {
			fieldVal := dstVal.FieldByName(field.Name)
			if fieldVal.CanSet() && fieldVal.Kind() == reflect.String {
				fieldVal.SetString(values[0])
			}
		}
	}
	return nil
}

func GetUUID() string {
	return uuid.New().String()
}
