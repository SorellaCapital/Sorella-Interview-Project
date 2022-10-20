# Sorella-Interview-Project

At Sorella, the problems we tackle are focused on taking deeply academic concepts and applying them as they benefit our use. As such, most of the development requires learning on the job in completely unfamiliar areas and referencing journal papers/niche sources to problem solve and puzzle when you get stuck.

I think it is only fitting that the interview test takes the form of a week long, open-resource coding project in an area tangential to what we do at Sorella. Your project will be to implement a total order multicast service using a distributed algorithm in Golang.

If you are unfamiliar with Golang, no worries! It is a relatively easy language to pick up and has an excellent community which makes it easy to just use Google when you get stuck. As a good primer, I recommend you complete the [online Go tutorial](https://go.dev/tour/welcome/1), especially spending time on concurrency. More resources are linked at the end.

## Implementation

A [theoretical overview of ISIS](https://www.cs.purdue.edu/homes/bb/cs542-15Spr/Birman-Reliable-Broadcast.pdf) will be helpful to look over before delving into the project. I also found the slides on total order multicast and the ISIS algorithm from this [slideshow](https://courses.grainger.illinois.edu/ece428/sp2021//assets/slides/lect8-after.pdf), developed by UIUC's Indranil Gupta to be beneficial.

You will need to implement a worker process that recieves requests and makes proposals to the master process which sends requests and distributes agreed messages to all nodes (slideshow slide 46). When an agreed sequence number is achieved, the master will broadcast this message to the nodes who will then change state accordingly. 

The worker is connected to a GUI process that represents a tuple. The initial state of the worker is the tuple (0, 0). Each message delivered to the worker should be a random integer from 0 to 25. The state change operation is to at this integer to the first value (mod 100) and then rotate the values. This operation is thus non-commutative – the sequence 3, 6, 20 will yield a different state than 20, 3, 6. As a result, should you succeed in your implementation, the state of all workers will stay the same. If you fail, the state will diverge as the messages aren't being broadcast in the agreed order.

First, the master sends a request to the workers. The workers then send a proposal of a random integer from 0 and 25 and a sequence number back to the master. Once all workers have responded, an agreed sequence number will be determined (look through slides/paper on how this is done) and the message will be broadcasted back to the workers in the proper order where they will change their state.

The workers will communicate with the master via RPC. I recommend looking at [Go's documentation](https://pkg.go.dev/net/rpc) for how this communication works. The outline for the worker and master files are included in the multicast folder; you will need to add to both files, as well as the RPC file and the main.go file to complete this project.

### Challenge

Implement a jitter parameter in the master, analagous to the sleep parameter in workers, to represent network delay in sending requests and agreed messages to workers.

## More Resources
* [Concurrency in Go Book](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) (pdf available online)
* [ISIS Algorithm Overview](https://studylib.net/doc/7830646/isis-algorithm-for-total-ordering-of-messages)
* [Worker Pool Video Example](https://www.youtube.com/watch?v=1iBj5qVyfQA)
* [Great Go Concurrency Tutorial Videos](https://www.youtube.com/playlist?list=PLsc-VaxfZl4do3Etp_xQ0aQBoC-x5BIgJ)
* [Helpful Goroutine Code Walkthrough (MIT 6.824)](https://www.youtube.com/watch?v=gA4YXUJX7t8&list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB&index=2) (skip to coding section)
* [TA Golang Concurrency Session (MIT 6.824)](https://www.youtube.com/watch?v=UzzcUS2OHqo&list=PLrw6a1wE39_tb2fErI4-WkMbsvGQk9_UB&index=5)

## Reach Out
Feel free to reach out to mozart@sorellacapital.com with any questions you have!
