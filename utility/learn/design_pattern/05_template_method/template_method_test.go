package _5_template_method

import "testing"

type WorkInterface interface {
	GetUp()
	Working()
	Sleep()
}
type Worker struct {
	WorkInterface
}

// 日常流程
func (w *Worker) DailyAction() {
	w.GetUp()
	w.Working()
	w.Sleep()
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}

type StudentWorker struct {
}

func (s StudentWorker) GetUp() {
	println("同学，起床了")
}

func (s StudentWorker) Working() {
	println("同学，上课了")
}

func (s StudentWorker) Sleep() {
	println("同学，睡觉了")
}

type Programmer struct{}

func (p Programmer) GetUp() {
	println("程序员，起床了")
}

func (p Programmer) Working() {
	println("程序员，写代码啦")
}

func (p Programmer) Sleep() {
	println("程序员，下班了 睡觉啦")
}

func TestWorkder(t *testing.T) {
	w := NewWorker(&StudentWorker{})
	w.DailyAction()
	w = NewWorker(&Programmer{})
	w.DailyAction()
}
