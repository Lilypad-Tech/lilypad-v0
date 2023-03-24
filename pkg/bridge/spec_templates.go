package bridge

import (
	"encoding/json"
	"fmt"

	"github.com/bacalhau-project/bacalhau/pkg/model"
)

type specTemplateFunction func(map[string]string) (model.Spec, error)

func ensureArgs(args map[string]string, keys ...string) error {
	for _, k := range keys {
		if _, ok := args[k]; !ok {
			return fmt.Errorf("missing argument %s", k)
		}
	}
	return nil
}

func getFastSpec(args map[string]string) (model.Spec, error) {
	spec := fastSpec
	return spec, nil
}

func getWaterlilySpec(args map[string]string) (model.Spec, error) {
	if err := ensureArgs(args, "prompt", "artistid", "imageid"); err != nil {
		return model.Spec{}, err
	}

	spec := waterlilySpec
	combinedPrompt := fmt.Sprintf("%s, in the style of %s", args["prompt"], args["artistid"])
	spec.Docker = getWaterlilyDockerSpec(
		combinedPrompt,
		args["artistid"],
		args["imageid"],
	)
	return spec, nil
}

var templateFunctions map[string]specTemplateFunction = map[string]specTemplateFunction{
	"fast":      getFastSpec,
	"waterlily": getWaterlilySpec,
}

// first let's just convert the data to a map[string]string
// and see if there is a "_lilypad_template" key
// if yes then let's run the template function passing the values
// this means we can have contracts in the wild that don't need
// upgrading because we can re-depoloy the bridge
// also it means that JSON wrangling in solidity is not required
// amd that for private deployments (like waterlily) - the details
// of the job are not public
// if we fail to find a "_lilypad_template" key then we assume
// it is normal model.Spec JSON
func convertContractJobSpec(eventData []byte) (model.Spec, error) {
	data := map[string]string{}
	err := json.Unmarshal(eventData, &data)
	if err != nil {
		return model.Spec{}, err
	}
	templateName, ok := data["_lilypad_template"]
	if !ok {
		// we should assume it's a raw format
		spec := model.Spec{}
		err = json.Unmarshal(eventData, &spec)
		return spec, err
	} else {
		// this means we should pass the data through to a template function
		templateHandler, ok := templateFunctions[templateName]
		if !ok {
			return model.Spec{}, fmt.Errorf("unknown template name: %s", templateName)
		}
		return templateHandler(data)
	}
}
