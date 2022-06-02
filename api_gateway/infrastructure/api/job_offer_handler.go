package api

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stojic19/XWS-TIM15/api_gateway/domain"
	"github.com/stojic19/XWS-TIM15/api_gateway/infrastructure/services"
	"github.com/stojic19/XWS-TIM15/common/proto/job_offers"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"net/http"
	"time"
)

type JobOffersHandler struct {
	jobOffersClientAddress string
	usersClientAddress     string
}

func NewJobOffersHandler(jobOffersClientAddress, usersClientAddress string) Handler {
	return &JobOffersHandler{
		jobOffersClientAddress: jobOffersClientAddress,
		usersClientAddress:     usersClientAddress,
	}
}

func (handler *JobOffersHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/job_offers/details", handler.GetJobOffersDetails)
	if err != nil {
		panic(err)
	}
}

func (handler *JobOffersHandler) GetJobOffersDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jobOffersInfo, err := initializeJobOffers(w, pathParams)
	if err != nil {
		return
	}

	err = handler.addJobOffersInfo(jobOffersInfo)
	if err != nil {
		return
	}

	for _, jobOfferInfo := range jobOffersInfo.JobOffers {
		for _, followerInfo := range jobOfferInfo.Subscribers {
			err = handler.addUserInfo(followerInfo)
			if err != nil {
				break
			}
		}
	}

	finishJobOffers(w, err, jobOffersInfo)
}

func initializeJobOffers(w http.ResponseWriter, pathParams map[string]string) (*domain.JobOffersUsersInfoList, error) {
	jobOffersInfo := &domain.JobOffersUsersInfoList{}
	jobOffersInfo.JobOffers = []*domain.JobOfferUsersInfo{}
	return jobOffersInfo, nil
}

func (handler *JobOffersHandler) addJobOffersInfo(jobOffersInfoList *domain.JobOffersUsersInfoList) error {
	jobOffersClient := services.NewJobOffersClient(handler.jobOffersClientAddress)
	jobOffers, err := jobOffersClient.GetAll(context.TODO(), &job_offers.GetAllRequest{})
	if err != nil {
		return err
	}
	for _, jobOfferPb := range jobOffers.JobOffers {
		jobOffer := domain.JobOfferUsersInfo{
			Id:           jobOfferPb.Id,
			Position:     jobOfferPb.Position,
			Description:  jobOfferPb.Description,
			Requirements: jobOfferPb.Requirements,
			IsActive:     jobOfferPb.IsActive,
			Subscribers:  []*domain.UserJobOfferInfo{},
		}
		for _, follower := range jobOfferPb.Subscribers {
			follower := domain.UserJobOfferInfo{
				Id: follower.Id,
			}
			jobOffer.Subscribers = append(jobOffer.Subscribers, &follower)
		}
		jobOffersInfoList.JobOffers = append(jobOffersInfoList.JobOffers, &jobOffer)
	}
	return nil
}

func (handler *JobOffersHandler) addUserInfo(jobOffersInfo *domain.UserJobOfferInfo) error {
	usersClient := services.NewUsersClient(handler.usersClientAddress)
	user, err := usersClient.GetUser(context.TODO(), &users.GetUserRequest{Id: jobOffersInfo.Id})
	if err != nil {
		return err
	}
	jobOffersInfo.Gender = user.User.Gender
	jobOffersInfo.Name = user.User.Name
	jobOffersInfo.Username = user.User.Username
	jobOffersInfo.DateOfBirth, _ = time.Parse("MM/DD/YYYY", user.User.DateOfBirth)
	return nil
}

func finishJobOffers(w http.ResponseWriter, err error, jobOffersInfo *domain.JobOffersUsersInfoList) {
	response, err := json.Marshal(jobOffersInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
