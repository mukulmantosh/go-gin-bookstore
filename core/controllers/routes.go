package controllers

func (s *Server) endpoints() {
	bookRoute(s)
	authorRoute(s)

}

func bookRoute(s *Server) {
	s.gin.POST("/book", s.CreateBook)
	s.gin.GET("/book", s.ListBook)
	s.gin.PUT("/book/:id", s.UpdateBook)
	s.gin.DELETE("/book/:id", s.DeleteBook)
	s.gin.POST("/book/cover/:id", s.UploadBookCover)
}

func authorRoute(s *Server) {
	s.gin.GET("/authors", s.ListAuthors)
	s.gin.POST("/author", s.CreateAuthor)
	s.gin.POST("/author/book/link", s.LinkBook)
}
