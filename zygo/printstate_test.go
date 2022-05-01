package zygo

import (
	"testing"

	"github.com/glycerine/goconvey/convey"
)

func Test040SeenMapWorks(t *testing.T) {

	convey.Convey(`To allow cycle detection, given a set of pointers of various types, Seen should set and note when they have seen.`, t, func() {
		ps := NewPrintState()
		a := &SexpPair{}
		b := &SexpPointer{}
		d := &SexpStr{}
		convey.So(ps.GetSeen(a), convey.ShouldBeFalse)
		convey.So(ps.GetSeen(b), convey.ShouldBeFalse)
		convey.So(ps.GetSeen(d), convey.ShouldBeFalse)

		ps.SetSeen(a, "a")
		ps.SetSeen(b, "b")

		convey.So(ps.GetSeen(a), convey.ShouldBeTrue)
		convey.So(ps.GetSeen(b), convey.ShouldBeTrue)
		convey.So(ps.GetSeen(d), convey.ShouldBeFalse)
	})
}
