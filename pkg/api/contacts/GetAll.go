package contacts

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/mmanjoura/niya-voyage/backend/pkg/common"
	"github.com/mmanjoura/niya-voyage/backend/pkg/database"
	"github.com/mmanjoura/niya-voyage/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// GetAll 			godoc
// @Summary 		Get all Contacts
// @Description 	Get all Contacts
// @Tags 			Contacts
// @Accept  		json
// @Produce 		json
// @Param 			offset query int false "Offset"
// @Param 			limit query int false "Limit"
// @Success 		200 {object} []Contact
// @Failure 		400 {string} string "Bad Request"
// @Failure 		500 {string} string "Internal Server Error"
// @Router 			/contacts [get]
func GetAll(c *gin.Context) {
	db := database.Database.DB
	offset, limit := common.GetPaginationParams(c)
	contacts, err := RetrieveContacts(c, db, limit, offset)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while retreiving Contacts": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

func RetrieveContacts(c *gin.Context, db *sql.DB, limit, offset int) ([]models.Contact, error) {

	// Build WHERE clause. Each part of the clause is AND-ed together to further
	// restrict the results. Placeholders are added to "args" and are used
	// to avoid SQL injection.
	// Each filter field is optional.
	_, args := []string{"1 = 1"}, []interface{}{}

	contacts := []models.Contact{}

	rows, err := db.QueryContext(c, `SELECT ID,
			name,
			email,
			subject,
			Created_At,
			Updated_At,
			COUNT(*) OVER()
		FROM Contacts `,
		args...,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		contact, err := scanContact(rows)

		if err != nil {
			return nil, err
		}

		if contact.ID != 0 {

			contacts = append(contacts, contact)
		}
	}
	defer rows.Close()

	return contacts, nil
}

func scanContact(rows *sql.Rows) (models.Contact, error) {
	contact := models.Contact{}
	var n int

	err := rows.Scan(&contact.ID,
		&contact.Name,
		&contact.Email,
		&contact.Subject,
		(*time.Time)(&contact.CreatedAt),
		(*time.Time)(&contact.UpdatedAt),
		&n,
	)

	return contact, err
}
