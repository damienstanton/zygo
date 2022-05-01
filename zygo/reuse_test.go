package zygo

import (
	"testing"

	"github.com/glycerine/goconvey/convey"
	"github.com/shurcooL/go-goon"
)

func Test101ConversionToAndFromMsgpackAndJson(t *testing.T) {

	convey.Convey(`
      SexpToGo() should notice when it sees the same hash/with-shadow-struct re-used,
      and doesn't generate a 2nd shadow struct but instead re-uses the prior one
`, t, func() {
		event := `(def reUseMe (snoopy id:123));(def dad (hornet id:8 friends:[reUseMe])); (def mom (hellcat id:7 friends:[reUseMe]));(setOfPlanes flyers:[mom dad])`
		env := NewZlisp()
		defer env.Close()

		env.StandardSetup()
		env.ImportDemoData()

		x, err := env.EvalString(event)
		panicOn(err)
		P("\n x = %#v /\n\n string: '%s'\n", x, x.SexpString(nil))

		var set SetOfPlanes
		_, err = SexpToGoStructs(x, &set, env, nil, 0, &set)
		panicOn(err)
		P("\n set = %#v\n", set)
		goon.Dump(set)
		shared := &Snoopy{
			Plane: Plane{
				ID: 123,
			}}
		_ = shared
		convey.So(&set, convey.ShouldResemble, &SetOfPlanes{Flyers: []Flyer{&Hellcat{Plane: Plane{ID: 7, Friends: []Flyer{shared}}}, &Hornet{Plane: Plane{ID: 8, Friends: []Flyer{shared}}}}})

		// should actually *be* the same struct pointed to.
		ptr0 := set.Flyers[0].(*Hellcat).Friends[0].(*Snoopy)
		ptr1 := set.Flyers[1].(*Hornet).Friends[0].(*Snoopy)
		P("ptr0 = %p, ptr1 = %p", ptr0, ptr1)
		convey.So(ptr0, convey.ShouldEqual, ptr1)
	})
}
