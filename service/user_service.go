package service

import (
	"bookapi/entity"
	"bookapi/repository"
	"errors"
	//"math/rand/v2"
)

func GetAllUsers() []entity.User {
	return repository.GetAllUsers()
}

func Register(user entity.User) entity.User {
	//user.ID = rand.Uint64N(100)
	user = repository.InsertUser(user)
	return user
}

func Profile(id uint64) (entity.User,error) {
	if user, err := repository.GetUser(id); err == nil{
		return user, nil
	}
	return entity.User{}, errors.New("user doesn't exists")
}

func UpdateProfile(user entity.User, id uint64) (error){
	user.ID=id
	if err := repository.UpdateUser(user); err == nil{
		return nil 
	}
	return errors.New("user do not exist")
}

func DeleteAccount(identifiant uint64) (error){
	if err := repository.DeleteUser(identifiant); err == nil{
		return nil
	}
	return errors.New("user do not exist") 
}