package kademlia
// Contains definitions mirroring the Kademlia spec. You will need to stick
// strictly to these to be compatible with the reference implementation and
// other groups' code.

import "net"


// Host identification.
type Contact struct {
    NodeID ID
    Host net.IP
    Port uint16
}


// PING
type Ping struct {
    Sender Contact
    MsgID ID
}

type Pong struct {
    MsgID ID
    Sender Contact
}

func (k *Kademlia) Ping(ping Ping, pong *Pong) error {
    // This one's a freebie.
    if err = k.HandleRPC(&args.RPCHeader, &response.RPCHeader); err == nil {
    	log.Printf("ping from: %s\n", args.RPCHeader) 
    return
}

// RPC Helper functions

type RPCHeader struct {
  Sender *Contact;
  NetworkId string;
}

func (k *Kademlia) HandleRPC(request, response *RPCHeader) error {
  if request.NetworkId != k.NetworkId {
  	log.Printf("Network IDs do not match"0)
	return error
  }
  if request.Sender != nil {
    k.table.Update(request.Sender);
  }
  response.Sender = &k.routes.node;
  return nil;
}






// STORE
type StoreRequest struct {
    Sender Contact
    MsgID ID
    Key ID
    Value []byte
}

type StoreResult struct {
    MsgID ID
    Err error
}

func (k *Kademlia) Store(req StoreRequest, res *StoreResult) error {
    if req.Value > [max_size_allowed] {
	return error.String("Value too large to store") }

//  k.[store](req.Key, req.Value)
    res.MsgID = CopyID(req.MsgID)
    return nil
}

// FIND_NODE
type FindNodeRequest struct {
    Sender Contact
    MsgID ID
    NodeID ID
}

type FoundNode struct {
    IPAddr string
    Port uint16
    NodeID ID
}

type FindNodeResult struct {
    MsgID ID
    Nodes []FoundNode
    Err error
}

func (k *kademlia) FindNode(req FindNodeRequest, res *FindNodeResult) error{
	Nodes := kc.kad.routes.FindClosest(req.NodeID, ListSize);
	res.Nodes = make([]FoundNode, Nodes.Len());
	
	for i := 0; i < Nodes.Len(); i++{
	//get node info
	}
	return;
}


// FIND_VALUE
type FindValueRequest struct {
    Sender Contact
    MsgID ID
    Key ID
}

// If Value is nil, it should be ignored, and Nodes means the same as in a
// FindNodeResult.
type FindValueResult struct {
    MsgID ID
    Value []byte
    Nodes []FoundNode
    Err error
}

func (k *Kademlia) FindValue(req FindValueRequest, res *FindValueResult) error {
    // TODO: Implement.
    return nil
}

