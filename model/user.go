package model


import(
	"github.com/ChanchalS7/practice_api/database"
	"html"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//user model
type User struct{
	gorm.Model
	ID			unit		`gorm:"primary_key"`
	RoleID		unit		`gorm:"not null; DEFAULT:3" json:"role_id"`
	Username	string		`gorm:"size:255; not null; unique" json:"email"`
	Email       string		`gorm:"size:255;not null; unique" json:"email"`
	password	string		`gorm:"size:255;not null" json:"-"`
	Role		Role 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"json:"-"`		
}

//Save user details
func (user *User) Save() (*User, error){
	err:= database.Db.Create(&user).Error
	if err!=nil{
		return &User{},err
	}
	return user,nil
}

//Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error{
	passwordHash, err:= bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
	if err!=nil{
		return err
	}
	user.password= string(passwordHash)
	user.Username=html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}
// Get all users
func GetUsers(User *[]User)(err error){
	err = database.Db.Find(User).Error
	if err!=nil{
		return err
	}
	return nil
}

//Validate user password

func (user *User) ValidateUserPassword( password string) error{
	return bcrypt.CompareHashAndPassword([]byte(user.password),[]byte(password))

}



//Get user by id
func GetUserById(id uint)(User, error){
	var user User

	err:=database.Db.Where("id=?",id).Find(&user).Error
	if err!=nil{
		return User{},err
	}
	return user,nil
}
//Update user

func UpdateUser(User *User)(err error){
	err = database.Db.Omit("password").Updates(User).Error
	if err!=nil{
		return nil
	}
	return nil
}



