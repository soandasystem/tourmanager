package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"tourmanager/config"
	"tourmanager/core/models"
	"tourmanager/core/ports"
	"tourmanager/util"
)

type flowService struct {
	config       config.Config
	gatewaysRepo ports.GatewaysRepository
	salesRepo    ports.SaleRepository
	cursoRepo    ports.CursoRepository
	paymentRepo  ports.PaymentRepository
}

// NewFlowService creates a new flow service
func NewFlowService(cfg config.Config, gatewaysRepo ports.GatewaysRepository, salesRepo ports.SaleRepository, cursoRepo ports.CursoRepository, paymentRepo ports.PaymentRepository) ports.FlowService {
	return &flowService{
		config:       cfg,
		gatewaysRepo: gatewaysRepo,
		salesRepo:    salesRepo,
		cursoRepo:    cursoRepo,
		paymentRepo:  paymentRepo,
	}
}

func (s *flowService) InitPayment(ctx context.Context, req models.InitFlowPaymentReq) (models.InitFlowPaymentResp, error) {
	// 1. Buscar la venta
	saleFilter := map[string]interface{}{"id": req.SaleID}
	saleResult, err := s.salesRepo.Get(ctx, saleFilter, nil, nil)
	if err != nil {
		return models.InitFlowPaymentResp{}, fmt.Errorf("error fetching sale: %v", err)
	}
	if len(saleResult) == 0 {
		return models.InitFlowPaymentResp{}, fmt.Errorf("sale not found")
	}
	saleList, ok := saleResult[0].(models.SaleListResponse)
	if !ok || len(saleList.Items) == 0 {
		return models.InitFlowPaymentResp{}, fmt.Errorf("sale no encontrado en response")
	}

	// 2. Buscar el curso
	cursoFilter := map[string]interface{}{"id": req.CursoID}
	cursoResult, err := s.cursoRepo.Get(ctx, cursoFilter, nil, nil)
	if err != nil {
		return models.InitFlowPaymentResp{}, fmt.Errorf("error fetching curso: %v", err)
	}
	if len(cursoResult) == 0 {
		return models.InitFlowPaymentResp{}, fmt.Errorf("curso not found")
	}
	cursoList, ok := cursoResult[0].(models.CursoListResponse)
	if !ok || len(cursoList.Items) == 0 {
		return models.InitFlowPaymentResp{}, fmt.Errorf("curso no encontrado en response")
	}
	curso := cursoList.Items[0]

	// 2. Get Gateway Config
	gatewayFilter := map[string]interface{}{"company_id": req.CompanyID, "gateway_id": 3} // 3 = Flow
	gatewayResult, err := s.gatewaysRepo.Get(ctx, gatewayFilter, nil, nil)
	if err != nil {
		return models.InitFlowPaymentResp{}, fmt.Errorf("error fetching gateway: %v", err)
	}
	if len(gatewayResult) == 0 {
		return models.InitFlowPaymentResp{}, fmt.Errorf("gateway not found")
	}
	gatewayList, ok := gatewayResult[0].(models.GatewaysListResponse)
	if !ok || len(gatewayList.Items) == 0 {
		return models.InitFlowPaymentResp{}, fmt.Errorf("gateway no encontrado en response")
	}
	gateway := gatewayList.Items[0]

	flowAPIKey := gateway.AdditionalConfig.FlowAPIKey
	flowSecretKey := gateway.AdditionalConfig.FlowSecretKey

	// 3. Prepare Flow params
	optionalData := map[string]string{
		"venta":  strconv.FormatInt(req.SaleID, 10),
		"alumno": req.UserRut,
	}
	optionalJSON, _ := json.Marshal(optionalData)

	params := map[string]string{
		"commerceOrder":   req.Identificador,
		"subject":         "Pago cuota viaje estudio",
		"currency":        "CLP",
		"amount":          strconv.Itoa(req.Monto),
		"email":           curso.Correo,
		"paymentMethod":   "9",
		"urlConfirmation": "https://flowresponse.onrender.com/token", // Webhook
		"urlReturn":       "https://tu-dominio.com/api/returnFlow",   // Return URL
		"optional":        string(optionalJSON),
	}

	flowAPI := util.NewFlowAPI(flowAPIKey, flowSecretKey, s.config.FlowAPIURL)
	response, err := flowAPI.Send("payment/create", params, "POST")
	if err != nil {
		return models.InitFlowPaymentResp{}, fmt.Errorf("error comunicándose con Flow: %v", err)
	}

	flowURL, hasUrl := response["url"].(string)
	flowToken, hasToken := response["token"].(string)
	if !hasUrl || !hasToken {
		return models.InitFlowPaymentResp{}, fmt.Errorf("respuesta inválida de flow: %v", response)
	}

	destinationURL := fmt.Sprintf("%s?token=%s", flowURL, flowToken)
	return models.InitFlowPaymentResp{RedirectURL: destinationURL}, nil
}
