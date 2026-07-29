package ports

import (
	"context"
	"tourmanager/core/models"
)

type FlowService interface {
	InitPayment(ctx context.Context, req models.InitFlowPaymentReq) (models.InitFlowPaymentResp, error)
}
