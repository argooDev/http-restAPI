package store

import "github.com/argooDev/http-restAPI/internal/app/model"

// Тип репозитория
type UserRepository struct {
	store *Store // Ссылка на главное хранилище
}

// Принимает на вход модель, возвращает модель
func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	// Передаем sql запрос, returning id - psql по умолчанию не возвращает id, поэтому делаю так
	// Scan - после того, как запрос возвращает строку, он сможет смапить в переданные аргументы(&u.ID)
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	// return user
	return u, nil
}

// Необходим при авторизации, принимает email, возвращает модель/ошибку
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	// Инициализируем юзера, в которого запишем данные из БД
	u := &model.User{}
	// QueryRow возвращает всегда 1 строку
	// Передаем в том же порядке параметры в Scan
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	// Если ошибок нет - возращаем заполненного юзера
	return u, nil
}