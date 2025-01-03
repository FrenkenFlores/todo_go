package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

// Method that will be used to initialize the endpoints
// gin.Engine is the main entry point for gin framework.
// gin is a framework used to develope API.
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			items := lists.Group("/:id/items")
			{
				items.GET("/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
				items.POST("/", h.createItem)
			}
		}
	}
	return router
}
