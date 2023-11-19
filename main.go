package main

import (
	// Local package
	"fmt"
	"net/http"
	"strconv"
	"strings"

	// framework & third party
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser Siteuser
	err := c.BindJSON(&newUser)
	//var Createuser Siteuser
	// check error existing
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Faild to Signup, Error has occured",
		})
		return
	}

	// Email validation
	checkValidEmail := newUser.Email
	var FirstValidEmail = strings.Contains(checkValidEmail, "@")
	var SecondValidEmail = strings.Contains(checkValidEmail, ".com")
	var TotalEmailValidation = FirstValidEmail && SecondValidEmail

	// Password validation
	checkValidPassword := newUser.Password
	var ValidPasswordContainsOne = strings.Contains(checkValidPassword, "@")
	var ValidPasswordContainsTwo = strings.Contains(checkValidPassword, "!")
	var ValidPasswordContainsThree = strings.Contains(checkValidPassword, "#")
	var ValidPasswordContainsFour = strings.Contains(checkValidPassword, "$")
	var ValidPasswordContainsFive = strings.Contains(checkValidPassword, "%")
	var ValidPasswordContainsSix = strings.Contains(checkValidPassword, "^")
	var ValidPasswordContainsSeven = strings.Contains(checkValidPassword, "&")
	var ValidPasswordContainsEight = strings.Contains(checkValidPassword, "*")
	var TotalValidPasswordContains = ValidPasswordContainsOne || ValidPasswordContainsTwo || ValidPasswordContainsThree || ValidPasswordContainsFour || ValidPasswordContainsFive || ValidPasswordContainsSix || ValidPasswordContainsSeven || ValidPasswordContainsEight
	var ValidPasswordLen = len(checkValidPassword) > 6
	var TotalPasswordValidation = TotalValidPasswordContains && ValidPasswordLen

	// Username validation
	var ValidUsername = len(newUser.Username) > 6

	// Check total validation
	if TotalEmailValidation && ValidUsername && TotalPasswordValidation {
		// users = append(Createuser, newUser)
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
			"data":    newUser,
		})
		Db.Create(&newUser)
	} else {
		if !TotalEmailValidation {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Faild to Signup, Email has not valid format",
			})
		}
		if !ValidUsername {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Faild to Signup, Username is less than 6 char",
			})
		}
		if !TotalPasswordValidation {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Faild to Signup,Password has not valid format, it must has one of this char: '!@#$%^&*', and more than 6 char",
			})
		}
		return
	}
}

func ReadUsers(c *gin.Context) {
	var Readusers []Siteuser
	if Db.Find(&Readusers).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Struct is empty",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok Read",
		"data":    Db.Find(&Readusers),
	})

}

func ReadUser(c *gin.Context) {
	var id = c.Param("id")
	var Readuser Siteuser
	if Db.First(&Readuser, id).RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user ID not found",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Data":    Db.First(&Readuser, id),
			"message": "This is user",
		})
	}
}

func UpdateUser(c *gin.Context) {
	var id, err1 = strconv.Atoi(c.Param("id"))

	if err1 != nil {
		fmt.Println(err1)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	var newUser Siteuser
	var err2 = c.BindJSON(&newUser)
	var Updateuser []Siteuser
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to update, error has occured",
		})
		return
	}

	// Email validation
	checkValidEmail := newUser.Email
	var FirstValidEmail = strings.Contains(checkValidEmail, "@")
	var SecondValidEmail = strings.Contains(checkValidEmail, ".com")
	var TotalEmailValidation = FirstValidEmail && SecondValidEmail

	// Password validation
	checkValidPassword := newUser.Password
	var ValidPasswordContainsOne = strings.Contains(checkValidPassword, "@")
	var ValidPasswordContainsTwo = strings.Contains(checkValidPassword, "!")
	var ValidPasswordContainsThree = strings.Contains(checkValidPassword, "#")
	var ValidPasswordContainsFour = strings.Contains(checkValidPassword, "$")
	var ValidPasswordContainsFive = strings.Contains(checkValidPassword, "%")
	var ValidPasswordContainsSix = strings.Contains(checkValidPassword, "^")
	var ValidPasswordContainsSeven = strings.Contains(checkValidPassword, "&")
	var ValidPasswordContainsEight = strings.Contains(checkValidPassword, "*")
	var TotalValidPasswordContains = ValidPasswordContainsOne || ValidPasswordContainsTwo || ValidPasswordContainsThree || ValidPasswordContainsFour || ValidPasswordContainsFive || ValidPasswordContainsSix || ValidPasswordContainsSeven || ValidPasswordContainsEight
	var ValidPasswordLen = len(checkValidPassword) > 6
	var TotalPasswordValidation = TotalValidPasswordContains && ValidPasswordLen

	// Username validation
	var ValidUsername = len(newUser.Username) > 6

	// Check total validation
	if TotalEmailValidation && ValidUsername && TotalPasswordValidation {
		Updateuser[id].Username = newUser.Username
		Updateuser[id].First_name = newUser.First_name
		Updateuser[id].Last_name = newUser.Last_name
		Updateuser[id].Email = newUser.Email
		Updateuser[id].Password = newUser.Password

		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
			"data":    newUser,
		})
		Db.Save(&Updateuser[id])
	} else {
		if !TotalEmailValidation {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Faild to Signup, Email has not valid format",
			})
		}
		if !ValidUsername {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Faild to Signup, Username is less than 6 char",
			})
		}
		if !TotalPasswordValidation {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Faild to Signup,Password has not valid format, it must has one of this char: '!@#$%^&*', and more than 6 char",
			})
		}
		return
	}
}

func DeleteUser(c *gin.Context) {
	var id, err1 = strconv.Atoi(c.Param("id"))

	var deleteuser Siteuser

	if err1 != nil {
		fmt.Println(err1)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	// firstHalf := users[:id]
	// secondHalf := users[id+1:]
	// users = append(firstHalf, secondHalf...)
	Db.Delete(&deleteuser, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})

}

// main function

func main() {
	r := gin.Default()

	r.POST("/create/", CreateUser)
	r.GET("/read/", ReadUsers)
	r.GET("/read/:id", ReadUser)
	r.PUT("/update/:id", UpdateUser)
	r.DELETE("/delete/:id", DeleteUser)

	Database()
	r.Run(":8000")
}
