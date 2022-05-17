// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type student struct {
// 	Name string
// 	Age  int8
// }

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()

// 	// r.LoadHTMLGlob("templates/*")

// 	// stu1 := &student{Name: "Geektutu", Age: 18}
// 	// stu2 := &student{Name: "Jack", Age: 20}

// 	// r.GET("/arr", func(ctx *gin.Context) {
// 	// 	ctx.HTML(http.StatusOK, "arr.tmpl", gin.H{
// 	// 		"title":  "Gin",
// 	// 		"stuArr": [2]*student{stu1, stu2},
// 	// 	})
// 	// })

// 	// GET 和 POST 混合
// 	// r.POST("/posts", func(c *gin.Context) {
// 	// 	id := c.Query("id")
// 	// 	page := c.DefaultQuery("page", "0")
// 	// 	username := c.PostForm("username")
// 	// 	password := c.DefaultPostForm("username", "000000") // 可设置默认值

// 	// 	c.JSON(http.StatusOK, gin.H{
// 	// 		"id":       id,
// 	// 		"page":     page,
// 	// 		"username": username,
// 	// 		"password": password,
// 	// 	})
// 	// })

// 	r.POST("/post", func(ctx *gin.Context) {
// 		ids := ctx.Query("ids")
// 		names := ctx.PostFormMap("names")

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"ids":   ids,
// 			"names": names,
// 		})
// 	})
// 	r.Run(":8000")
// }

// package main

// import (
// 	"fmt"
// )

// type Mutatable struct {
// 	a int
// 	b int
// }

// func (m Mutatable) StayTheSame() {
// 	m.a = 5
// 	m.b = 7
// }

// func (m *Mutatable) Change() {
// 	m.a = 5
// 	m.b = 7
// }

// func main() {
// 	m := &Mutatable{0, 0}
// 	fmt.Println("初始值为:", "a=", m.a, "b=", m.b)

// 	m.StayTheSame()
// 	fmt.Println("StayTheSame后:", "a=", m.a, "b=", m.b)
// 	m.Change()
// 	fmt.Println("Change后:", "a=", m.a, "b=", m.b)
// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)
	})
	router.GET("/albums/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Loop through the list of albums, looking for
		// an album whose ID value matches the parameter.
		for _, a := range albums {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	})

	router.POST("/savealbums", func(c *gin.Context) {
		var newAlbum album
		// newAlbum.ID = c.PostForm("id")
		// newAlbum.Title = c.PostForm("title")
		// newAlbum.Artist = c.PostForm("artist")
		// newAlbum.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newAlbum); err != nil {
			fmt.Println(err)
			return
		}

		// Add the new album to the slice.
		albums = append(albums, newAlbum)
		//c.IndentedJSON(http.StatusCreated, newAlbum)
		c.IndentedJSON(http.StatusOK, albums)
	})
	router.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	router.Run(":8000")
}

/* win10下
curl   -H "Accept:application/json" -H  "Content-Type:application/json" -X POST -d "{\"id\":\"4\",\"title\":\"123\",\"artist\":\"1234\",\"price\":4.1}"  http://127.0.0.1:8000/savealbums
*/
