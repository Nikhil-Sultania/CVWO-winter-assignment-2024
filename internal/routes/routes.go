package routes

import (
	"github.com/CVWO/sample-go-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.PUT("/signup", handlers.CreateUser)
	r.POST("/login", handlers.GetUser)
	r.POST("/createPost", handlers.CreatePost)
	r.GET("/thread/:id", handlers.GetPostByID)

	// r.DELETE("/users/:id", users.DeleteUser)
}

// func GetRoutes() func(r chi.Router) {
// 	return func(r chi.Router) {
// 		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
// 			response, _ := users.HandleList(w, req)

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(response)
// 		})
// 	}
// }
