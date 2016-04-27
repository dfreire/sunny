package handlers

import "github.com/labstack/echo"

func GetDoc(c echo.Context) error {
	return c.File("doc.txt")
}
