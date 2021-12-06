package asyncjob

import (
	"context"
	"time"
)

//1 job nhu 1 hjam
//2 job có thể retries dc
//3 co thể config retries
//4 job manager sẽ quản lý nhieu job
type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout   = time.Second * 5
	defaultMaxRetryTime = 3
)

var defaultRetryTime = []time.Duration{time.Second * 1, time.Second * 3, time.Second * 5}

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

type JobConfig struct {
	MaxTimeOut time.Duration
	Retries    []time.Duration
}
func(js JobState) String() string{
	return []string{"Init","Running","Failed","TimeOut","Completed","RetryFailed"}[js]
}

type job struct {
	config JobConfig
	handler JobHandler
	state JobState
	retryIndex int
	stopChan chan bool
}
func NewJob(handler JobHandler) *job{
	return &job{
		config : JobConfig{
			MaxTimeOut: defaultMaxTimeout,
			Retries: defaultRetryTime,
		},
		handler: handler,
		state: StateInit,
		retryIndex: -1,
		stopChan: make(chan bool ),
	}
}

func (j *job) Execute(ctx context.Context) error{
	j.state =  StateRunning
	if err := j.handler(ctx); err!= nil{
		j.state = StateFailed
	}
	j.state = StateCompleted
	return nil
}

func (j *job)Retry (ctx context.Context) error{
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	 err:= j.handler(ctx)
	 if err== nil{
		 j.state = StateCompleted
	 }

	 if j.retryIndex == len(j.config.Retries) - 1{
		 j.state = StateRetryFailed
		 return err
	 }
	 j.state = StateFailed
	 return nil
}

func (j *job) State() JobState{ return j.state}
func (j *job) RetryIndex() int{ return j.retryIndex}

func(j *job) SetRetryDurations(times []time.Duration){
	if len(times) == 0 {
			return
	}
	j.config.Retries = times
}
//1h