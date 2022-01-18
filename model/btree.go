package model

type Node struct { //创建一个左右斜树
	Left  *Node
	right *Node
	Data  interface{}
}
