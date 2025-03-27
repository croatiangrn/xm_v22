package company

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestCompany_AssignAmountOfEmployees(t *testing.T) {
	type fields struct {
		ID                uuid.UUID
		Name              string
		Description       string
		AmountOfEmployees int
		Registered        bool
		Type              string
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Amount of imployees - success",
			fields: fields{
				AmountOfEmployees: 10,
			},
			args: args{
				amount: 10,
			},
			wantErr: false,
		},
		{
			name: "Amount of imployees - can't be negative",
			fields: fields{
				AmountOfEmployees: -10,
			},
			args: args{
				amount: -10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Company{
				ID:                tt.fields.ID,
				Name:              tt.fields.Name,
				Description:       tt.fields.Description,
				AmountOfEmployees: tt.fields.AmountOfEmployees,
				Registered:        tt.fields.Registered,
				Type:              tt.fields.Type,
				CreatedAt:         tt.fields.CreatedAt,
				UpdatedAt:         tt.fields.UpdatedAt,
			}
			if err := c.AssignAmountOfEmployees(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("AssignAmountOfEmployees() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
