package transport

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"hrm/pkg/candidate/domain"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	candidateService    domain.CandidateService
	candidateRepository domain.CandidateRepository
}

func NewHandler(s domain.CandidateService) *Handler {
	return &Handler{
		candidateService: s,
	}
}

func NewRouter(s *Handler) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	sr := router.PathPrefix("/api/v1").Subrouter()

	sr.Name("Register Candidate").
		Methods(http.MethodPost).
		Path("/candidate/register").
		HandlerFunc(createHandlerFunc(s.RegisterCandidate))

	sr.Name("Make Offer").
		Methods(http.MethodPut).
		Path("/candidate/{candidateId}/offer").
		HandlerFunc(createHandlerFunc(s.MakeOffer))

	sr.Name("Decline Candidate").
		Methods(http.MethodPost).
		Path("/candidate/{candidateId}/decline").
		HandlerFunc(createHandlerFunc(s.DeclineCandidate))

	sr.Name("Hire candidate").
		Methods(http.MethodPost).
		Path("/candidate/{candidateId}/hire").
		HandlerFunc(createHandlerFunc(s.HireCandidate))

	sr.Name("Get candidate").
		Methods(http.MethodGet).
		Path("/candidate/{candidateId}").
		HandlerFunc(createHandlerFunc(s.GetCandidate))

	sr.Name("Get candidate list").
		Methods(http.MethodGet).
		Path("/candidate/list").
		HandlerFunc(createHandlerFunc(s.GetCandidateList))

	return router
}

func (s *Handler) RegisterCandidate(w http.ResponseWriter, r *http.Request) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	var createCandidateRequestData createCandidateRequest
	err = json.Unmarshal(b, &createCandidateRequestData)
	if err != nil {
		return err
	}

	registerOptions := []domain.CandidateOption{
		domain.WithName(createCandidateRequestData.Name),
		domain.WithEmail(createCandidateRequestData.Email),
		domain.WithPhone(createCandidateRequestData.Phone),
		domain.WithAddress(createCandidateRequestData.Address),
	}

	c, err := s.candidateService.Register(registerOptions...)
	if err != nil {
		return err
	}

	err = jsonResponse(w, candidateResponse{
		Id:      c.Id,
		Name:    c.Name,
		Email:   c.Email,
		Phone:   c.Phone,
		Address: c.Address,
		Status:  int(c.Status.Type),
	})

	return err
}

func (s *Handler) MakeOffer(w http.ResponseWriter, r *http.Request) error {
	candidateId, ok := mux.Vars(r)["candidateId"]
	if !ok {
		badRequestResponse(w, "OrderId not found")
		return nil
	}

	return s.candidateService.MakeOffer(candidateId)
}

func (s *Handler) HireCandidate(w http.ResponseWriter, r *http.Request) error {
	candidateId, ok := mux.Vars(r)["candidateId"]
	if !ok {
		badRequestResponse(w, "OrderId not found")
		return nil
	}

	return s.candidateService.Hire(candidateId)
}

func (s *Handler) DeclineCandidate(w http.ResponseWriter, r *http.Request) error {
	candidateId, ok := mux.Vars(r)["candidateId"]
	if !ok {
		badRequestResponse(w, "OrderId not found")
		return nil
	}

	return s.candidateService.Decline(candidateId)
}

func (s *Handler) GetCandidate(w http.ResponseWriter, r *http.Request) error {
	candidateId, ok := mux.Vars(r)["candidateId"]
	if !ok {
		badRequestResponse(w, "OrderId not found")
		return nil
	}

	c, err := s.candidateRepository.GetById(candidateId)
	if err != nil {
		return err
	}
	if c == nil {
		notFoundResponse(w, "")
		return nil
	}

	return jsonResponse(w, candidateResponse{
		Id:      c.Id,
		Name:    c.Name,
		Email:   c.Email,
		Phone:   c.Phone,
		Address: c.Address,
		Status:  int(c.Status.Type),
	})
}

func (s *Handler) GetCandidateList(w http.ResponseWriter, r *http.Request) error {
	candidates, err := s.candidateRepository.GetAll()
	if err != nil {
		return err
	}

	var response candidateList
	for _, c := range candidates {
		response.candidateResponse = append(response.candidateResponse, candidateResponse{
			Id:      c.Id,
			Name:    c.Name,
			Email:   c.Email,
			Phone:   c.Phone,
			Address: c.Address,
			Status:  int(c.Status.Type),
		})
	}

	return jsonResponse(w, response)
}

func createHandlerFunc(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			var e *responseError
			if errors.As(err, e) {
				http.Error(w, e.Msg, e.Status)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func jsonResponse(w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8\"")
	resp, err := json.Marshal(r)
	if err != nil {
		return err
	}
	_, err = w.Write(resp)
	return err
}

func badRequestResponse(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func notFoundResponse(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusNotFound)
}
