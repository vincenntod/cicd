package actors

import "github.com/gin-gonic/gin"

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
func (u UseCase) ShowActors() ([]Actors, error) {
	return u.repo.ShowActors()
}
func (u UseCase) ShowActorsByName(c *gin.Context) ([]Actors, error) {
	return u.repo.ShowActorsByName(c)
}
func (u UseCase) EditActors(c *gin.Context) (Actors, error) {
	return u.repo.EditActors(c)
}
func (u UseCase) DeleteActors(c *gin.Context) (Actors, error) {
	return u.repo.DeleteActors(c)
}
func (u UseCase) ShowRegister() ([]Actors, error) {
	return u.repo.ShowRegister()
}
func (u UseCase) ApprovalRegister(c *gin.Context) (Actors, error) {
	return u.repo.ApprovalRegister(c)
}
func (u UseCase) RegisActors(c *gin.Context) (Actors, error) {
	return u.repo.RegisActors(c)
}
func (u UseCase) LoginActors(c *gin.Context) (Actors, error) {
	return u.repo.LoginActors(c)
}
