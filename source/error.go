package source

import (
	"context"

	"github.com/MasterOfBinary/gobatch/batch"
)

type errorSource struct {
	err error
}

// Error returns a Source that returns an error and then closes immediately.
// It can be used as a mock Source.
func Error(err error) batch.Source {
	return &errorSource{
		err: err,
	}
}

// Read returns an error and then closes.
func (s *errorSource) Read(ctx context.Context, ps batch.PipelineStage) {
	ps.Error() <- s.err
	ps.Close()
}
