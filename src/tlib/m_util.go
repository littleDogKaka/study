package tlib

import "strconv"

// StoreDef 中间存储接口定义
type StoreDef interface {
	Set(v string)
	Get() int
	Run()
}

// Store 中间存储
type Store struct {
	in  string
	out int
}

// New 返回一个新的Store
func New(in string) *Store {
	return &Store{in: in}
}

// Set 设置输入值
func (m *Store) Set(v string) {
	m.in = v
}

// Get 读取处理结果
func (m *Store) Get() int {
	return m.out
}

// Run 执行转换
func (m *Store) Run() {
	m.out, _ = strconv.Atoi(m.in)
}
