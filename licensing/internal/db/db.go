package db

import (
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/sqlite"
    "github.com/dchest/uniuri"
    "os"
    "time"
    "github.com/fatih/color"
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
    Expires time.Time
}


func deleteExpiredKeys(db *gorm.DB) {
    //delete keys that have expired from about 7 or more days ago
    db.Where("expires < ?", time.Now().AddDate(0, 0, -7)).Delete(&Key{})
}

func DeleteExpiredKeys(db *gorm.DB) {
    for {
        deleteExpiredKeys(db)
        time.Sleep(24 * time.Hour)
    }
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
    //expires in a year 
    expires := time.Now().AddDate(1, 0, 0)
    db.Create(&Key{Key: key, User: user, Expires: expires})
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

func DeleteKey(db *gorm.DB, key string) {
    db.Where("key = ?", key).Delete(&Key{})
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

func OnlyAdmin(db *gorm.DB, key string) bool {
    var user User 
    db.Where("api_key = ?", key).First(&user)
    if user.Username == "admin" {
        return true
    }
    return false
}

func CreateUser(db *gorm.DB, username string) string {
    key := generateApiKey()
    db.Create(&User{Username: username, ApiKey: key})
    return string(key)
}

func DeleteUser(db *gorm.DB, key string) bool {
    var user User 
    err := db.Where("api_key = ?", key).First(&user).Error
    if err != nil {
        return false
    }
    if user.Username == "admin" {
        return false
    }
    db.Delete(&user)
    return true
}

func UpdateUserKey(db *gorm.DB, key string) string {
    var user User 
    err := db.Where("api_key = ?", key).First(&user).Error
    if err != nil {
        return ""
    }
    //if the user is the admin, don't update the key 
    if user.Username == "admin" {
        return string(user.ApiKey)
    }
    newKey := generateApiKey()
    user.ApiKey = newKey
    db.Save(&user)
    return string(newKey)
}

func GetApiKey(db *gorm.DB, username string) (string, bool) {
    var user User 
    err := db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return "", false
    }
    return string(user.ApiKey), true
}

func Init() *gorm.DB { 
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
        SkipDefaultTransaction: true,
    })
    if err != nil {
        color.Red("Failed to connect to database")
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
        color.Yellow("Admin user does not exist, creating one")
        db.Create(&User{Username: "admin", ApiKey: os.Getenv("ADMIN_KEY")})
        color.Green("Admin user created")
    }
    //run the function to delete expired keys in a goroutine
    go DeleteExpiredKeys(db)
    return db
}
