package transport

import (
	"errors"
	"github.com/gorilla/mux"
	"hrm/pkg/candidate/domain"
	"net/http"
)

type Handler struct {
	candidateService domain.CandidateService
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
	return nil
}

func (s *Handler) MakeOffer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Handler) HireCandidate(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Handler) DeclineCandidate(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Handler) GetCandidate(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Handler) GetCandidateList(w http.ResponseWriter, r *http.Request) error {
	return nil
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
