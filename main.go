package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var id uint64

type response struct {
	ID uint64 `json:"id"`
}

var counter uint8
var workerid uint8

func main() {
	router := gin.Default()

	router.Use(getWorkerIdContext())

	router.GET("/id", getID)
	router.Run("0.0.0.0:8080")
}

func getID(c *gin.Context) {
	var res response
	res.ID = generateID(c)
	c.IndentedJSON(http.StatusOK, res)
}

func getWorkerIdContext() gin.HandlerFunc {
	s := os.Getenv("POD_NAME")
	//s := "id-maker-74dbf8c59b-64d8t"
	i := strings.LastIndex(s, "-")
	s = s[i+1 : len(s)]
	fmt.Println(s)
	runes := []rune(s)
	var result []int
	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	workid := strings.Trim(strings.Replace(fmt.Sprint(result), " ", "", -1), "[]")
	wid, err := strconv.Atoi(workid)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return func(c *gin.Context) {
		c.Set("wid", wid)
		c.Next()
	}
}

func generateID(c *gin.Context) uint64 {
	currentTime := time.Now().UnixMilli()
	timeSinceEpoch := currentTime
	id := timeSinceEpoch << 8
	counter++
	id |= int64(counter)
	id = id << 8
	wid2 := c.GetInt("wid")
	id |= int64(wid2)
	return uint64(id)

}
