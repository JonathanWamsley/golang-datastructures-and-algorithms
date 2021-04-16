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
    

### Stacks

- What: A Stack is data structure that mimics how you would add and remove a stack of plates in real life. You can only add a new plate on top, which is known as push, or you can remove the top most item, which is called pop. This is why stacks are referred to as a First In, Last Out (FILO) or Last In First Out (LIFO) data structure. Stacks can be implemented with a linked list, fixed size array, or a dynamically sized array.

- Why: Stacks are useful when you need store data and wish to get the most recently added data back in that order.

- implementation: Each stack has the push(to add data), pop(to remove and get), peek(to get without removal), and Empty(returns true if no data in stack).

- Performance of methods for a stacked using a fixed size array
    - O(1)
        - New - creates a new fixed stack with the specified capacity
        - Push - adds data to the stack
        - Pop - removes and returns the most recently pushed data
        - Peek - returns the most recently pushed daa
        - Empty - returns if the stack has no items
        - Full - returns if the stack has no more space
        - Len - returns the amount of items in stack
 
- Performance of methods for a stacked using a dynamically size array
    - O(1)
        - New - creates a new stack using a slice
        - Push - adds data to the stack
        - Pop - removes and returns the most recently pushed data
        - Peek - returns the most recently pushed data
        - Empty - returns if the stack has no items
        - Len - returns the amount of items in stack

- Performance of methods for a stacked using a linked list array
     - O(1)
        - New - creates a new stack using a linked list
        - Push - adds data to the stack and returns node back
        - Pop - removes and returns the most recently pushed data
        - Peek - returns the most recently pushed data
        - Empty - returns if the stack has no nodes
        - Len - returns the amount of nodes in stack

### Queues

- What: A queue is a data structure that mimics standing in line to get some food. New people arrive at the end of the line, which is known as enqueue. People at the front of the line can get served and are removed, which is known as dequeue. This is why queues are known as First In First Out(FIFO) or Last In Last Out (LILO). Queues can be implemented by using a double ended linked list, circular fixed array, or a circular dynamic array.

- Why: Queues are used when ever you have tasks or jobs that need to be done in a certain order or schedule.

- Implementation: Each queue has Enqueue(to add data), Dequeue(to remove data), and Front(to view the data the would be removed) properties

- Performance of methods for a queue using a double ended linked list
    - O(1)
        - Enqueue - adds a node to the end
        - Dequeue - returns and remove the first node
        - Front - returns the first node

- Performance of methods for a queue using a circular fixed array 
    - O(1)
        - Enqueue - adds item to end of queue
        - Dequeue - removes front item in queue
        - Front - returns the first element
        - Full - returns true if queue is at max capacity

- Performance of methods for a queue using a circular fixed array 
    - O(1)
        - Enqueue - adds item to end of queue
        - Dequeue - removes front item in queue
        - Front - returns the first element
        - Full - returns true if queue is at max capacity

- Performance of methods for a queue using a circular dynamic array
    - O(1)
        - Enqueue - adds item to end of queue
        - Dequeue - removes front item in queue
        - Front - returns the first element
