package controllers

func (s *Server) endpoints() {
	bookRoute(s)
	authorRoute(s)
	customerRoute(s)
	reviewRoute(s)
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
func customerRoute(s *Server) {
	s.gin.POST("/customer", s.CreateCustomer)
	s.gin.PUT("/customer/:id", s.UpdateCustomer)
	s.gin.DELETE("/customer/:id", s.DeleteCustomer)
}
func reviewRoute(s *Server) {
	s.gin.POST("/review", s.CreateReview)
	s.gin.GET("/review/book/:id", s.ListReviewByBook)
}
