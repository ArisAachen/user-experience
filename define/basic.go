package define

import "io"

type Operator interface {
	write(o io.Writer)
}