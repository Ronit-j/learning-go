package main

import "fmt"

type Vertex struct {
	X int
	Y *Vertex
}
func display_linked_list(v Vertex){	
	cur_temp := &v
	for cur_temp != nil {
		fmt.Println(cur_temp.X)
		cur_temp = cur_temp.Y
	}
}

func add_to_linked_list(ll Vertex, x int){
	cur_temp := &ll
	for cur_temp.Y != nil {
		fmt.Println(cur_temp.X)
		cur_temp = cur_temp.Y
	}
	cur_temp.Y = &Vertex{x, nil}
}

func delete_node_from_linked_list(ll *Vertex, x int){
	cur_temp := ll
	prev_temp := ll
	for cur_temp.Y != nil {
		if cur_temp.X == x{
			break
		}
		prev_temp = cur_temp
		cur_temp = cur_temp.Y
		
	}
	fmt.Println(cur_temp, prev_temp)
	prev_temp.Y = cur_temp.Y
	fmt.Println(prev_temp)
	
}

func main() {
	
	v1 := Vertex{1, nil}
	v2 := Vertex{2, &v1}
	v3 := Vertex{3, &v2}
	head := Vertex{-1, &v3}
	display_linked_list(head)
	add_to_linked_list(head, 5)
	fmt.Println("--------------------")
	delete_node_from_linked_list(&head, 3)
	fmt.Println(head)
	display_linked_list(head)
}

