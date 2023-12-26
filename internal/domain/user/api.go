package user

import (
	"github.com/google/uuid"
	"github.com/happsie/go-webserver-template/internal/architecture"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Api struct {
	Container  *architecture.Container
	Repository Repository
}

func (a Api) Register(e *echo.Echo) {
	g := e.Group("/api/users/user-v1")
	g.POST("", a.create)
	g.GET("{id}", a.read)
	g.PUT("", a.update)
	g.DELETE("{id}", a.delete)
}

func (a Api) create(c echo.Context) error {
	req := CreateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	creationTime := time.Now()
	user := User{
		ID:          uuid.New(),
		DisplayName: req.DisplayName,
		CreatedAt:   creationTime,
		UpdatedAt:   creationTime,
		Version:     1,
	}
	err = a.Repository.Create(user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusCreated, user)
}

func (a Api) read(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user, err := a.Repository.Read(ID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, user)
}

func (a Api) update(c echo.Context) error {
	req := UpdateUserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user, err := a.Repository.Read(req.ID)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	user.DisplayName = req.DisplayName
	err = a.Repository.Update(user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, user)
}

func (a Api) delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	err = a.Repository.Delete(ID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusNoContent)
}
