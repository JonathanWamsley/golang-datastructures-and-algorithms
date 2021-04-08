![go test](https://github.com/JonathanWamsley/golang-datastructures-and-algorithms/actions/workflows/go.yml/badge.svg)
![lint](https://github.com/JonathanWamsley/golang-datastructures-and-algorithms/actions/workflows/golangci-lint.yml/badge.svg)

# Golang Data Structures and Algorithms

This repo is for learning purposes only. I am referencing and old java book, data and algorithms made easy I used at uni.

### Linked List

- What: A data structure that consist of many nodes which are either singly or doubly linked. Each node stores data and has a pointer or a link connecting to the next node. When each node has a one way connection pointing only to the next node it is called a linked list. When each node has a two connections pointing to the prev and next node, it is a double linked list. The node value at the front in the end will point to (nil, null, none) indicating ends.

- Why: Linked list are good if you want to add or remove items at ends only. This makes them great for stacks and queues. They also don't need additional wasted memory allocated, so they can grow as needed.

- Why not: Linked list are bad at finding elements that are not on the ends. This is because a sequential search has to be done, iterating from one end to another. They also require additional memory because they have so store pointer values.

- Implementation: They way I implemented a linked list and double linked list is by creating struct that contains a pointer to the head node. Additional information can also be stored, such as the amount of nodes or the tail of a node in the case of a double linked list.

- Performance of methods for linked list
    - O(1)
        - New - creates an empty linked list
        - Insert - inserts Data at the front
        - DeleteHead - deletes the first node
        - Len - returns the amount of nodes
    - O(n)
        - Append - creates a new node at the end
            - Append could be made to O(1) if linked list had a tail pointer
        - Delete - searches for a node and removes it
        - String - prints the Data of each node

- Implemented methods for double linked list
    - O(1)
        - New - creates an empty linked list
        - Insert - inserts Data at the front
        - Append - creates a new node at the end        
        - DeleteHead - deletes the first node
        - DeleteTail - deletes the last node
        - Len - returns the amount of nodes
    - O(n)
        - Delete - searches for a node and removes it
        - String - prints the Data of each node
    

