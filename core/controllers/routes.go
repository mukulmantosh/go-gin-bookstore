package controllers

func (s *Server) endpoints() {
	bookRoute(s)

}

func bookRoute(s *Server) {
	s.gin.POST("/book", s.CreateBook)
	s.gin.GET("/book", s.ListBook)
	s.gin.POST("/book/upload", s.UploadBookCover)
}
