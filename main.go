package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func main() {
	list := linkedlist{}
	list.insert(5)
	list.insert(4)
	list.insert(3)
	list.insert(2)
	list.insert(1)
	list.addAthead(25)
	list.addAthead(10)
	list.addAthead(10)
	list.delete(6)
	list.addafter(11, 67)
	list.addbefore(5, 2)

	list.dltend()
	list.dlthead()
	list.reveres()
	list.display()
	fmt.Println()
	fmt.Println(list.sum())

}

// insert element

func(list *linkedlist)insert(value int){
	newnode:=&node{data: value}
	temp:=list.head

	if temp==nil{
		list.head=newnode
		return
	}
	for temp.next!=nil{
		temp=temp.next
	}
	temp.next=newnode
	
}

// add node at head

func (list *linkedlist) addAthead(value int){
	newnode:=&node{data: value }
	if list.head==nil{
		list.head=newnode
	}else{
		newnode.next=list.head
		list.head=newnode
	}
}

// delete target in a node

func (list *linkedlist) delete(target int) {
	temp := list.head
	if temp.data == target {
		temp.next = temp.next.next
		return
	}
	for temp.next != nil {
		if temp.next.data == target {

			temp.next = temp.next.next
			if temp.next == nil {
				return
			}
		}
		temp = temp.next
	}
}

// delete node at head

func (list *linkedlist) dlthead() {
	list.head = list.head.next

}

// delete node at tail

func (list *linkedlist) dltend() {
	temp := list.head
	for temp.next != nil {
		if temp.next.next == nil {
			temp.next = nil
			break
		}
		temp = temp.next
	}
}

// add a node after a target node

func (list linkedlist) addafter(value, target int) {
	newnode := &node{value, nil}
	temp := list.head
	if temp.data == target {
		newnode.next = list.head.next
		list.head.next = newnode
		return
	}

	for temp.data != target {
		if temp.next == nil {
			fmt.Println("not found")
			break
		}
		temp = temp.next
	}
	if temp.next!=nil{
	newnode.next = temp.next
	temp.next = newnode
	}
}


// add before a target node

func(list *linkedlist)addbefore(value,target int){
	newnode:=&node{value,nil}
	temp:=list.head
	if temp.data==target{
		newnode.next=list.head
		list.head=newnode
	}
	for temp.next.data!=target{
		temp=temp.next
	}
	newnode.next=temp.next
	temp.next=newnode
}

// reverese a linked list

func(list *linkedlist)reveres(){
	var next,prev *node
	temp:=list.head
	for temp.next!=nil{
		next=temp.next
		temp.next=prev
		prev=temp
		temp=next
	}
	list.head=temp
	temp.next=prev
}

// display linkedlist

func (list *linkedlist) display() {
	temp := list.head
	if temp == nil {
		fmt.Println("empty list")
		return
	}
	for temp!=nil{
		fmt.Printf("-->%d",temp.data)
		temp=temp.next
	}
}

// sum of node

func (list *linkedlist) sum() int {
	sum := 0
	temp := list.head
	for temp != nil{
		sum=sum+temp.data
		temp=temp.next
	}

return sum
}

