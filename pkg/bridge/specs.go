package bridge

import "github.com/bacalhau-project/bacalhau/pkg/model"

var fastSpec = model.Spec{
	Engine:    model.EngineDocker,
	Verifier:  model.VerifierNoop,
	Publisher: model.PublisherEstuary,
	Docker: model.JobSpecDocker{
		Image:      "ubuntu",
		Entrypoint: []string{"echo"},
	},
	Deal: model.Deal{
		Concurrency: 1,
	},
}
