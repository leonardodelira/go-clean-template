package services

import (
	"context"
	"errors"
	"fmt"
	"leonardodelira/go-clean-template/internal/core/domain"
	"leonardodelira/go-clean-template/mocks/easymocks"
	mockupsGateway "leonardodelira/go-clean-template/mocks/mockups/gateway"
	mockupsRepo "leonardodelira/go-clean-template/mocks/mockups/repositorie"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type deps struct {
	translationRepo    *mockupsRepo.MockTranslationRepository
	translationGateway *mockupsGateway.MockTranslatorGateway
}

func makeDependencies(t *testing.T) deps {
	return deps{
		translationRepo:    mockupsRepo.NewMockTranslationRepository(gomock.NewController(t)),
		translationGateway: mockupsGateway.NewMockTranslatorGateway(gomock.NewController(t)),
	}
}

func TestDoTranslation(t *testing.T) {
	type args struct {
		input domain.TranslationInput
	}

	type want struct {
		response *domain.Translation
		err      error
	}

	tests := []struct {
		name string
		mock func(m deps, args *args, want *want)
		args func() *args
		want func() *want
	}{
		{
			name: "Success Translation and save postgres",
			mock: func(m deps, args *args, want *want) {
				resultGateway := easymocks.TranslationGatewayMock()
				gomock.InOrder(
					m.translationGateway.EXPECT().Translate(context.Background(), args.input.Text, args.input.LanguageDestination).Return(resultGateway, nil),
					m.translationRepo.EXPECT().SaveTranslation(context.Background(), resultGateway).Return(int32(1), nil),
				)
			},
			args: func() *args {
				return &args{
					input: domain.TranslationInput{
						Text:                "Olá Mundo!",
						LanguageDestination: "EN",
					},
				}
			},
			want: func() *want {
				resultGateway := easymocks.TranslationGatewayMock()
				return &want{
					response: resultGateway,
				}
			},
		},
		{
			name: "Error on translate in gateway",
			mock: func(m deps, args *args, want *want) {
				error := errors.New("any error gateway")
				gomock.InOrder(
					m.translationGateway.EXPECT().Translate(context.Background(), args.input.Text, args.input.LanguageDestination).Return(nil, error),
				)
			},
			args: func() *args {
				return &args{
					input: domain.TranslationInput{
						Text:                "Olá Mundo!",
						LanguageDestination: "EN",
					},
				}
			},
			want: func() *want {
				return &want{
					response: nil,
					err:      fmt.Errorf("error on translate in the gateway: any error gateway"),
				}
			},
		},
		{
			name: "Success Translation but error postgres",
			mock: func(m deps, args *args, want *want) {
				resultGateway := easymocks.TranslationGatewayMock()
				error := errors.New("any error postgres")
				gomock.InOrder(
					m.translationGateway.EXPECT().Translate(context.Background(), args.input.Text, args.input.LanguageDestination).Return(resultGateway, nil),
					m.translationRepo.EXPECT().SaveTranslation(context.Background(), resultGateway).Return(int32(-1), error),
				)
			},
			args: func() *args {
				return &args{
					input: domain.TranslationInput{
						Text:                "Olá Mundo!",
						LanguageDestination: "EN",
					},
				}
			},
			want: func() *want {
				return &want{
					response: nil,
					err:      fmt.Errorf("error on save translation on database: any error postgres"),
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, want := tt.args(), tt.want()
			deps := makeDependencies(t)
			tt.mock(deps, args, want)

			srv := NewTranslationService(deps.translationRepo, deps.translationGateway)
			r, err := srv.DoTranslation(context.Background(), args.input)

			assert.Equal(t, want.response, r)
			assert.Equal(t, want.err, err)
		})
	}
}

func TestGetTranslations(t *testing.T) {
	type args struct {
	}

	type want struct {
		response []domain.Translation
		err      error
	}

	tests := []struct {
		name string
		mock func(m deps, args *args, want *want)
		args func() *args
		want func() *want
	}{
		{
			name: "Success on GET all translations",
			mock: func(m deps, args *args, want *want) {
				gomock.InOrder(
					m.translationRepo.EXPECT().GetTranslations(gomock.Any()).Return(want.response, nil),
				)
			},
			args: func() *args {
				return &args{}
			},
			want: func() *want {
				translations := easymocks.TranslationRepoPGMock()
				return &want{
					response: translations,
					err:      nil,
				}
			},
		},
		{
			name: "Error on GET all translations",
			mock: func(m deps, args *args, want *want) {
				error := errors.New("any error postres")
				gomock.InOrder(
					m.translationRepo.EXPECT().GetTranslations(gomock.Any()).Return(want.response, error),
				)
			},
			args: func() *args {
				return &args{}
			},
			want: func() *want {
				return &want{
					response: nil,
					err:      fmt.Errorf("error on get all translations: any error postres"),
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, want := tt.args(), tt.want()
			deps := makeDependencies(t)
			tt.mock(deps, args, want)

			srv := NewTranslationService(deps.translationRepo, deps.translationGateway)
			r, err := srv.GetTranslations(context.Background())

			assert.Equal(t, want.response, r)
			assert.Equal(t, want.err, err)
		})
	}
}
