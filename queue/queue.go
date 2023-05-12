package queue

import (
	"fmt"
	"time"
)

func newQueue() *Queue {
	return &Queue{
		Jobs:          make(chan Job, 100),
		Present:       NewMutexChan(),
		WaitingList:   &LinkedList{},
		executionTerm: 180,
	}
}

type Queue struct {
	Jobs          chan Job    // 곧 실행 될 목록
	Present       *MutexChan  // 현재 진행중 목록 채널로 되어있음
	WaitingList   *LinkedList // 실행 대기 (순서 이동 가능) : 링크드 리스트로 간다아아아악
	executionTerm int         // job이 실행되는 주기
	//workers       int        // 몇개 씩 병렬처리할거임?
	//name          string     // 만들어진 큐이름 (중복일 경우 worker를 더 만듦
	// 사전작업 함수
}

func (q *Queue) AddJob() {

}
func (q *Queue) Run() { // consume
	for {
		select {
		case job := <-q.Jobs:
			q.Present.Add(job.Event) // 실행 목록 관리를 위한 적재
			job.Execute()            // go routine 사용금지
		}

		// 실행주기
		time.Sleep(time.Second * time.Duration(q.executionTerm))
	}
}
func (q *Queue) Progress() map[string]interface{} {


	return map[string]interface{}{
		"now_working":  q.Present.GetListAll(),
		"waiting_list": "",
	}
}

// Job!!

type Job struct {
	name       string
	Event      string //  run, delete, continue
	Done       bool
	number     int
	preprocess bool
}

func (j Job) Execute() {
	switch j.Event {
	case "run":
	case "delete":
	case "continue":
	default:
		fmt.Println("job name is not exist.")
	}
}
