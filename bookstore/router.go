package bookstore

func Router() {
	router.POST("/api/book", CreateBook)
	router.GET("/api/book/:id", GetBook)
	router.PUT("/api/book/:id", UpdateBook)
	router.DELETE("/api/book/:id", DeleteBook)
	router.GET("/api/book", GetAllBook)
}
