package flogotimestamp

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-flogotimestamp")

type TimestampActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &TimestampActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *TimestampActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *TimestampActivity) Eval(context activity.Context) (done bool, err error) {

	inputdate, _ := context.GetInput("inputdate").(string)
	t, err := time.Parse(time.RFC3339, inputdate)
	if err != nil {
		log.Info(err)
	}
	context.SetOutput("timestamp", t.Unix())

	return true, nil
}
