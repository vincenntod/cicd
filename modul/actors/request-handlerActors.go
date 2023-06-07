package actors

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"mini-project-2/auth"
	"net/http"
	"time"
)

type RequestHandler struct {
	ctrl *Controller
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}
func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{
		ctrl: ctrl,
	}
}

func (h RequestHandler) ShowActors(c *gin.Context) {
	res, err := h.ctrl.ShowActors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h RequestHandler) ShowActorsByName(c *gin.Context) {
	res, err := h.ctrl.ShowActorsByName(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)

}
func (h RequestHandler) EditActors(c *gin.Context) {
	res, err := h.ctrl.EditActors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) DeleteActors(c *gin.Context) {
	_, err := h.ctrl.DeleteActors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Terhapus"})
}

func (h RequestHandler) ShowRegister(c *gin.Context) {
	res, err := h.ctrl.ShowRegister()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) ApprovalRegister(c *gin.Context) {
	res, err := h.ctrl.ApprovalRegister(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) RegisActors(c *gin.Context) {
	res, err := h.ctrl.RegisActors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h RequestHandler) LoginActors(c *gin.Context) {
	_, err := h.ctrl.LoginActors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login Berhasil"})
}
func (h RequestHandler) Logout(c *gin.Context) {
	c.SetCookie("gin_cookie", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "Logout Berhasil"})
}
func (h RequestHandler) MiddlewareAdmin(c *gin.Context) {
	tokenString, err := c.Cookie("gin_cookie")
	if err != nil {
		c.AbortWithStatus(401)
		c.JSON(401, gin.H{"message": "Silahkan Login"})
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return auth.JWT_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(401)
			c.JSON(401, gin.H{"message": "Silahkan Login Kembali"})
		}
		var actor Actors
		h.ctrl.useCase.repo.Db.First(&actor, claims["Id"])
		if actor.Id == 0 {
			c.AbortWithStatus(401)
		}
		c.Set("user", actor)
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}

}
func (h RequestHandler) MiddlewareSuperAdmin(c *gin.Context) {

	tokenString, err := c.Cookie("gin_cookie")
	if err != nil {
		c.JSON(401, gin.H{"message": "Silahkan Login"})
		c.AbortWithStatus(401)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return auth.JWT_KEY, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{"message": "Silahkan Login Kembali"})
			c.AbortWithStatus(401)
		}
		var actor Actors
		h.ctrl.useCase.repo.Db.First(&actor, claims["Id"])
		if actor.Id == 0 {
			c.AbortWithStatus(401)
		}
		if actor.RoleId != 1 {
			c.JSON(403, gin.H{"message": "forbidden"})
			c.AbortWithStatus(401)
		}
		c.Set("user", actor)
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}

}
