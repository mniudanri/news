package jobs

// NOTE: queue for background process
// number of worker can be set from initialization
import(
	"log"
)

type Job interface {
	Process()
}

type Worker struct {
	WorkerId int
	Status chan bool
	JobChannel chan Job
}

type JobQueue struct {
	Workers []*Worker
	JobChannel chan Job
	Status chan bool
}

func CreateWorker(workerId int, jobChan chan Job) *Worker {
	return &Worker{
		WorkerId : workerId,
		Status: make(chan bool),
		JobChannel: jobChan,
	}
}

func (w *Worker) Run() {
    log.Print("Run worker id ", w.WorkerId)
	go func() {
		for {
			select {
				case job := <- w.JobChannel:
					log.Print("Running job at id ", w.WorkerId)
					job.Process()
				case <-w.Status:
					log.Print("Job Done at id ", w.WorkerId )
					return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	w.Status <- true
}

func InitJobQueue(n int) JobQueue {
	// NOTE: make array worker's cap and length based on number worker requested
	workers := make([]*Worker, n, n)
	jobChannel := make(chan Job)

	for i := 0; i < n; i++ {
		workers[i] = CreateWorker(i, jobChannel)
	}

	return JobQueue {
		Workers: workers,
		JobChannel: jobChannel,
		Status: make(chan bool),
	}
}

func (queue *JobQueue) Push(job Job) {
	queue.JobChannel <- job
}

func (queue *JobQueue) Stop() {
	queue.Status <- true
}

func (queue *JobQueue) Start() {
	go func() {
		for i := 0; i < len(queue.Workers); i++ {
			queue.Workers[i].Run()
		}
	}()

	go func() {
		for {
			select {
				case <-queue.Status:
					for i := 0; i < len(queue.Workers); i++ {
						queue.Workers[i].StopWorker()
					}
					return
			}
		}
	}()
}