package main

import (
	"fmt"
	"sync"
	"time"
)

//1.使用sync.once ，单例模式
//2.使用mutex
//3.使用rwmutex

//实际的单例对象
type Manager struct {
	name   string
	nameLk sync.Mutex
	ageLk  sync.RWMutex
	age    int
}

//初始化函数
func (m *Manager) Init() {
	m.name = "I'm a manager"
}
func (m *Manager) SetName(newName string) {
	m.nameLk.Lock()
	defer m.nameLk.Unlock()

	m.name = newName
}

func (m *Manager) SetAge(newAge int) {
	m.ageLk.Lock()
	defer m.ageLk.Unlock()
	m.age = newAge
}
func (m *Manager) GetAge() int {
	m.ageLk.RLock()
	defer m.ageLk.RUnlock()
	return m.age
}

//成员函数
func (m *Manager) PrintName() {
	fmt.Println(m.name)
}

//产生单例
type ManagerSingleton struct {
	initLk sync.Once
	m      *Manager
}

func (ms *ManagerSingleton) init() {
	fmt.Println("manager Init")
	ms.m = new(Manager)
}
func (ms *ManagerSingleton) getInstance() *Manager {
	ms.initLk.Do(ms.init)
	return ms.m
}

var manager ManagerSingleton
var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	manager.getInstance().Init()

	go manager.getInstance().PrintName()
	wg.Add(4)
	go manager.getInstance().PrintName()
	go func() {
		for i := 0; i < 10; i++ {
			manager.getInstance().SetAge(i)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(manager.getInstance().GetAge())
		}

	}()
	go func() {
		wg.Done()
		wg.Add(-3)
	}()
	wg.Wait()
	time.Sleep(10 * time.Second)
}
