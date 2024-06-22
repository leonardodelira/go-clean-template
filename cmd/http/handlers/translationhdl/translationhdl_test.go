package translationhdl

import (
	"encoding/json"
	"errors"
	"leonardodelira/go-clean-template/internal/core/domain"
	"leonardodelira/go-clean-template/mocks/easymocks"
	mockups "leonardodelira/go-clean-template/mocks/mockups/services"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type deps struct {
	serviceTranslation *mockups.MockTranslationService
}

func makeDependencies(t *testing.T) deps {
	return deps{
		serviceTranslation: mockups.NewMockTranslationService(gomock.NewController(t)),
	}
}

func TestGetTranslation(t *testing.T) {
	type args struct {
		body string
		url  string
	}

	type want struct {
		status int
	}

	tests := []struct {
		name string
		mock func(m deps, args *args, want *want)
		args func() *args
		want func() *want
	}{
		{
			name: "Success - 200",
			mock: func(m deps, args *args, want *want) {
				result := easymocks.TranslationMock()
				gomock.InOrder(
					m.serviceTranslation.EXPECT().GetTranslations(gomock.Any()).Return(result, nil),
				)
			},
			args: func() *args {
				return &args{
					url: "/v1/translation",
				}
			},
			want: func() *want {
				return &want{
					status: 200,
				}
			},
		},
		{
			name: "Error - Some error on service - 400",
			mock: func(m deps, args *args, want *want) {
				error := errors.New("some error occurs")
				gomock.InOrder(
					m.serviceTranslation.EXPECT().GetTranslations(gomock.Any()).Return(nil, error),
				)
			},
			args: func() *args {
				return &args{
					url: "/v1/translation",
				}
			},
			want: func() *want {
				return &want{
					status: 400,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//setup
			args, want := tt.args(), tt.want()
			deps := makeDependencies(t)
			tt.mock(deps, args, want)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", args.url, strings.NewReader(args.body))
			c.Request.Header.Set("Content-Type", "application/json")

			hdl := NewTranslationHandler(deps.serviceTranslation)
			hdl.GetTranslations(c)

			assert.Equal(t, want.status, w.Code)
			assert.NotNil(t, w.Body)
		})
	}
}

func TestDoTranslation(t *testing.T) {
	type args struct {
		body string
		url  string
	}

	type want struct {
		status int
	}

	tests := []struct {
		name string
		mock func(m deps, args *args, want *want)
		args func() *args
		want func() *want
	}{
		{
			name: "Success - 200",
			mock: func(m deps, args *args, want *want) {
				input := domain.TranslationInput{}
				_ = json.Unmarshal([]byte(args.body), &input)

				result := &easymocks.TranslationMock()[0]

				gomock.InOrder(
					m.serviceTranslation.EXPECT().DoTranslation(gomock.Any(), input).Return(result, nil),
				)
			},
			args: func() *args {
				input := `{
					"text": "Olá",
    				"language_destination": "EN"
				}`
				return &args{
					url:  "/v1/translation",
					body: input,
				}
			},
			want: func() *want {
				return &want{
					status: 200,
				}
			},
		},
		{
			name: "Error - Invalid Body - 400",
			mock: func(m deps, args *args, want *want) {
			},
			args: func() *args {
				input := `{
					"text": "Olá",
    				"language_destinati": "EN"
				}`

				return &args{
					url:  "/v1/translation",
					body: input,
				}
			},
			want: func() *want {
				return &want{
					status: 400,
				}
			},
		},
		{
			name: "Error - Some error on service - 400",
			mock: func(m deps, args *args, want *want) {
				input := domain.TranslationInput{}
				_ = json.Unmarshal([]byte(args.body), &input)

				error := errors.New("some error occurs")

				gomock.InOrder(
					m.serviceTranslation.EXPECT().DoTranslation(gomock.Any(), input).Return(nil, error),
				)
			},
			args: func() *args {
				input := `{
					"text": "Olá",
    				"language_destination": "EN"
				}`
				return &args{
					url:  "/v1/translation",
					body: input,
				}
			},
			want: func() *want {
				return &want{
					status: 400,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//setup
			args, want := tt.args(), tt.want()
			deps := makeDependencies(t)
			tt.mock(deps, args, want)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("POST", args.url, strings.NewReader(args.body))
			c.Request.Header.Set("Content-Type", "application/json")

			hdl := NewTranslationHandler(deps.serviceTranslation)
			hdl.DoTranslation(c)

			assert.Equal(t, want.status, w.Code)
			assert.NotNil(t, w.Body)
		})
	}
}
