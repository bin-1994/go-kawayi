package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(router *gin.Engine, db *gorm.DB) {
	admin := router.Group("/x", Logger())

	admin.GET("/bangumi")

	//hlsRepository := repository.NewHlsRepository(db)
	//hlsService := service.NewHlsService(hlsRepository)
	//hlsController := system.NewHlsController(hlsService)
	//
	//admin := router.Group("/x", auth())
	//{
	//	//admin.GET("/bangumi/:media_id", bangumi.NewBangumiController)
	//	// 以下是旧的
	//	admin.GET("/media", bangumi.FetchAllMedia)
	//	admin.GET("/media/:media_id", bangumi.FetchMedia)
	//	admin.POST("/media/add", bangumi.AddMedia)
	//	admin.POST("/media/edit", bangumi.EditMedia)
	//	admin.POST("/media/update", new(bangumi.MediaController).UpdateTime)
	//	admin.POST("/media/cache", new(bangumi.MediaController).DeleteCache)
	//
	//	admin.GET("/comment", new(bangumi.CommentController).List)
	//	admin.POST("/comment/del", new(bangumi.CommentController).Delete)
	//
	//	admin.GET("/danmaku", new(bangumi.DanmakuController).FetchDanmaku)
	//	admin.POST("/danmaku/del", new(bangumi.DanmakuController).DelDanmaku)
	//
	//	admin.GET("/answer", bangumi.FetchAllAnswer)
	//	admin.GET("/answer/:qid", bangumi.FetchAnswer)
	//	admin.POST("/answer/add", bangumi.AddAnswer)
	//	admin.POST("/answer/edit", bangumi.EditAnswer)
	//
	//	admin.GET("/vip", user.FetchAllVip)
	//	admin.GET("/vip/:vid", user.FetchVip)
	//	admin.GET("/vip/:vid/check", user.CheckVip)
	//	admin.GET("/vip/:vid/notpass", user.NotPassVip)
	//	admin.POST("/vip/edit", user.EditVip)
	//
	//	admin.GET("/user", user.FetchAllUser)
	//	admin.GET("/user/:user_id", user.FetchUser)
	//	admin.POST("/user/edit", user.EditUser)
	//
	//	admin.GET("/season", bangumi.FetchAllSeason)
	//	admin.GET("/season/:season_id", bangumi.FetchSeason)
	//	admin.POST("/season/add", bangumi.AddSeason)
	//	admin.POST("/season/edit", bangumi.EditSeason)
	//
	//	admin.GET("/banner", bangumi.FetchAllBanner)
	//	admin.GET("/banner/:id", bangumi.FetchBanner)
	//	admin.POST("/banner/add", bangumi.AddBanner)
	//	admin.POST("/banner/edit", bangumi.EditBanner)
	//
	//	admin.GET("/style", new(bangumi.StyleController).FetchAllStyle)
	//	admin.POST("/style/add", new(bangumi.StyleController).Create)
	//	admin.POST("/style/edit", new(bangumi.StyleController).Update)
	//
	//	admin.POST("/episode/add", new(bangumi.EpisodeController).Create)
	//	admin.POST("/episode/edit", new(bangumi.EpisodeController).Update)
	//	admin.GET("/media/:media_id/episode", new(bangumi.EpisodeController).Get)
	//	//admin.GET("/maxep/:media_id", bangumi.FetchMaxEp)
	//	//admin.GET("/media/:media_id/ep", bangumi.FetchMediaToEp)
	//	//admin.POST("/ep/add", bangumi.AddEp)
	//	//admin.POST("/ep/edit", bangumi.EditEp)
	//	admin.GET("/eps", bangumi.FetchAllEp)
	//	admin.POST("/eps/add", bangumi.AddEps)
	//	admin.POST("/eps/replace", bangumi.ReplaceEp)
	//
	//	admin.GET("/search/list", bangumi.FetchSearchList)
	//	admin.POST("/search/del", bangumi.DeleteKeyword)
	//	admin.POST("/search/delall", bangumi.DelAllKeyword)
	//
	//	admin.GET("/feedback/list", new(bangumi.FeedbackController).FeedbackList)
	//	admin.POST("/feedback/del", new(bangumi.FeedbackController).DelFeedback)
	//	admin.POST("/feedback/delall", new(bangumi.FeedbackController).DelAllFeedback)
	//
	//	admin.GET("/timeline", bangumi.FetchTimeline)
	//
	//	admin.GET("/system", system.GetSystemInfo)
	//	admin.POST("/upload/image", system.UploadImage)
	//	admin.GET("/upload/list", system.FetchAllUpload)
	//	admin.POST("/upload/del", system.DeleteUpload)
	//	admin.POST("/upload/webp", system.CreateWebp)
	//
	//	admin.POST("/upload/hls", hlsController.Upload)
	//	admin.GET("/hls/list", hlsController.List)
	//	admin.GET("/hls/file", hlsController.Get)
	//	admin.POST("/hls/update", hlsController.Update)
	//
	//}
	//router.POST("/x/login", user.Login)
	//router.GET("/x/upload/move", system.MoveImg)
}
