package middleware

//func Cors(c *gin.Context)  {
//	origin := c.Request.Header.Get("Origin")
//	if origin != "" {
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
//		c.Header("Access-Control-Allow-Credentials", "true")
//		c.Set("content-type", "application/json")
//	}
//	//放行所有OPTIONS方法if method == "OPTIONS" {
//	c.AbortWithStatus(http.StatusNoContent)
//
//    c.Next()
//}
