// wap in go prints "from 1" every 2 seconds and "from 2" every 3 seconds.
// select picks the first channel that is ready and receives from it. if more than one of the channels are ready then it
// randomly picks which one to receive from. if none of the channels are ready. the statement blocks until one becomes available.
