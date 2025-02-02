package events

import (
	"github.com/runatlantis/atlantis/server/events/models"
	"github.com/runatlantis/atlantis/server/logging"
	"github.com/runatlantis/atlantis/server/metrics"
)

type InstrumentedProjectCommandBuilder struct {
	ProjectCommandBuilder
	Logger logging.SimpleLogging
}

func (b *InstrumentedProjectCommandBuilder) BuildApplyCommands(ctx *CommandContext, comment *CommentCommand) ([]models.ProjectCommandContext, error) {
	scope := ctx.Scope.SubScope("builder")

	timer := scope.Timer(metrics.ExecutionTimeMetric).Start()
	defer timer.Stop()

	executionSuccess := scope.Counter(metrics.ExecutionSuccessMetric)
	executionError := scope.Counter(metrics.ExecutionErrorMetric)

	projectCmds, err := b.ProjectCommandBuilder.BuildApplyCommands(ctx, comment)

	if err != nil {
		executionError.Inc(1)
		b.Logger.Err("Error building apply commands: %s", err)
	} else {
		executionSuccess.Inc(1)
	}

	return projectCmds, err

}
func (b *InstrumentedProjectCommandBuilder) BuildAutoplanCommands(ctx *CommandContext) ([]models.ProjectCommandContext, error) {
	scope := ctx.Scope.SubScope("builder")

	timer := scope.Timer(metrics.ExecutionTimeMetric).Start()
	defer timer.Stop()

	executionSuccess := scope.Counter(metrics.ExecutionSuccessMetric)
	executionError := scope.Counter(metrics.ExecutionErrorMetric)

	projectCmds, err := b.ProjectCommandBuilder.BuildAutoplanCommands(ctx)

	if err != nil {
		executionError.Inc(1)
		b.Logger.Err("Error building auto plan commands: %s", err)
	} else {
		executionSuccess.Inc(1)
	}

	return projectCmds, err

}
func (b *InstrumentedProjectCommandBuilder) BuildPlanCommands(ctx *CommandContext, comment *CommentCommand) ([]models.ProjectCommandContext, error) {
	scope := ctx.Scope.SubScope("builder")

	timer := scope.Timer(metrics.ExecutionTimeMetric).Start()
	defer timer.Stop()

	executionSuccess := scope.Counter(metrics.ExecutionSuccessMetric)
	executionError := scope.Counter(metrics.ExecutionErrorMetric)

	projectCmds, err := b.ProjectCommandBuilder.BuildPlanCommands(ctx, comment)

	if err != nil {
		executionError.Inc(1)
		b.Logger.Err("Error building plan commands: %s", err)
	} else {
		executionSuccess.Inc(1)
	}

	return projectCmds, err

}
