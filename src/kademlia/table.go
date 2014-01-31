package kademlia

import (
	//stuff
	"container/list";
)

const ListSize = 10; //how many buckets?

type Table struct {
	owner Contact; 
	buckets [IDBytes*8]*List;
}

//make a table
func NewTable(owner_node Contact) (table *Table) {
	table = new(Table);
	table.owner = owner_node;
	//initialize buckets
	for i := 0;  i < IDBytes*8; i++{
		buckets[i] := list.New();
	}
	return;
}