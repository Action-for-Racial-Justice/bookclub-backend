package service

import (
	"log"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/bcerrors"
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

	if err := svc.mysql.CreateUserClubMember(joinRequest); err != nil {
		return "", err
	}

	return id.String(), nil
}

//UserLeaveClub checks if the user is the leader of the club or a member, and deletes or leaves the club respectively
func (svc *BookClubService) UserLeaveClub(leaveRequest *models.LeaveClubRequest) error {

	leader, err := svc.mysql.IsUserClubLeader(leaveRequest)
	if err != nil {
		return err
	}

	if leader {
		if err := svc.mysql.DeleteClub(leaveRequest); err != nil {
			return err
		}
	} else {
		if err := svc.mysql.DeleteUserClubMember(leaveRequest); err != nil {
			return err
		}
	}

	return nil
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
		svc.logger.Errorw(err.Error())
		return "", bcerrors.NewError("MYSQL failed creating club", bcerrors.InternalError).
			WithExternalMessage("could not creat club").
			WithRootCause(err)
	}

	_, err := svc.mysql.UserJoinClub(&models.JoinClubRequest{
		UserID: createRequest.LeaderID,
		ClubID: createRequest.EntryID.String(),
	})

	if err != nil {
		log.Printf("Error joining user to club -> %s", err.Error())
		return "", err
	}

	return createRequest.EntryID.String(), nil
}
