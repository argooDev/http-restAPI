package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User соответствует схеме в Бд
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email), // Указываем значение для валидации, далее указываем правила валидации

		// При чтении юзера из БД, мы не записываем ему password, мы записываем шифрованный пароль в encryptedPassword
		// Это можно посмотреть в internal/app/store/userrepository FindByEmail
		// Из-за этого password будет пустым и модель пользователя будет невалидной при следующем сохранении юзера в БД
		// validation.By() фиксит это. В нее добавляем кастомные проверки, они будут в validations.go
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

// Callback функция, вызывается каждый раз при попытке сохранения пользователя в БД с помощью Create
func (u *User) BeforeCreate() error {

	// Проверка пароля на пустоту, если не пуст - передаем пароль в encryptString
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		// Если ошибок нет - записываем зашифрованную строку в EncryptedPassword
		u.EncryptedPassword = enc
	}

	return nil
}

// Принимает и возвращает строки
func encryptString(s string) (string, error) {
	// С помощью bcrypt.GenerateFromPassword, она принимает строку, приведенную в массив байтов, втором bcrypt.MinCos

	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	// Если ошибок нет - возвращаем шифрованные байты и приводим их обратно в строку
	return string(b), err
}
