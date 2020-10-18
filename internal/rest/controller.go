package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/bhakiyakalimuthu/blogger/internal/svc"
)

type RegisterRoute interface {
	Register(router chi.Router)
}
type Controller struct {
	logger  *zap.Logger
	service *svc.Service
}

func NewController(logger *zap.Logger, service *svc.Service) *Controller {
	return &Controller{
		logger:  logger,
		service: service,
	}
}

func (c *Controller) SetupRouter(router chi.Router) chi.Router {
	router.Route("/user", func(r chi.Router) {
		r.Get("/{id}", c.Get)
		r.Post("/create", c.Create)
		r.Patch("/update", c.Update)
		r.Delete("/delete", c.Delete)
	})
	return router
}

func (c *Controller) Get(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	user, err := c.service.GetUser(request.Context(), id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	resp, _ := json.Marshal(&user)

	response.WriteHeader(http.StatusOK)
	response.Write(resp)
	return

}
func (c *Controller) Create(response http.ResponseWriter, request *http.Request) {
	var p svc.Payload
	if err := json.NewDecoder(request.Body).Decode(&p); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`unknown payload content`))
		return
	}

	if err := c.service.CreateUser(request.Context(), p); err != nil {
		c.logger.Info("user created")
		response.WriteHeader(http.StatusCreated)
		return
	}
}
func (c *Controller) Update(response http.ResponseWriter, request *http.Request) {

}
func (c *Controller) Delete(response http.ResponseWriter, request *http.Request) {

}
