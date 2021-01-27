package format

import (
	"github.com/geekerstar/libv1/av/avutil"
	"github.com/geekerstar/libv1/format/aac"
	"github.com/geekerstar/libv1/format/flv"
	"github.com/geekerstar/libv1/format/mp4"
	"github.com/geekerstar/libv1/format/rtmp"
	"github.com/geekerstar/libv1/format/rtsp"
	"github.com/geekerstar/libv1/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
