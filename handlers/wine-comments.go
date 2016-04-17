package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dfreire/sunny/crud"
	"github.com/dfreire/sunny/middleware"
	"github.com/dfreire/sunny/model"
	"github.com/labstack/echo"
	"gopkg.in/Masterminds/squirrel.v1"
)

// http http://localhost:3500/customer-wine-comments?customerId="customer-1"
func GetCustomerWineComments(c echo.Context) error {
	log.Println("GetCustomerWineComments")

	customerId := c.QueryParam("customerId")

	db := c.Get(middleware.DB).(*sql.DB)

	comments := []model.WineComment{}

	rows, err := squirrel.
		Select("id", "wineId", "wineYear", "comment").
		From("WineComment").
		Where(squirrel.Eq{
			"customerId": customerId,
		}).
		RunWith(db).Query()
	if err != nil {
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	}

	for rows.Next() {
		var c model.WineComment
		rows.Scan(&c.Id, &c.WineId, &c.WineYear, &c.Comment)
		comments = append(comments, c)
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true, Data: comments})
}

// http POST http://localhost:3500/wine-comment wineId="wine-1" wineYear=2014 customerId="customer-1" comment="great"
// http POST http://localhost:3500/wine-comment wineId="wine-1" wineYear=2014 customerId="customer-2" comment="fantastic" id="5713f0bf5a1d1801bb000001"
func UpsertWineComment(c echo.Context) error {
	log.Println("UpsertWineComment")

	db := c.Get(middleware.DB).(*sql.DB)

	var comment crud.Record
	c.Bind(&comment)

	var err error

	if comment["id"] == nil || comment["id"] == "" {
		createdAt := time.Now().Format(time.RFC3339)
		comment["createdAt"] = createdAt
		comment["updatedAt"] = createdAt
		err = crud.Create(db, "WineComment", comment)

	} else {
		id := comment["id"].(string)
		updatedAt := time.Now().Format(time.RFC3339)
		comment["updatedAt"] = updatedAt
		delete(comment, "customerId")
		delete(comment, "wineId")
		delete(comment, "wineYear")
		err = crud.Update(db, "WineComment", comment, []string{id})
	}

	if err != nil {
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	}

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}

// http http://localhost:3500/signup-customer-with-wine-comment email="dario.freire@gmail.com" role="wine_lover" wineComments:='[{"wineId": "wine-1", "wineYear": 2015, "comment": "great"}]'
func SignupCustomerWithWineComment(c echo.Context) error {
	log.Println("RegisterCustomerWithWineComment")

	var reqData struct {
		Email        string `json:"email"`
		Role         string `json:"role"`
		WineComments []struct {
			WineId   string `json:"wineId"`
			WineYear int    `json:"wineYear"`
			Comment  string `json:"comment"`
		} `json:"wineComments"`
	}

	c.Bind(&reqData)

	db := c.Get(middleware.DB).(*sql.DB)

	id, err := getOrCreateCustomerId(db, reqData.Email, reqData.Role)
	if err != nil {
		log.Printf("error: %+v", err)
		return c.JSON(http.StatusInternalServerError, JsonResponse{Ok: false})
	}

	log.Printf("customer id: %+v", id)

	return c.JSON(http.StatusOK, JsonResponse{Ok: true})
}

func getOrCreateCustomerId(db *sql.DB, email string, role string) (string, error) {
	log.Println("getOrCreateCustomer")

	rows, err := squirrel.
		Select("id").
		From("Customer").
		Where(squirrel.Eq{"email": email}).
		RunWith(db).Query()
	if err != nil {
		return "", err
	}

	var id string
	for rows.Next() {
		rows.Scan(&id)
	}

	if id == "" {
		customer := crud.Record{
			"email":     email,
			"role":      role,
			"createdAt": time.Now().Format(time.RFC3339),
		}
		err = crud.Create(db, "Customer", customer)
		if err != nil {
			return "", err
		}
		id = customer["id"].(string)
	}

	return id, nil
}
