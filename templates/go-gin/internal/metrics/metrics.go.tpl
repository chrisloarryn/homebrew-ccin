{{- if .GCPProject}}
package metrics

import (
	"context"
	"fmt"
	"log"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	client     *monitoring.MetricClient
	projectID  string
)

// Initialize initializes the metrics client
func Initialize(gcpProjectID string) error {
	var err error
	projectID = gcpProjectID
	
	client, err = monitoring.NewMetricClient(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create monitoring client: %w", err)
	}

	return nil
}

// RecordHTTPRequest records an HTTP request metric
func RecordHTTPRequest(method, path string, statusCode int, duration time.Duration) {
	if client == nil {
		return
	}

	// Create a time series data point
	now := time.Now()
	dataPoint := &monitoringpb.Point{
		Interval: &monitoringpb.TimeInterval{
			EndTime: timestamppb.New(now),
		},
		Value: &monitoringpb.TypedValue{
			Value: &monitoringpb.TypedValue_DoubleValue{
				DoubleValue: float64(duration.Milliseconds()),
			},
		},
	}

	// Create the time series
	timeSeries := &monitoringpb.TimeSeries{
		Metric: &monitoringpb.Metric{
			Type: "custom.googleapis.com/{{.ProjectName}}/http_request_duration",
			Labels: map[string]string{
				"method":      method,
				"path":        path,
				"status_code": fmt.Sprintf("%d", statusCode),
			},
		},
		Resource: &monitoringpb.MonitoredResource{
			Type: "global",
		},
		Points: []*monitoringpb.Point{dataPoint},
	}

	// Create the request
	req := &monitoringpb.CreateTimeSeriesRequest{
		Name:       fmt.Sprintf("projects/%s", projectID),
		TimeSeries: []*monitoringpb.TimeSeries{timeSeries},
	}

	// Send the request (async to avoid blocking)
	go func() {
		if err := client.CreateTimeSeries(context.Background(), req); err != nil {
			log.Printf("Failed to write time series data: %v", err)
		}
	}()
}
{{- else}}
package metrics

import (
	"time"
)

// Initialize is a no-op when GCP is not enabled
func Initialize(gcpProjectID string) error {
	return nil
}

// RecordHTTPRequest is a no-op when GCP is not enabled
func RecordHTTPRequest(method, path string, statusCode int, duration time.Duration) {
	// No-op
}
{{- end}}
