package server

import (
	"log"
	"minder/src/server/controller"
	"minder/src/server/middleware"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	router     *gin.Engine
	auth       *controller.AuthHandler
	user       *controller.UserHandler
	membership *controller.MembershipHandler
	location   *controller.LocationHandler
	media      *controller.MediaHandler
	swipe      *controller.SwipeHandler
	purchase   *controller.PurchaseHandler
	middleware *middleware.Middleware
}

func NewRouterGin(
	router *gin.Engine,
	auth *controller.AuthHandler,
	user *controller.UserHandler,
	membership *controller.MembershipHandler,
	location *controller.LocationHandler,
	media *controller.MediaHandler,
	swipe *controller.SwipeHandler,
	purchase *controller.PurchaseHandler,
	middleware *middleware.Middleware,
) *GinRouter {
	return &GinRouter{
		router:     router,
		auth:       auth,
		user:       user,
		membership: membership,
		location:   location,
		media:      media,
		swipe:      swipe,
		purchase:   purchase,
		middleware: middleware,
	}
}

func (r *GinRouter) Start(port string) {
	r.router.Use(r.middleware.Trace)

	auth := r.router.Group("/auth")
	auth.POST("/register", r.auth.Register)
	auth.POST("/login", r.auth.Login)

	user := r.router.Group("/users")
	user.GET("/me", r.middleware.Auth, r.user.Profile)
	user.GET("/explore", r.middleware.Auth, r.user.Explore)
	user.PUT("/me", r.middleware.Auth, r.user.Update)
	user.DELETE("/me", r.middleware.Auth, r.user.DeleteUser)
	user.GET("/count", r.middleware.Auth, r.middleware.AdminOnly(r.user.CountUsers))

	membership := r.router.Group("/memberships")
	membership.GET("/", r.middleware.Auth, r.membership.GetMemberships)
	membership.POST("/", r.middleware.Auth, r.middleware.AdminOnly(r.membership.CreateMembership))
	membership.GET("/:membershipId", r.middleware.Auth, r.membership.GetMembershipById)
	membership.PUT("/:membershipId", r.middleware.Auth, r.middleware.AdminOnly(r.membership.UpdateMembership))
	membership.DELETE("/:membershipId", r.middleware.Auth, r.middleware.AdminOnly(r.membership.DeleteMembership))

	location := r.router.Group("/locations")
	location.GET("/", r.middleware.Auth, r.location.GetLocations)
	location.POST("/", r.location.CreateLocation)
	location.GET("/:locationId", r.middleware.Auth, r.location.GetLocationById)
	location.PUT("/:locationId", r.middleware.Auth, r.middleware.AdminOnly(r.location.UpdateLocation))
	location.DELETE("/:locationId", r.middleware.Auth, r.middleware.AdminOnly(r.location.DeleteLocation))

	media := r.router.Group("/media")
	media.POST("", r.middleware.Auth, r.media.UploadMedia)
	media.GET("/:mediaId", r.middleware.Auth, r.media.GetMedia)
	media.DELETE("/:mediaId", r.middleware.Auth, r.media.DeleteMedia)

	swipe := r.router.Group("/swipes")
	swipe.POST("/like/:targetId", r.middleware.Auth, r.swipe.Like)
	swipe.POST("/pass/:targetId", r.middleware.Auth, r.swipe.Pass)
	swipe.POST("/favourite/:targetId", r.middleware.Auth, r.swipe.Favourite)

	purchase := r.router.Group("/purchases")
	purchase.POST("/", r.middleware.Auth, r.purchase.CreatePurchase)
	purchase.GET("/:purchaseId", r.middleware.Auth, r.purchase.GetPurchase)
	purchase.POST("/cancel/:purchaseId", r.middleware.Auth, r.purchase.CancelPurchase)
	purchase.POST("/sync", r.middleware.Auth, r.middleware.AdminOnly(r.purchase.SyncPurchase))

	err := r.router.Run(port)
	if err != nil {
		log.Fatal(err.Error())
		panic("error while running apps")
	}
}
