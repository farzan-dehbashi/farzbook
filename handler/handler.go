package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FarzbookHandler struct {
        counter int
}

func NewFarzbookHandler() *FarzbookHandler {
        return &FarzbookHandler{}
}

func (h *FarzbookHandler) Home(c *gin.Context) {
        c.String(http.StatusOK, "home")
        fmt.Println("home")
}

func (h *FarzbookHandler) Login(c *gin.Context) {
        c.String(http.StatusOK, "login")
        fmt.Println("login")
}

func (h *FarzbookHandler) Profile(c *gin.Context) {
        c.String(http.StatusOK, "profile")
        fmt.Println("profile")
}

func (h *FarzbookHandler) Logout(c *gin.Context) {
        c.String(http.StatusOK, "logout")
        fmt.Println("logout")
}

func (h *FarzbookHandler) Add(c *gin.Context) {
        h.counter++
        counterStr := strconv.Itoa(h.counter)
        c.String(http.StatusOK, "add, counter: %s", counterStr)
        fmt.Println("add", h.counter)
}

func (h *FarzbookHandler) Sub(c *gin.Context) {
        h.counter--
        counterStr := strconv.Itoa(h.counter)
        c.String(http.StatusOK, "sub, counter: %s", counterStr)
        fmt.Println("sub", h.counter)
}
