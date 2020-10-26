package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/anovanmaximuz/go-jwt/structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{	
			"code":400,		
			"data": err.Error(),
		}
	} else {
		result = gin.H{
			"code":200,
			"data": person,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"code":400,
			"data": nil,			
		}
	} else {
		result = gin.H{
			"code":200,			
			"data": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	password  := c.PostForm("password")
	person.First_Name = first_name
	person.Last_Name = last_name
	if password != "" {
     		hash, err := HashPassword(password)
     		if err != nil {
        		return
     		}
     		password = hash
  	}



	person.Password = password
	idb.DB.Create(&person)
	result = gin.H{
		"code":200,
		"data": person,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"code":400,
			"data": "data not found",
		}
	}
	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"code":301,
			"data": "update failed",
		}
	} else {
		result = gin.H{
			"code":200,
			"data": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"code":301,
			"data": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"code":302,
			"data": "delete failed",
		}
	} else {
		result = gin.H{
			"code":200,
			"data": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
