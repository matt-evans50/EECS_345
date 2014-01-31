package kademlia
// Contains the core kademlia type. In addition to core state, this type serves
// as a receiver for the RPC methods, which is required by that package.
import (
	//stuff
	"container/list";
)
// Core Kademlia type. You can put whatever state you want in this.
type Kademlia struct {
    info *Contact
    rtable *Table;
}

func NewKademlia() *Kademlia {
    // TODO: Assign yourself a random ID and prepare other state here.
    ret = new(Kademlia);
    ret.info = new(Contact);
    ret.info.NodeID = NewRandomID();
    ret.rtable = NewTable(ret.NodeID)
    return ret;
}

const ListSize = 10; //how many buckets?

type Table struct {
	owner *Contact; 
	buckets [IDBytes*8]*list.List;
}

//make a table
func NewTable(owner_node *Contact) (table *Table) {
	table = new(Table);
	table.owner = owner_node;
	//initialize buckets
	for i := 0;  i < IDBytes*8; i++{
		buckets[i] := list.New();
	}
	return;
}

//update table
func (table *Table) Update(node *Contact) {
	//how to initialize to nil?
	var node_holder Element = nil;

	//identify correct bucket
	bucket_num := node.NodeID.Xor(table.owner.NodeID).PrefixLen();

	//check if node already in list
	bucket = table.buckets[bucket_num];
	for i := bucket.Front(); i != nil; i = i.Next() {
		if i.Value.NodeID.Equals(node.NodeID {
			node_holder = i;
			break;
		}
	}
	//if old, move to end
	if node_holder!=nil {
		bucket.MoveToBack(node_holder);
	}

	//if new and list not full, add to end
	else if node_holder==nil && bucket.Len()!=ListSize {
		bucket.PushBack(node_holder);
	}
	//if new and list full, ping oldest
		//if oldest responds, do nothing
		//else drop oldest, add new to end
	else if node_holder==nil && bucket.Len()==ListSize {
		Ping(///front of list);
		if { //no pong
			bucket.Remove(bucket.Front());
			bucket.PushBack(node_holder);
		}
	}
	else{
		//error message
	}
}