package cron

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/goplugin/plugin-common/pkg/services"

	"github.com/goplugin/pluginv3.0/v2/core/logger"
	"github.com/goplugin/pluginv3.0/v2/core/services/job"
	"github.com/goplugin/pluginv3.0/v2/core/services/pipeline"
)

// Cron runs a cron jobSpec from a CronSpec
type Cron struct {
	cronRunner     *cron.Cron
	logger         logger.Logger
	jobSpec        job.Job
	pipelineRunner pipeline.Runner
	chStop         services.StopChan
}

// NewCronFromJobSpec instantiates a job that executes on a predefined schedule.
func NewCronFromJobSpec(
	jobSpec job.Job,
	pipelineRunner pipeline.Runner,
	logger logger.Logger,
) (*Cron, error) {
	cronLogger := logger.Named("Cron").With(
		"jobID", jobSpec.ID,
		"schedule", jobSpec.CronSpec.CronSchedule,
	)
	if id := jobSpec.CronSpec.EVMChainID; id != nil {
		cronLogger = logger.With("evmChainID", id)
	}

	return &Cron{
		cronRunner:     cronRunner(),
		logger:         cronLogger,
		jobSpec:        jobSpec,
		pipelineRunner: pipelineRunner,
		chStop:         make(chan struct{}),
	}, nil
}

// Start implements the job.Service interface.
func (cr *Cron) Start(context.Context) error {
	cr.logger.Debug("Starting")

	_, err := cr.cronRunner.AddFunc(cr.jobSpec.CronSpec.CronSchedule, cr.runPipeline)
	if err != nil {
		cr.logger.Errorw(fmt.Sprintf("Error running cron job %d", cr.jobSpec.ID), "err", err)
		return err
	}
	cr.cronRunner.Start()
	return nil
}

// Close implements the job.Service interface. It stops this job from
// running and cleans up resources.
func (cr *Cron) Close() error {
	cr.logger.Debug("Closing")
	cr.cronRunner.Stop()
	return nil
}

func (cr *Cron) runPipeline() {
	ctx, cancel := cr.chStop.NewCtx()
	defer cancel()

	jobSpec := map[string]interface{}{
		"databaseID":    cr.jobSpec.ID,
		"externalJobID": cr.jobSpec.ExternalJobID,
		"name":          cr.jobSpec.Name.ValueOrZero(),
	}
	if id := cr.jobSpec.CronSpec.EVMChainID; id != nil {
		jobSpec["evmChainID"] = id.String()
	}

	vars := pipeline.NewVarsFrom(map[string]interface{}{
		"jobSpec": jobSpec,
		"jobRun": map[string]interface{}{
			"meta": map[string]interface{}{},
		},
	})

	run := pipeline.NewRun(*cr.jobSpec.PipelineSpec, vars)

	_, err := cr.pipelineRunner.Run(ctx, run, false, nil)
	if err != nil {
		cr.logger.Errorf("Error executing new run for jobSpec ID %v", cr.jobSpec.ID)
	}
}

func cronRunner() *cron.Cron {
	return cron.New(cron.WithSeconds())
}
