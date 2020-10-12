package resp

import "github.com/gin-gonic/gin"

func Ok(c *gin.Context){
	c.JSON(200, gin.H{"status": Success, "info": "success"})
}

func OkWithData(c *gin.Context, status int, data interface{}){
	c.JSON(200, gin.H{"status": status, "info": "success","data": data})
}

func Error(c *gin.Context, status int, err error){
	c.JSON(200, gin.H{"status": status, "info": err.Error()})
}
