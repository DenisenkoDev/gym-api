package cryp

import "golang.org/x/crypto/bcrypt"

// GenerateHashFromPassword генерирует хеш из исходного пароля
func GenerateHashFromPassword(password string) (string, error) {
	// Генерируем хеш с использованием bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CompareHashes сравнивает два хеша (но это не рекомендуется, см. пояснения)
func CompareHashes(storedHash, providedHash string) bool {
	// Проверяем, совпадают ли строки напрямую
	return storedHash == providedHash
}

// ComparePasswordWithHash правильно сравнивает пароль с сохранённым хешем
func ComparePasswordWithHash(password, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	return err == nil
}
