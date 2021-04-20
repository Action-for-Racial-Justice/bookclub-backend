package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/google/uuid"
)

//UserJoinClub adds the user to the club and returns the club member entry id
func (svc *BookClubService) UserJoinClub(joinRequest *models.JoinClubRequest) (string, error) {

	if _, err := svc.mysql.GetUserDataForUserID(joinRequest.UserID); err != nil {
		return "", err
	}

	id := uuid.New()
	joinRequest.EntryID = id
	//TODO validate user struct values exist

	if err := svc.mysql.CreateUserClubMember(joinRequest); err != nil {
		return "", err
	}

	return id.String(), nil
}

//GetClubData returns club data for a club entry ID
func (svc *BookClubService) GetClubData(entryID string) *models.Club {
	//TODO return error here

	clubData, err := svc.mysql.GetClubDataForEntryID(entryID)
	if err != nil {
		log.Printf("Error while retrieving club data from mysql database: %s", err)
		return nil
	}

	return clubData
}

//GetClubs gets a slice of all clubs
func (svc *BookClubService) GetClubs() *models.Clubs {
	//TODO return error here

	clubs, err := svc.mysql.GetListClubs()
	if err != nil {
		log.Printf("Error while retrieving a list of clubs from mysql database: %s", err)
		return nil
	}

	return clubs
}

//GetUserClubs gets a list of clubs that a userID is in
func (svc *BookClubService) GetUserClubs(userID string) (*models.Clubs, error) {

	clubMembers, err := svc.mysql.GetUserClubMembers(userID)
	if err != nil {
		log.Printf("Error while retrieving a list of the users club member entries from mysql database: %s", err)
		return nil, err
	}

	clubs, err := svc.mysql.GetUserClubs(clubMembers)
	if err != nil {
		log.Printf("Error while retrieving a list of users clubs from mysql database: %s", err)
		return nil, err
	}

	return clubs, nil
}

//CreateClub takes in a createRequest, creates a new club,
// adds the leader to the club as a club member
// and returns the club member entry id
func (svc *BookClubService) CreateClub(createRequest *models.CreateClubRequest) (string, error) {
	createRequest.EntryID = uuid.New()

	if err := svc.mysql.CreateClub(createRequest); err != nil {
		log.Printf("Error creating club -> %s", err.Error())
		return "", err
	}

	_, err := svc.UserJoinClub(&models.JoinClubRequest{
		UserID: createRequest.LeaderID,
		ClubID: createRequest.EntryID.String(),
	})

	if err != nil {
		log.Printf("Error joining user to club -> %s", err.Error())
		return "", err
	}

	return createRequest.EntryID.String(), nil
}
