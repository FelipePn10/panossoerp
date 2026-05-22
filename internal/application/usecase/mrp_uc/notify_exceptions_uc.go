package mrp_uc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/FelipePn10/panossoerp/internal/domain/mrp_calculation/entity"
	mrprepo "github.com/FelipePn10/panossoerp/internal/domain/mrp_calculation/repository"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/notification"
)

type NotifyExceptionsUseCase struct {
	repo     mrprepo.MRPCalculationRepository
	webhook  *notification.WebhookClient
	emailSvc *notification.EmailService
}

func NewNotifyExceptionsUseCase(repo mrprepo.MRPCalculationRepository, emailSvc *notification.EmailService) *NotifyExceptionsUseCase {
	return &NotifyExceptionsUseCase{
		repo:     repo,
		webhook:  notification.NewWebhookClient(),
		emailSvc: emailSvc,
	}
}

type NotifyExceptionsDTO struct {
	PlanCode   int64    `json:"plan_code"`
	WebhookURL string   `json:"webhook_url"`
	EmailTo    []string `json:"email_to"`
}

type ExceptionSummary struct {
	PlanCode    int64            `json:"plan_code"`
	GeneratedAt time.Time        `json:"generated_at"`
	Total       int              `json:"total"`
	ByType      map[string]int   `json:"by_type"`
	Exceptions  []*ExceptionItem `json:"exceptions"`
}

type ExceptionItem struct {
	ItemCode    int64  `json:"item_code"`
	MessageType string `json:"message_type"`
	Description string `json:"description"`
}

func (uc *NotifyExceptionsUseCase) Execute(ctx context.Context, dto NotifyExceptionsDTO) (*ExceptionSummary, error) {
	exceptions, err := uc.repo.ListExceptionsByPlan(ctx, dto.PlanCode)
	if err != nil {
		return nil, fmt.Errorf("listing exceptions for plan %d: %w", dto.PlanCode, err)
	}

	summary := buildSummary(dto.PlanCode, exceptions)

	if dto.WebhookURL != "" {
		if err := uc.webhook.Send(ctx, dto.WebhookURL, summary); err != nil {
			return summary, fmt.Errorf("webhook delivery failed: %w", err)
		}
	}

	if len(dto.EmailTo) > 0 && uc.emailSvc != nil {
		subject := fmt.Sprintf("[MRP] Exceções no Plano %d — %d ocorrência(s)", dto.PlanCode, summary.Total)
		if err := uc.emailSvc.Send(dto.EmailTo, subject, summary.FormatText()); err != nil {
			return summary, fmt.Errorf("email delivery failed: %w", err)
		}
	}

	return summary, nil
}

func buildSummary(planCode int64, exceptions []*entity.MRPExceptionMessage) *ExceptionSummary {
	byType := make(map[string]int)
	items := make([]*ExceptionItem, 0, len(exceptions))
	for _, e := range exceptions {
		t := string(e.MessageType)
		byType[t]++
		items = append(items, &ExceptionItem{
			ItemCode:    e.ItemCode,
			MessageType: t,
			Description: e.Description,
		})
	}
	return &ExceptionSummary{
		PlanCode:    planCode,
		GeneratedAt: time.Now(),
		Total:       len(exceptions),
		ByType:      byType,
		Exceptions:  items,
	}
}

// FormatText returns a plain-text digest suitable for email body or logs.
func (s *ExceptionSummary) FormatText() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Relatório de Exceções MRP — Plano %d\n", s.PlanCode))
	sb.WriteString(fmt.Sprintf("Gerado em: %s\n", s.GeneratedAt.Format("02/01/2006 15:04")))
	sb.WriteString(fmt.Sprintf("Total de exceções: %d\n", s.Total))
	sb.WriteString("\nPor tipo:\n")
	for t, count := range s.ByType {
		sb.WriteString(fmt.Sprintf("  %-30s %d\n", t, count))
	}
	sb.WriteString("\nDetalhes:\n")
	for _, item := range s.Exceptions {
		sb.WriteString(fmt.Sprintf("  Item %-6d [%-25s] %s\n", item.ItemCode, item.MessageType, item.Description))
	}
	return sb.String()
}
