package ferryman_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	. "github.com/marconi/ferryman"
	. "github.com/marconi/rivers"
)

func TestFerryman(t *testing.T) {
	Convey("should be able to move jobs from delayed to urgent", t, func() {
		urgent := NewQueue("styx", "urgent")
		delayed := NewQueue("acheron", "delayed")
		quit := Run(urgent, delayed)

		delayed.Push(NewJob(time.Now().UTC()))
		delayed.Push(NewJob(time.Now().UTC()))

		size, err := delayed.GetSize()
		So(err, ShouldEqual, nil)
		So(size, ShouldEqual, 2)

		time.Sleep(1 * time.Second)

		size, err = delayed.GetSize()
		So(err, ShouldEqual, nil)
		So(size, ShouldEqual, 0)

		size, err = urgent.GetSize()
		So(err, ShouldEqual, nil)
		So(size, ShouldEqual, 2)

		Reset(func() {
			close(quit)
			urgent.Destroy()
			delayed.Destroy()
		})
	})
}
