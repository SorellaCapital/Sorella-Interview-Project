package main

import (
	"github.com/SorellaCapital/Sorella-Interview-Project/multicast"
	//"math/rand"
	//"time"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	// Spawning master node with 5 workers, can change as you see fit
	m := multicast.MakeMaster(5)
	
	// Spawning worker nodes
	for i := 0; i < m.Concurrency; i++ {
		// Spawning as many worker threads as there are in the concurrency
		// parameter of the master node
		go func(id int){
			wk := multicast.MakeWorker(id)
			w := wk.MakeGUI(a)
			w.Show()

			
			//Uncomment to see how state changes in worker threads;
			/*t := time.NewTicker(time.Second)
			for range t.C {
				wk.ChangeNumber(wk.T.N2 + rand.Intn(5))
				wk.UpdateNumber(wk.GUI)
			}*/

		}(i)

	}
	
	a.Run()
	
}
