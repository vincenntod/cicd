package actors

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mini-project-2/auth"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
func (r Repository) ShowActors() ([]Actors, error) {
	var actors []Actors

	err := r.Db.Find(&actors).Error
	return actors, err
}
func (r Repository) ShowActorsByName(c *gin.Context) ([]Actors, error) {
	var actors []Actors
	name := c.Param("name")
	err := r.Db.Model(&actors).Where("username LIKE ?", name+"%").Find(&actors).Error

	return actors, err

}
func (r Repository) EditActors(c *gin.Context) (Actors, error) {
	var actors Actors
	id := c.Param("id")
	if err := c.BindJSON(&actors); err != nil {
		c.JSON(500, gin.H{"error": err})
		return actors, err
	}
	err := r.Db.Where("id=?", id).Updates(&actors).Error
	return actors, err
}
func (r Repository) DeleteActors(c *gin.Context) (Actors, error) {
	var actors Actors
	id := c.Param("id")
	if err := c.BindJSON(&actors); err != nil {
		c.JSON(500, gin.H{"error": err})
		return actors, err
	}
	err := r.Db.Delete(&actors, id).Error
	return actors, err
}
func (r Repository) ShowRegister() ([]Actors, error) {
	var actors []Actors

	err := r.Db.Where("verified IS NULL OR active IS NULL").Find(&actors).Error
	return actors, err
}
func (r Repository) ApprovalRegister(c *gin.Context) (Actors, error) {
	var actors Actors
	id := c.Param("id")
	if err := c.BindJSON(&actors); err != nil {
		c.JSON(500, gin.H{"error": err})
		return actors, err
	}
	err := r.Db.Where("id=?", id).Updates(&actors).Error
	return actors, err
}
func (r Repository) RegisActors(c *gin.Context) (Actors, error) {
	var actors Actors
	if err := c.BindJSON(&actors); err != nil {
		c.JSON(500, gin.H{"error": err})
		return actors, err
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte((actors.Password)), bcrypt.DefaultCost)
	actors.Password = string(hashPassword)
	err := r.Db.Select("username", "password").Create(&actors).Error

	return actors, err
}
func (r Repository) LoginActors(c *gin.Context) (Actors, error) {
	var actors Actors
	if err := c.BindJSON(&actors); err != nil {
		return actors, err
	}
	var actor Actors
	if err := r.Db.Where("username= ?", actors.Username).First(&actor).Error; err != nil {
		c.JSON(401, gin.H{"error": "User tidak ditemukan"})
		return actor, err
	}
	if err := r.Db.Where("active = true").First(&actor).Error; err != nil {
		c.JSON(401, gin.H{"error": "User belum di setujui"})
		return actor, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(actor.Password), []byte(actors.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Password salah"})
		return actor, err
	}
	expTime := time.Now().Add(time.Minute * 60)

	claims := &auth.JWTClaim{
		Id:       actor.Id,
		Username: actor.Username,
		RoleId:   actor.RoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime)}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenGenerate, err := token.SignedString(auth.JWT_KEY)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return actor, err
	}
	c.SetCookie("gin_cookie", tokenGenerate, 3600, "/", "localhost", false, true)
	return actor, err
}
