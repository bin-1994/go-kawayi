package api

import (
	"go-biu/app/api/main/controller"
	"go-biu/app/api/main/repository"
	"go-biu/app/api/main/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, db *gorm.DB) {
	//api := router.Group("/api", cors())
	//{
	api := router.Group("/api")
	api.Use(cors())

	styleRepository := repository.NewStyleRepository(db)
	episodeRepository := repository.NewEpisodeRepository(db)

	bangumiRepository := repository.NewBangumiRepository(db)
	bangumiService := service.NewBangumiService(bangumiRepository, styleRepository, episodeRepository)
	bangumiController := controller.NewBangumiController(bangumiService)

	api.GET("/bangumi/list", bangumiController.FindAll)
	api.GET("/bangumi/media", bangumiController.FindById)
	api.GET("/bangumi/filter", bangumiController.FindFilter)

	searchRepository := repository.NewSearchRepository(db)
	searchService := service.NewSearchService(searchRepository)
	searchController := controller.NewSearchController(searchService)

	api.GET("/search", searchController.FindAll)
	api.GET("/search/suggest", searchController.FetchSugget)
	api.GET("/search/hotword", searchController.FetchHotword)

	rankRepository := repository.NewRankRepository(db)
	rankService := service.NewRankService(rankRepository)
	rankController := controller.NewRankController(rankService)

	api.GET("/rank", rankController.FindAll)
	api.POST("/rank/add", rankController.Create)

	feedbackService := service.NewFeedbackService()
	feedbackController := controller.NewFeedbackController(feedbackService)

	api.POST("/feedback/add", feedbackController.Create)

	userReposition := repository.NewUserRepository(db)
	userService := service.NewUserService(userReposition)
	passportController := controller.NewPassportController(userService)
	userController := controller.NewUserController(userService)

	api.POST("/login", passportController.Login)
	api.POST("/register/email", passportController.Register)
	api.POST("/check/nickname", passportController.CheckName)

	bannerRepository := repository.NewBannerRepository(db)
	bannerService := service.NewBannerService(bannerRepository)
	bannerController := controller.NewBannerController(bannerService)

	api.GET("/banner", bannerController.FindAll)

	roomService := service.NewRoomService()
	roomController := controller.NewRoomController(roomService)

	api.POST("/room/:id", roomController.Online)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService)

	api.GET("/comment", commentController.FindAll)
	api.GET("/comment/reply", commentController.FindReply)

	//api.GET("/timeline", super(), controller.FetchTimeline)
	//api.GET("/timeline/online", controller.FetchOnline)

	//api.GET("/danmaku", controller.FetchDanmaku)

	//api.GET("/review", controller.FetchReview)

	//api.GET("/answer/question", member.FetchQuestion)
	//api.POST("/answer/post", member.PostAnswer)

	// 权限认证
	api.Use(auth())

	api.GET("/user/info", userController.FindById)
	api.GET("/user/state", userController.FindState)

	followRepository := repository.NewFollowRepository(db)
	followService := service.NewFollowService(followRepository)
	followController := controller.NewFollowController(followService)

	api.GET("/follow/list", followController.FindAll)
	api.POST("/follow/add", followController.Follow)
	api.POST("/follow/del", followController.UnFollow)
	api.POST("/follow/edit", followController.Update)

	//api.POST("/bangumi/follow", controller.Follow)
	//api.POST("/bangumi/unfollow", controller.UnFollow)
	//api.POST("/bangumi/progress", controller.Progress)
	//api.POST("/bangumi/rest", controller.RestFollowNum)
	//
	//api.GET("/bangumi/list", member.FetchMyBangumi)
	//api.POST("/bangumi/state", controller.UpdateFollowState)

	//api.GET("/space/review", member.FetchMyReview)
	//
	api.POST("/comment/add", commentController.Create)
	api.POST("/comment/del", commentController.Delete)
	//api.POST("/danmaku/post", controller.AddDanmaku)
	//
	//api.POST("/review/post", controller.PostReview)
	//api.POST("/review/delete", controller.DeleteReview)

	// 个人中心
	//api.GET("/account/setting", member.FetchSetting)
	//api.POST("/account/setting", member.UpdateSetting)
	//api.POST("/account/password", member.UpdatePassword)
	//}

	//video := router.Group("/player", cors())
	//{
	//	video.GET("/v/:value", player.Player)
	//	video.GET("/video/:value", player.Player)
	//	video.GET("/weiyun", player.WeiYun)
	//}

	// hls := router.Group("/hls", cors())
	// {
	// 	hls.GET("/:value", player.Hls)
	// }
}
