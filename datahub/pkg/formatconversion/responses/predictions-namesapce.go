package responses

import (
	DaoPredictionTypes "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/predictions/types"
	FormatEnum "github.com/turtacn/alameda/datahub/pkg/formatconversion/enumconv"
	ApiPredictions "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/predictions"
)

type NamespacePredictionExtended struct {
	*DaoPredictionTypes.NamespacePrediction
}

func (d *NamespacePredictionExtended) ProducePredictions() *ApiPredictions.NamespacePrediction {
	var (
		rawDataChan        = make(chan ApiPredictions.MetricData)
		upperBoundDataChan = make(chan ApiPredictions.MetricData)
		lowerBoundDataChan = make(chan ApiPredictions.MetricData)
		numOfGoroutine     = 0

		datahubNamespacePrediction ApiPredictions.NamespacePrediction
	)

	datahubNamespacePrediction = ApiPredictions.NamespacePrediction{
		ObjectMeta: NewObjectMeta(&d.ObjectMeta),
	}

	// Handle prediction raw data
	numOfGoroutine = 0
	for metricType, samples := range d.PredictionRaw {
		if datahubMetricType, exist := FormatEnum.TypeToDatahubMetricType[metricType]; exist {
			numOfGoroutine++
			go producePredictionMetricDataFromSamples(datahubMetricType, samples.Granularity, samples.Data, rawDataChan)
		}
	}
	for i := 0; i < numOfGoroutine; i++ {
		receivedPredictionData := <-rawDataChan
		datahubNamespacePrediction.PredictedRawData = append(datahubNamespacePrediction.PredictedRawData, &receivedPredictionData)
	}

	// Handle prediction upper bound data
	numOfGoroutine = 0
	for metricType, samples := range d.PredictionUpperBound {
		if datahubMetricType, exist := FormatEnum.TypeToDatahubMetricType[metricType]; exist {
			numOfGoroutine++
			go producePredictionMetricDataFromSamples(datahubMetricType, samples.Granularity, samples.Data, upperBoundDataChan)
		}
	}
	for i := 0; i < numOfGoroutine; i++ {
		receivedPredictionData := <-upperBoundDataChan
		datahubNamespacePrediction.PredictedUpperboundData = append(datahubNamespacePrediction.PredictedUpperboundData, &receivedPredictionData)
	}

	// Handle prediction lower bound data
	numOfGoroutine = 0
	for metricType, samples := range d.PredictionLowerBound {
		if datahubMetricType, exist := FormatEnum.TypeToDatahubMetricType[metricType]; exist {
			numOfGoroutine++
			go producePredictionMetricDataFromSamples(datahubMetricType, samples.Granularity, samples.Data, lowerBoundDataChan)
		}
	}
	for i := 0; i < numOfGoroutine; i++ {
		receivedPredictionData := <-lowerBoundDataChan
		datahubNamespacePrediction.PredictedLowerboundData = append(datahubNamespacePrediction.PredictedLowerboundData, &receivedPredictionData)
	}

	return &datahubNamespacePrediction
}
