package metrics

import (
	"runtime"
)

var (
	rtMem runtime.MemStats
)

// a few simple stats from go runtime, in your app, if you want these 
// add them as a metrics function
//
//		func init() {
//			metrics.AddMetricsFunction(metrics.GoRuntimeMetrics)
//		}
func GoRuntimeMetrics(snap *Snapshot) {
<<<<<<< HEAD
	snap.Ints["go_cgocall.ct"] = uint64(runtime.NumCgoCall())
	snap.Ints["go_numgoroutine.ct"] = uint64(runtime.NumGoroutine())
	// http://weekly.golang.org/pkg/runtime/#MemStats
	runtime.ReadMemStats(&rtMem)
	snap.Ints["go_memalloc.avg"] = rtMem.Alloc / 1048576
	snap.Ints["go_memtotalloc.avg"] = rtMem.TotalAlloc / 1048576
	snap.Ints["go_memsys.avg"] = rtMem.Sys / 1048576
}
=======

	snap.Ints["go_cgocall"] = uint64(runtime.NumCgoCall())
	snap.Ints["go_numgoroutine"] = uint64(runtime.NumGoroutine())
	// http://weekly.golang.org/pkg/runtime/#MemStats
	runtime.ReadMemStats(&rtMem)
	snap.Ints["go_memalloc"] = rtMem.Alloc / 1048576
	snap.Ints["go_memtotalloc"] = rtMem.TotalAlloc / 1048576
	snap.Ints["go_memsys"] = rtMem.Sys / 1048576
	
}
>>>>>>> 0b419bfe79f1924261590bb2e2af4087631e1a94
