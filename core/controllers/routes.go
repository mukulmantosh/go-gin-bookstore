package controllers

func (s *Server) endpoints() {
	bookRoute(s)

}

func bookRoute(s *Server) {
	s.gin.POST("/book", s.AddBookToDB)
}
