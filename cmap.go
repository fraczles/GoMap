package main


type emap struct{
    sharedMap map[string]int
    stop chan bool
    //Several channels
}

//Constructor?
func (c *emap) NewChannelMap() *emap {
    var ret EmergingMap
    x := emap{}
    return &x
}

//Implement interface functions
func (c *emap) Listen() {
    for {
        select {
        case <- //stuff
        case <-c.stop:
            return
        }
    }
}
func (c *emap) Stop() {
    c.listen <- false
}
func (c *emap) Reduce(functor ReduceFunc, accum_str string, accum_int int) (string, int) {
    //Code here
}
func (c *emap) AddWord(word string) {
    //Code here
}
func (c *emap) GetCount(word string) int {
    //Code here
}
