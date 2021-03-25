package mysql 

import(
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
)

const GET_USER_DATA_QUERY = `SELECT`
func (sql *BookClubMysql) GetUserDataForUserID(userID string) (*models.BookClub, error){

	
}