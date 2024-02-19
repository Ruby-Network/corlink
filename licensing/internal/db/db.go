package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "fmt"
    "github.com/dchest/uniuri"
    "os"
)

type User struct {
    gorm.Model
    Username string
    ApiKey string
}

type Key struct {
    gorm.Model
    Key string
    User User `gorm:"foreignKey:ID"`
}

func GenerateKey(db *gorm.DB, user User) string {
    key := uniuri.NewLen(6)
    //get all the keys for the user to make sure we don't generate the same key twice
    var keys []Key 
    db.Where("ID = ?", user.ID).Find(&keys)
    for _, k := range keys {
        if k.Key == key {
            return GenerateKey(db, user)
        }
    }
    db.Create(&Key{Key: key, User: user})
    fmt.Println("Generated key: " + key)
    return string(key)
}

func VerifyKey(db *gorm.DB, key string) bool {
    var k Key
    db.Where("key = ?", key).First(&k)
    if k.ID == 0 {
        return false
    }
    return true
}

func generateApiKey() string {
    key := uniuri.NewLen(64)
    return string(key)
}

func GetUserByApiKey(db *gorm.DB, key string) User {
    var user User 
    db.Where("api_key = ?", key).First(&user)
    return user
}

func Init() *gorm.DB { 
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("Error connecting to database")
        return nil
    }
    //if the tables don't exist, create them
    if db.Migrator().HasTable(&User{}) == false {
        db.Migrator().CreateTable(&User{})
    }
    if db.Migrator().HasTable(&Key{}) == false {
        db.Migrator().CreateTable(&Key{})
    }
    //create the admin user if it doesn't exist
    if db.Where("username = ?", "admin").First(&User{}).RowsAffected == 0 {
        fmt.Println("Creating admin user")
        db.Create(&User{Username: "admin", ApiKey: os.Getenv("ADMIN_KEY")})
    }
    return db
}
