package data_structures

// Linked List Implementation

type Node struct{
	data int
	next *Node
}
var head *Node
func Add(elem int) {
	if head==nil {
		head=new(Node)
		head.data=elem
	}else{
		temp:=head
		for temp.next!=nil{
			temp=temp.next
		}
		newNode:=new(Node)
		newNode.data=elem
		temp.next=newNode
	}
}

func Search(elem int) *Node {
	temp:=head
	for temp!=nil{
		if temp.data==elem {
			return temp
		}
		temp=temp.next
	}
	return nil
}