/*
Package mod はモジュール
*/
package mod

const (
	// Example1 は名称その1
	Example1 Name = "example1"
	// Example2　は名称その2
	Example2 Name = "example2"
)

const (
	// Id0 id 0
	Id0 ID = 0
	// Id1 id 1
	Id1 ID = 1
	// Id2 id 2
	Id2 ID = 2
)

// ID は Example を一意に決めるための識別子
type ID int

// Name は名前
type Name string

// Example はこのお試し用で使うための構造体
type Example struct {
	id   ID
	name Name
}

// Id は id のゲッター
func (e Example) Id() ID {
	return e.id
}

// Name は name のゲッター
func (e Example) Name() Name {
	return e.name
}

// NewExample は Example のコンストラクタ
func NewExample(id ID, name Name) Example {
	return Example{
		id:   id,
		name: name,
	}
}
