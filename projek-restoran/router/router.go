package router

import (
	"projek-restoran/resto/handler"
	"projek-restoran/resto/repository"
	"projek-restoran/resto/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handlers struct {
	DB  *gorm.DB
	R   *gin.Engine
}

func (h *Handlers) Routes() {
	RepoCust := repository.NewRepoCust(h.DB)
	UsecaseCust := usecase.NewUsecaseCust(RepoCust)

	RepoKtg := repository.NewRepoKtg(h.DB)
	UsecaseKtg := usecase.NewUsecaseKtg(RepoKtg)

	RepoMenu := repository.NewRepoMenu(h.DB)
	UsecaseMenu := usecase.NewUsecaseMenu(RepoMenu)

	RepoMeja := repository.NewRepoMeja(h.DB)
	UsecaseMeja := usecase.NewUsecaseMeja(RepoMeja)

	RepoReservasi := repository.NewRepoReservasi(h.DB)
	UsecaseReservasi := usecase.NewUsecaseReservasi(RepoReservasi)

	RepoKasir := repository.NewRepoKasir(h.DB)
	UsecaseKasir := usecase.NewUsecaseKasir(RepoKasir)
	
	
	v1 := h.R.Group("api")
	handler.RoutesCust(UsecaseCust, v1)
	handler.RoutesKtg(UsecaseKtg, v1)
	handler.RoutesMenu(UsecaseMenu, v1)
	handler.RoutesMeja(UsecaseMeja, v1)
	handler.RoutesReservasi(UsecaseReservasi, v1)
	handler.RoutesKasir(UsecaseKasir, v1)
}