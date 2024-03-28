package definitions

type JobService interface {
	Create(Request) JobResponse
}

type JobResponse struct{}
