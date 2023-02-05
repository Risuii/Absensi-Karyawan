package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt interface {
	HashPassword(plain string) (string, error)
	ComparePasswordHash(plain, hash string) bool
}

type BcryptImpl struct {
	HashCost int
}

func NewBcrypt(hashCost int) Bcrypt {
	return &BcryptImpl{
		HashCost: hashCost,
	}
}

func (bi *BcryptImpl) HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bi.HashCost)
	return string(bytes), err
}

func (bi *BcryptImpl) ComparePasswordHash(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
