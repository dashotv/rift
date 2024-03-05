package definitions

type JobService interface {
	Index(Request) Response
	Create(Request) Response
}

// Job tracks jobs in the system
// model: true
type Job struct {
	Kind     string        `json:"kind" bson:"kind"`
	Args     string        `json:"args" bson:"args"`
	Status   string        `json:"status" bson:"status"`
	Queue    string        `json:"queue" bson:"queue"`
	Attempts []*JobAttempt `bson:"attempts" json:"attempts"`
}

// JobAttempt tracks the attempts made to process a job
type JobAttempt struct { // struct
	StartedAt  string   `bson:"started_at" json:"started_at"`
	Duration   float64  `bson:"duration" json:"duration"`
	Status     string   `bson:"status" json:"status"`
	Error      string   `bson:"error" json:"error"`
	Stacktrace []string `bson:"stacktrace" json:"stacktrace"`
}
