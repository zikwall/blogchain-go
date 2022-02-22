package drop

// The Drop interface and the Droppable structure implement a custom solution for closing
// and destroying (clearing) resources after the application is terminated (including an emergency,
// except a power outage ^_^)
// The main use of the Drop interface is to free the resources that the implementor instance owns.
// The Drop interface can also be manually implemented for any custom data type.
// ```code
// 	func (conn *SomeConnection) Drop() error {
//		return conn.db.Close()
//	}
// ```
type Drop interface {
	Drop() error
}

// Debug interface implementation in DropDebug + Drop interface
// allows you to output user messages during resource cleanup
// ```code
// 	func (conn SomeConnection) DropMsg() string {
//		return "close something connection pool"
//	}
// ```
type Debug interface {
	DropMsg() string
}

type Droppable struct {
	droppers []Drop
}

func (n *Droppable) AddDroppers(droppers ...Drop) {
	for _, dropper := range droppers {
		n.AddDropper(dropper)
	}
}

func (n *Droppable) AddDropper(dropper Drop) {
	n.droppers = append(n.droppers, dropper)
}

func (n *Droppable) EachDroppers(callback func(Drop)) {
	for _, dropper := range n.droppers {
		callback(dropper)
	}
}
