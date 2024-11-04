package model

import "golang.org/x/crypto/bcrypt"

// User соответствует схеме в Бд
type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
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
