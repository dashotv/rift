package definitions

type WorkerService interface {
	Enqueue(Request) Response
}
