package main

import (
	"minder/src/db"
	"minder/src/seed"
	"minder/src/server"
	"minder/src/server/controller"
	"minder/src/server/middleware"
	"minder/src/server/repository/repo_impl"
	"minder/src/server/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := db.ConnectGormDB()
	if err != nil {
		panic(err)
	}

	seed.SeedLocation(db)
	seed.SeedInterest(db)
	seed.SeedUser(db, 500)
	seed.SeedUserInterest(db)
	seed.SeedUserMembership(db)
	seed.SeedUserPhoto(db)
	seed.SeedMembership(db)
	seed.SeedMembershipPrivilege(db)
	seed.SeedPrivilege(db)

	userRepo := repo_impl.NewUserRepo(db)
	membershipRepo := repo_impl.NewMembershipRepo(db)
	userMembershipRepo := repo_impl.NewUserMembershipRepo(db)
	userSwipeRepo := repo_impl.NewUserSwipeRepo(db)
	locationRepo := repo_impl.NewLocationRepo(db)
	privilegeRepo := repo_impl.NewPrivilegeRepo(db)
	membershipPrivilegeRepo := repo_impl.NewMembershipPrivilegeRepo(db)
	purchaseRepo := repo_impl.NewPurchaseRepo(db)
	interestRepo := repo_impl.NewInterestRepo(db)

	authService := service.NewAuthServices(userRepo, locationRepo, membershipRepo, userMembershipRepo)
	userService := service.NewUserServices(userRepo, locationRepo, interestRepo, userSwipeRepo, userMembershipRepo)
	locationService := service.NewLocationServices(locationRepo)
	membershipService := service.NewMembershipServices(membershipRepo, userRepo, privilegeRepo, membershipPrivilegeRepo)
	mediaService := service.NewMediaServices()
	userSwipeService := service.NewUserSwipeServices(userSwipeRepo, userRepo, userMembershipRepo, membershipPrivilegeRepo, privilegeRepo)
	purchaseService := service.NewPurchaseServices(purchaseRepo, membershipRepo, userMembershipRepo, userRepo)
	privilegeService := service.NewPrivilegeServices(privilegeRepo, membershipPrivilegeRepo)

	authHandler := controller.NewAuthHandler(authService)
	userHandler := controller.NewUserHandler(userService)
	membershipHandler := controller.NewMembershipHandler(membershipService)
	locationHandler := controller.NewLocationHandler(locationService)
	mediaHandler := controller.NewMediaHandler(mediaService)
	swipeHandler := controller.NewSwipeHandler(userSwipeService)
	purchaseHandler := controller.NewPurchaseHandler(purchaseService)
	privilegeHandler := controller.NewPrivilegeHandler(privilegeService)

	router := gin.Default()
	router.Use(gin.Logger())

	middleware := middleware.NewMiddleware(userService)

	app := server.NewRouterGin(
		router,
		authHandler,
		userHandler,
		membershipHandler,
		privilegeHandler,
		locationHandler,
		mediaHandler,
		swipeHandler,
		purchaseHandler,
		middleware,
	)

	app.Start(":4444")
}
