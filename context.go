package main

//this context contains all the information about the request that the handler might need to process it
//it easy to pass request-scoped values , cancelation signals,and deadlines across API boundries to all the goroutines involved in handling a request
//:9090/api/path?n=vanshika&email=somevalue
//after ? all are parameters & adds parameters
//like header cookies etc
//json {},[{},{}]
//XML <tag><tag>

//usage of context
//err:=c.Bind(&obj)//json or xml body
//err:=c.BindQuery(&obj)
//err:=c.BindXML(&obj)
//err:=c.BindYAML(&obj)
//err:=c.BindHeader(&obj)
//err:=c.BindJSON(&obj)
//c.Header("user-id","72348723")
//c.SetCookie(name,value string,maxAge int,path,domain,string,secure,httpOnly bool)
//val,err:=c.Cookie("name")
//err:=c.SaveUploadedFile(file,dest)
//form:=c.MultipartForm()
//key:=c.Postform("key","default value")
//id :=c.Query("id")
//id :=c.Param("id")
//name:=c.DefaultQuery("name","jack")
//c.GetFloat64("mykey")
//c.Set("mykey","my value")
//c.Get("mykey")
//c.MustGet("mykey")
//c.GetString("mykey")
//c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"message":"something went wrong"})
//c.IsAborted()
//c.AbortWithStatus(http.StatusBadRequest)
//c.Abort()

//c.HTML(
//	http.StatusOK,"index.html",gin.H{
//"message":"this is a message"
//	})

//c.JSON(
//	http.StatusOK,"index.html",gin.H{
//"message":"this is a message"
//	})

//c.XML(
//	http.StatusOK,"index.html",gin.H{
//"message":"this is a message"
//	})

//c.YAML(
//	http.StatusOK,"index.html",gin.H{
//"message":"this is a message"
//	})
