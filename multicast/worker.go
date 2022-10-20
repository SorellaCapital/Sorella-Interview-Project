package multicast

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Your definition for Worker goes here -
// Tuple is already filled in as this is what the
// state is represented by on the GUI;
// GUI is the pointer to change the GUI interface itself,
// and ID is the worker ID
//
type Worker struct {
	T Tuple
	GUI *widget.Label
	ID int
}

type Tuple struct {
	N1 int
	N2 int
}

//Function to call to change worker state
func (wk *Worker) ChangeNumber(n int) {
	a,b := wk.T.N1, wk.T.N2
	wk.T.N1 = b
	wk.T.N2 = (a + n) % 100
}

//Function to call to implement changed worker state
// in the GUI
func (wk *Worker) UpdateNumber(num *widget.Label) {
	num.SetText(strconv.Itoa(wk.T.N1) + " " + strconv.Itoa(wk.T.N2))
}

// Returns pointer to initialization of a worker thread, called in main.go
func MakeWorker(id int) *Worker {
	wk := Worker {
		T: Tuple{0,0},
		ID: id,
	}

	fmt.Printf("Worker %d spawned \n", id)

	return &wk
}

// Initializes a GUI interface for each worker thread, called in main.go
func (wk *Worker) MakeGUI(a fyne.App) fyne.Window {
	w := a.NewWindow("Worker " + strconv.Itoa(wk.ID))
	num := widget.NewLabel(strconv.Itoa(wk.T.N1) + " " + strconv.Itoa(wk.T.N2))
	wk.GUI = num
	w.SetContent(num)
	w.Resize(fyne.NewSize(150,100))

	return w
}

//
// example function to show how to make an RPC call to the master node.
//
// the RPC argument and reply types are defined in rpc.go.
//
func CallExample() {

	// declare an argument structure.
	args := ExampleArgs{}

	// fill in the argument(s).
	args.X = 99

	// declare a reply structure.
	reply := ExampleReply{}

	// send the RPC request, wait for the reply.
	// the "Master.Example" tells the
	// receiving server that we'd like to call
	// the Example() method of struct Master.
	ok := call("Master.Example", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("reply.Y %v\n", reply.Y)
	} else {
		fmt.Printf("call failed!\n")
	}
}

//
// send an RPC request to the master, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	//sockname := coordinatorSock()
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
