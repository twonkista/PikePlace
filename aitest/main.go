package main

// imported as openai
import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func GenerateText(prompt string) (string, error) {

	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")), // or set OPENAI_API_KEY in your env
	)

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-5-nano",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String("Try to be a fair betting assistant that helps users create betting pools and place wagers. " +
			"Respond only with valid JSON objects that conform to the data models. " +
			"When responding to requests about pools, include only the fields: ID, Title, Description, CreatorID, Status, SLTotal, SWTotal, Outcome, Level of Proof (LoP). " +
			"Here is the user prompt: " + prompt)},
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp.OutputText())
	return resp.OutputText(), nil
}

func main() {
	out, err := GenerateText("Mitra crash out during the 2008 economic recession?")

	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
