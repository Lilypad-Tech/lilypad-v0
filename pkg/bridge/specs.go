package bridge

import (
	"fmt"
	"os"
	"strings"

	"github.com/filecoin-project/bacalhau/pkg/model"
)

var stableDiffusionSpec = model.Spec{
	Engine:    model.EngineDocker,
	Verifier:  model.VerifierNoop,
	Publisher: model.PublisherIpfs,
	Docker: model.JobSpecDocker{
		Image:      "ghcr.io/bacalhau-project/examples/stable-diffusion-gpu:0.0.1",
		Entrypoint: []string{"python", "main.py", "--o", "./outputs", "--p"},
	},
	Resources: model.ResourceUsageConfig{
		GPU: "1",
	},
	Outputs: []model.StorageSpec{
		{
			Name: "outputs",
			Path: "/outputs",
		},
	},
	Deal: model.Deal{
		Concurrency: 1,
	},
}

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

func getWaterlilyImage(artistid string) string {
	return fmt.Sprintf("algoveraai/sdprojectv2:%s", artistid)
}

func getWaterlilyEntrypoint(prompt string, imageid string) []string {
	escapedPrompt := strings.ReplaceAll(prompt, "\"", "\\\"")
	fullCommand := fmt.Sprintf(`
	(apt update && apt install -y curl) &
	(python main.py --o /outputs --seed '%s' --p "%s") &
	wait &&
		curl -o /upload.py https://raw.githubusercontent.com/bacalhau-project/WaterLily/main/scripts/upload.py &&
		python3 /upload.py /outputs;
	`, imageid, escapedPrompt)
	singleLineCommand := strings.ReplaceAll(fullCommand, "\n", " ")
	return []string{"bash", "-c", singleLineCommand}
}

func getWaterlilyEnv(imageid string) []string {
	chain := os.Getenv("CHAIN_ID") // hack!
	return []string{
		fmt.Sprintf("WATERLILY_JOB_ID=%s-%s", chain, imageid),
		fmt.Sprintf("FILESTORE_ACCESS_TOKEN=%s", os.Getenv("FILESTORE_TOKEN")),
	}
}

func getWaterlilyDockerSpec(prompt string, artistid string, imageid string) model.JobSpecDocker {
	return model.JobSpecDocker{
		Image:                getWaterlilyImage(artistid),
		Entrypoint:           getWaterlilyEntrypoint(prompt, imageid),
		EnvironmentVariables: getWaterlilyEnv(imageid),
	}
}

var waterlilySpec = model.Spec{
	Engine:    model.EngineDocker,
	Verifier:  model.VerifierNoop,
	Publisher: model.PublisherIpfs,
	Docker:    getWaterlilyDockerSpec("", "", ""),
	Resources: model.ResourceUsageConfig{
		GPU: "1",
	},
	Outputs: []model.StorageSpec{
		{
			Name: "outputs",
			Path: "/outputs",
		},
	},
	Deal: model.Deal{
		Concurrency: 1,
	},
	Network: model.NetworkConfig{
		Type: model.NetworkFull,
	},
}
