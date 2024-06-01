package db

import (
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
    "github.com/dchest/uniuri"
    "os"
    "time"
    "fmt"
    "github.com/fatih/color"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    ApiKey string
    Keys []Key `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Key struct {
    gorm.Model
    Key string
    UserID uint
    Used bool
    Expires time.Time
}


func deleteExpiredKeys(db *gorm.DB) {
    //delete keys that have expired from about 7 or more days ago
    db.Where("expires < ?", time.Now().AddDate(0, 0, -7)).Unscoped().Delete(&Key{})
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
    db.Where("user_id = ?", user.ID).Find(&keys)
    for _, k := range keys {
        if k.Key == key {
            return GenerateKey(db, user)
        }
    }
    //expires in a year 
    expires := time.Now().AddDate(1, 0, 0)
    //expires in 5 minutes for testing
    //expires := time.Now().Add(2 * time.Minute)
    db.Create(&Key{Key: string(key), UserID: user.ID, Expires: expires})
    return string(key)
}

func VerifyKey(db *gorm.DB, key string) bool {
    var k Key
    db.Where("key = ?", key).First(&k)
    if k.ID == 0 {
        return false
    }
    if k.Used {
        return false
    }
    db.Model(&k).Update("used", true)
    return true
}

func DeleteKey(db *gorm.DB, key string) {
    db.Where("key = ?", key).Unscoped().Delete(&Key{})
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

func CreateUser(db *gorm.DB, username string) (error, string) {
    key := generateApiKey()
    err := db.Create(&User{Username: username, ApiKey: key}).Error
    if err != nil {
        error := fmt.Errorf("Username already exists")
        return error, ""
    }
    return nil, string(key)
}

func DeleteUser(db *gorm.DB, username string) bool {
    var user User 
    err := db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return false
    }
    if user.Username == "admin" {
        return false
    }
    err = db.Unscoped().Delete(&user).Error 
    if err != nil {
        return false
    }
    return true
}

func UpdateUserKey(db *gorm.DB, username string) string {
    var user User 
    err := db.Where("username = ?", username).First(&user).Error
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

func Init(lite bool) *gorm.DB {
    //blank assignment to avoid error 
    var db *gorm.DB
    var err error
    if lite {
        db, err = gorm.Open(sqlite.Open("corlink.db"), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Silent),
            SkipDefaultTransaction: true,
        })
        if err != nil { color.Red("Failed to connect to database"); return nil }
    } else {
        dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
        db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Silent),
            SkipDefaultTransaction: true,
        })
        if err != nil {
            color.Red("Failed to connect to database")
            return nil
        }
    }
    //if the tables don't exist, create them
    db.AutoMigrate(&User{}, &Key{})
    if db.Where("username = ?", "admin").First(&User{}).RowsAffected == 0 {
        color.Yellow("Admin user does not exist, creating one")
        db.Create(&User{Username: "admin", ApiKey: os.Getenv("ADMIN_KEY")})
        color.Green("Admin user created")
    }
    //run the function to delete expired keys in a goroutine
    go DeleteExpiredKeys(db)
    return db
}
