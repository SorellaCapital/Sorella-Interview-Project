package multicast


import (
	"net"
	"net/http"
	"log"
	"net/rpc"
)

type Master struct {
	//concurrency is the number of worker nodes sharing state
	Concurrency int
	// Your definitions here.

}

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (m *Master) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	//sockname := coordinatorSock()
	//os.Remove(sockname)
	//l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}


//
// create a Master, will be called in main.go
// Pass in other arguments as you see fit
//
func MakeMaster(concurrency int) *Master {
	m := Master{
		Concurrency: concurrency,
	}

	// Your code here.


	m.server()
	return &m
}
