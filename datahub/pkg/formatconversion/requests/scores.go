package requests

import (
	DaoScoreTypes "github.com/turtacn/alameda/datahub/pkg/dao/interfaces/scores/types"
	DBCommon "github.com/turtacn/alameda/internal/pkg/database/common"
	ApiScores "github.com/turtacn/api/alameda_api/v1alpha1/datahub/scores"
)

type ListSimulatedSchedulingScoresRequestExtended struct {
	Request *ApiScores.ListSimulatedSchedulingScoresRequest
}

func (r *ListSimulatedSchedulingScoresRequestExtended) ProduceRequest() DaoScoreTypes.ListRequest {
	var (
		queryCondition DBCommon.QueryCondition
	)

	queryCondition = QueryConditionExtend{r.Request.GetQueryCondition()}.QueryCondition()
	listRequest := DaoScoreTypes.ListRequest{
		QueryCondition: queryCondition,
	}

	return listRequest
}
