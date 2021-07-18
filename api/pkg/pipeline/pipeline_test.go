package pipeline_test

import (
	"testing"

	"github.com/infraboard/workflow/api/pkg/pipeline"
	"github.com/stretchr/testify/assert"
)

func TestPipelineNextStep(t *testing.T) {
	sample := SamplePipeline()
	t.Log(sample)
	t.Log(sample.NextStep())
	sample.Stages[0].Steps[0].Success(map[string]string{"status": "ok"})
	t.Log(sample.NextStep())
	sample.Stages[0].Steps[1].Success(map[string]string{"status": "ok"})
	t.Log(sample.Stages[0].IsComplete())
	t.Log(sample.NextStep())

	sample.Stages[1].Steps[0].Success(map[string]string{"status": "ok"})
	t.Log(sample.NextStep())
	sample.Stages[1].Steps[1].Success(map[string]string{"status": "ok"})
	t.Log(sample.Stages[1].IsComplete())
	t.Log(sample.NextStep())

	t.Log(sample.IsComplete())
}

func TestStageNextStep(t *testing.T) {
	sample := SampleStage("stage01")
	t.Log(sample)
	t.Log(sample.NextStep())
	sample.Steps[0].Success(map[string]string{"status": "ok"})
	t.Log(sample.NextStep())
	sample.Steps[1].Success(map[string]string{"status": "ok"})
	t.Log(sample)
	t.Log(sample.NextStep())
	t.Log(sample.IsComplete())
}

func TestUpdateStep(t *testing.T) {
	should := assert.New(t)
	sample := SamplePipeline()
	s1 := pipeline.NewDefaultStep()
	s1.Key = "..1.1"
	s1.Action = "update01"
	err := sample.UpdateStep(s1)

	if should.NoError(err) {
		t.Log(sample)
	}
}

func SamplePipeline() *pipeline.Pipeline {
	p := pipeline.NewDefaultPipeline()
	p.AddStage(SampleStage("stage01"))
	p.AddStage(SampleStage("stage02"))
	return p
}

func SampleStage(name string) *pipeline.Stage {
	stage := pipeline.NewDefaultStage()
	stage.Name = name

	s1 := pipeline.NewDefaultStep()
	s1.Action = name + ".action01"
	s1.Key = "..1.1"
	s2 := pipeline.NewDefaultStep()
	s2.Action = name + ".action02"
	stage.AddStep(s1)
	stage.AddStep(s2)
	return stage
}
