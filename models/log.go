package models

//Log represents a worker log
type Log struct {
	WorkerID string
	JobID    string
	Message  []byte
}
