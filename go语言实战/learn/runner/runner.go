package runner

//runner包管理处理任务的运行和生命周期
import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner在给定的超时时间内执行一组任务,
//并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//interrupt通道报告从操作系统
	//发送的信号
	interrupt chan os.Signal
	//complete通道报告处理任务已经完成
	complete chan error
	//timeout报告处理任务已经超时
	timeout <-chan time.Time
	//tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

//ErrTimeout会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

//ErrorInterrupt会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

//New返回一个新的准备使用的Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add将一个任务附加到Runner上。这个任务是一个接收一个int类型的ID作为参数的函数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete: //任务完成时发出的信号
		return err
	case <-r.timeout: //超时时发出的信号
		return ErrTimeout
	}
}

//run执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}
		//执行已注册的任务
		task(id)
	}
	return nil
}

//验证是否接收到了中断信号
func (r *Runner) goInterrupt() bool {
	select {
	//当中断事件被触发时发出的信号
	case <-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
		//继续正常运行
	default:
		return false
	}
}
