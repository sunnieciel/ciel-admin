package _4_observer

import (
	"fmt"
	"testing"
)

type TeacherSubject struct {
	observers []Observer
	context   string
}

func NewTeacherSubject() *TeacherSubject {
	return &TeacherSubject{observers: make([]Observer, 0)}
}

// 添加观察者
func (s *TeacherSubject) AddPerson(o Observer) {
	s.observers = append(s.observers, o)
}
func (s *TeacherSubject) notify() {
	for _, o := range s.observers {
		o.Update(s) //通知所有观察者更新自己的状态
	}
}

// 主题的更新方法
func (s TeacherSubject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

// 观察者接口
type Observer interface {
	Update(*TeacherSubject)
}
type Boy struct{}

func (m Boy) Update(bossStatus *TeacherSubject) {
	fmt.Println(bossStatus.context, " 男孩们学习啦")
}

type Girl struct{}

func (g Girl) Update(status *TeacherSubject) {
	fmt.Println(status.context, " 女孩们回家啦")
}
func TestObserver(t *testing.T) {
	teacherSubject := NewTeacherSubject()
	teacherSubject.AddPerson(Boy{})
	teacherSubject.AddPerson(Girl{})
	teacherSubject.UpdateContext("老师下课了")
}
