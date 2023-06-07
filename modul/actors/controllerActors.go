package actors

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Controller struct {
	useCase *UseCase
}
type ReadResponse struct {
	Data []ActorsResponse `json:"data"`
}
type ActorsResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
	Verified  string    `json:"verified"`
	Active    string    `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type RegisterApprovalResponse struct {
	Id           int    `json:"id"`
	AdminId      int    `json:"admin_id"`
	SuperAdminId int    `json:"super_admin_id"`
	Status       string `json:"status"`
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}
func (c Controller) ShowActors() (*ReadResponse, error) {
	actors, err := c.useCase.ShowActors()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, actor := range actors {
		item := ActorsResponse{
			Id:        actor.Id,
			Username:  actor.Username,
			Password:  actor.Password,
			RoleId:    actor.RoleId,
			Verified:  actor.Verified,
			Active:    actor.Active,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) ShowActorsByName(ctx *gin.Context) (*ReadResponse, error) {
	actors, err := c.useCase.ShowActorsByName(ctx)
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, actor := range actors {
		item := ActorsResponse{
			Id:        actor.Id,
			Username:  actor.Username,
			Password:  actor.Password,
			RoleId:    actor.RoleId,
			Verified:  actor.Verified,
			Active:    actor.Active,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) EditActors(ctx *gin.Context) (Actors, error) {
	actors, err := c.useCase.EditActors(ctx)
	if err != nil {
		return actors, err
	}

	return actors, nil
}
func (c Controller) DeleteActors(ctx *gin.Context) (Actors, error) {
	actors, err := c.useCase.DeleteActors(ctx)
	if err != nil {
		return actors, err
	}

	return actors, nil
}
func (c Controller) ShowRegister() (*ReadResponse, error) {
	actors, err := c.useCase.ShowRegister()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, actor := range actors {
		item := ActorsResponse{
			Id:        actor.Id,
			Username:  actor.Username,
			Password:  actor.Password,
			RoleId:    actor.RoleId,
			Verified:  actor.Verified,
			Active:    actor.Active,
			CreatedAt: actor.CreatedAt,
			UpdatedAt: actor.UpdatedAt,
		}
		res.Data = append(res.Data, item)
	}

	return res, nil
}
func (c Controller) ApprovalRegister(ctx *gin.Context) (Actors, error) {
	actors, err := c.useCase.ApprovalRegister(ctx)
	if err != nil {
		return actors, err
	}

	return actors, nil
}
func (c Controller) RegisActors(ctx *gin.Context) (Actors, error) {
	actors, err := c.useCase.RegisActors(ctx)
	if err != nil {
		return actors, err
	}

	return actors, nil
}
func (c Controller) LoginActors(ctx *gin.Context) (Actors, error) {
	actors, err := c.useCase.LoginActors(ctx)
	if err != nil {
		return actors, err
	}

	return actors, nil
}
