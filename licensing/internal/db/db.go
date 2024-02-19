package db

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "fmt"
    "github.com/dchest/uniuri"
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

//retunr the db and the type User
func Init() *gorm.DB {
    key := generateApiKey()
    fmt.Println("Generated admin key: " + key)
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println("Error connecting to database")
        return nil
    }
    db.AutoMigrate(&User{}, &Key{})
    db.Create(&User{Username: "admin", ApiKey: key})
    return db
}
