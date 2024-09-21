package http

import "github.com/gin-gonic/gin"




type ScreenshotController struct {

}


// NewScreenshotController creates a new instance of ScreenshotController
func NewScreenshotController(r *gin.RouterGroup) {
	sc := &ScreenshotController{}
	r.GET("/screenshot", sc.GetScreenshot)
	
}


func (sc *ScreenshotController) GetScreenshot(c *gin.Context) {
	c.JSON(200, gin.H{"message": "GetScreenshot"})
}

