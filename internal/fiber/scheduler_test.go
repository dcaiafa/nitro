package fiber

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestScheduler(t *testing.T) {
	var res strings.Builder

	s := NewScheduler()

	main := s.NewFiber(func() {
		res.WriteString("main started\n")
		child1 := s.NewFiber(func() {
			res.WriteString("child1 started\n")
			grandChild1 := s.NewFiber(func() {
				for i := 0; i < 3; i++ {
					fmt.Fprintf(&res, "grandChild1 block %v\n", i)
					s.Block(context.Background(), func(ctx context.Context) {
						time.Sleep(10 * time.Millisecond)
					})
				}
				res.WriteString("grandChild1 finished\n")
			})
			s.SwitchToNew(grandChild1)
			res.WriteString("child1 ended\n")
		})
		s.SwitchToNew(child1)
		child2 := s.NewFiber(func() {
			for i := 0; i < 3; i++ {
				fmt.Fprintf(&res, "child2 block %v\n", i)
				s.Block(context.Background(), func(ctx context.Context) {
					time.Sleep(5 * time.Millisecond)
				})
			}
			res.WriteString("child2 finished\n")
		})
		s.SwitchToNew(child2)
		res.WriteString("main ended\n")
	})

	s.Run(main)

	expected := `
main started
child1 started
grandChild1 block 0
child2 block 0
child1 ended
main ended
child2 block 1
grandChild1 block 1
child2 block 2
child2 finished
grandChild1 block 2
grandChild1 finished
`
	require.Equal(t, strings.TrimSpace(expected), strings.TrimSpace(res.String()))
}
