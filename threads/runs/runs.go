package runs

import (
	"context"
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/internal/client"
	"github.com/CharLemAznable/dashscope-sdk-go/threads"
	"github.com/CharLemAznable/dashscope-sdk-go/threads/messages"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gvalid"
	"time"
)

type RunParam struct {
	AssistantId            string                   `json:"assistant_id" v:"required"`
	Model                  string                   `json:"model,omitempty"`
	Instructions           string                   `json:"instructions,omitempty"`
	AdditionalInstructions string                   `json:"additional_instructions,omitempty"`
	AdditionalMessages     []*messages.MessageParam `json:"additional_messages,omitempty"`
	Tools                  []tools.Tool             `json:"tools,omitempty"`
	Metadata               map[string]string        `json:"metadata,omitempty"`
	Temperature            float64                  `json:"temperature,omitempty"`
	Stream                 bool                     `json:"stream"`
	MaxPromptTokens        int64                    `json:"max_prompt_tokens,omitempty"`
	MaxCompletionTokens    int64                    `json:"max_completion_tokens,omitempty"`
	TruncationStrategy     TruncationStrategy       `json:"truncation_strategy,omitempty"`
	ToolChoice             interface{}              `json:"tool_choice,omitempty"`
	ResponseFormat         interface{}              `json:"response_format,omitempty" v:"in:json_object"`
	ParallelToolCalls      bool                     `json:"parallel_tool_calls,omitempty"`
}

type ThreadAndRunParam struct {
	RunParam
	Thread *threads.ThreadParam `json:"thread,omitempty"`
}

type SubmitToolOutputsParam struct {
	ToolOutputs []tools.ToolOutput `json:"tool_outputs" v:"required"`
	Stream      bool               `json:"stream"`
}

//goland:noinspection GoUnusedExportedFunction
func Create(ctx context.Context, threadId string, param *RunParam) (Run, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	fillDefaultParam(param)
	if err := validateRunParam(ctx, param); err != nil {
		return nil, err
	}
	checkNonStreamParam(ctx, &param.Stream)
	return client.Post(ctx, NewRunFromJson, "/threads/"+threadId+"/runs", param)
}

//goland:noinspection GoUnusedExportedFunction
func CreateThreadAndRun(ctx context.Context, param *ThreadAndRunParam) (Run, error) {
	fillDefaultParam(&param.RunParam)
	if err := validateRunParam(ctx, &param.RunParam); err != nil {
		return nil, err
	}
	checkNonStreamParam(ctx, &param.RunParam.Stream)
	return client.Post(ctx, NewRunFromJson, "/threads/runs", param)
}

//goland:noinspection GoUnusedExportedFunction
func CreateStream(ctx context.Context, threadId string, param *RunParam) (RunStream, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	fillDefaultParam(param)
	if err := validateRunParam(ctx, param); err != nil {
		return nil, err
	}
	checkStreamParam(ctx, &param.Stream)
	return callSteam(ctx, "/threads/"+threadId+"/runs", param), nil
}

//goland:noinspection GoUnusedExportedFunction
func CreateStreamThreadAndRun(ctx context.Context, param *ThreadAndRunParam) (RunStream, error) {
	fillDefaultParam(&param.RunParam)
	if err := validateRunParam(ctx, &param.RunParam); err != nil {
		return nil, err
	}
	checkStreamParam(ctx, &param.RunParam.Stream)
	return callSteam(ctx, "/threads/runs", param), nil
}

//goland:noinspection GoUnusedExportedFunction
func Retrieve(ctx context.Context, threadId string, runId string) (Run, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateRunId(ctx, runId); err != nil {
		return nil, err
	}
	return client.Get(ctx, NewRunFromJson, "/threads/"+threadId+"/runs/"+runId)
}

//goland:noinspection GoUnusedExportedFunction
func RetrieveStep(ctx context.Context, threadId string, runId string, stepId string) (RunStep, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateRunId(ctx, runId); err != nil {
		return nil, err
	}
	if err := validateStepId(ctx, stepId); err != nil {
		return nil, err
	}
	return client.Get(ctx, NewRunStepFromJson, "/threads/"+threadId+"/runs/"+runId+"/steps/"+stepId)
}

//goland:noinspection GoUnusedExportedFunction
func Update(ctx context.Context, threadId string, runId string, param *common.UpdateMetadataParam) (Run, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateRunId(ctx, runId); err != nil {
		return nil, err
	}
	return client.Post(ctx, NewRunFromJson, "/threads/"+threadId+"/runs/"+runId, param)
}

//goland:noinspection GoUnusedExportedFunction
func List(ctx context.Context, threadId string, param ...*common.ListParam) (*common.ListResult[Run], error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	listParam := &common.ListParam{}
	if len(param) > 0 && param[0] != nil {
		listParam = param[0]
	}
	return client.Get(ctx, common.ListResultFunc(NewRunFromJson), "/threads/"+threadId+"/runs", listParam)
}

//goland:noinspection GoUnusedExportedFunction
func ListSteps(ctx context.Context, threadId string, runId string, param ...*common.ListParam) (*common.ListResult[RunStep], error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateRunId(ctx, runId); err != nil {
		return nil, err
	}
	listParam := &common.ListParam{}
	if len(param) > 0 && param[0] != nil {
		listParam = param[0]
	}
	return client.Get(ctx, common.ListResultFunc(NewRunStepFromJson), "/threads/"+threadId+"/runs/"+runId+"/steps", listParam)
}

//goland:noinspection GoUnusedExportedFunction
func SubmitToolOutputs(ctx context.Context, threadId string, runId string, param *SubmitToolOutputsParam) (Run, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateRunId(ctx, runId); err != nil {
		return nil, err
	}
	if err := validateSubmitToolOutputsParam(ctx, param); err != nil {
		return nil, err
	}
	checkNonStreamParam(ctx, &param.Stream)
	return client.Post(ctx, NewRunFromJson, "/threads/"+threadId+"/runs/"+runId+"/submit_tool_outputs", param)
}

//goland:noinspection GoUnusedExportedFunction
func SubmitStreamToolOutputs(ctx context.Context, threadId string, runId string, param *SubmitToolOutputsParam) (RunStream, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateRunId(ctx, runId); err != nil {
		return nil, err
	}
	if err := validateSubmitToolOutputsParam(ctx, param); err != nil {
		return nil, err
	}
	checkStreamParam(ctx, &param.Stream)
	return callSteam(ctx, "/threads/"+threadId+"/runs/"+runId+"/submit_tool_outputs", param), nil
}

func Wait(ctx context.Context, threadId string, runId string) (run Run, err error) {
	for {
		run, err = Retrieve(ctx, threadId, runId)
		if err != nil {
			return
		}
		if run.GetStatus() == RunCancelled ||
			run.GetStatus() == RunFailed ||
			run.GetStatus() == RunCompleted ||
			run.GetStatus() == RunExpired ||
			run.GetStatus() == RunRequiresAction {
			return
		}
		time.Sleep(time.Second)
	}
}

func validateThreadId(ctx context.Context, threadId string) error {
	return gvalid.New().
		Rules("required").
		Messages("threadId is required").
		Data(threadId).Run(ctx)
}

func fillDefaultParam(param *RunParam) {
	param.ResponseFormat = "json_object"
}

func validateRunParam(ctx context.Context, param *RunParam) error {
	return gvalid.New().Data(param).Run(ctx)
}

func checkNonStreamParam(ctx context.Context, stream *bool) {
	if *stream {
		g.Log().Warning(ctx, "call non-stream method with stream=true, changed to false")
		*stream = false
	}
}

func checkStreamParam(ctx context.Context, stream *bool) {
	if !*stream {
		g.Log().Warning(ctx, "call stream method with stream=false, changed to true")
		*stream = true
	}
}

func validateRunId(ctx context.Context, runId string) error {
	return gvalid.New().
		Rules("required").
		Messages("runId is required").
		Data(runId).Run(ctx)
}

func validateStepId(ctx context.Context, stepId string) error {
	return gvalid.New().
		Rules("required").
		Messages("stepId is required").
		Data(stepId).Run(ctx)
}

func validateSubmitToolOutputsParam(ctx context.Context, param *SubmitToolOutputsParam) error {
	return gvalid.New().Data(param).Run(ctx)
}

func callSteam(ctx context.Context, url string, data ...interface{}) RunStream {
	stream, cli := newRunStream(), client.Client(ctx)
	g.Go(ctx, func(ctx context.Context) {
		eventSource := cli.PostEventSource(ctx, url, data...)
		defer func() {
			eventSource.Close()
			stream.Close(eventSource.Err())
		}()

		for event := range eventSource.Event() {
			stream.Push(newRunStreamEvent(event))
		}

	}, func(ctx context.Context, exception error) {
		g.Log().Errorf(ctx, "%+v", exception)
	})
	return stream
}
