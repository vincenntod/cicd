package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestUseCase_ShowCustomer(t *testing.T) {
	type fields struct {
		repo CustomerRepository
	}
	type args struct {
		c *gin.Context
	}
	mocking := NewCustomerRepository(t)
	mocking.On("ShowCustomer", mock.AnythingOfType("*gin.Context")).Return([]Customer{}, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Customer
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				repo: mocking,
			},
			args: args{
				c: &gin.Context{},
			},
			want:    []Customer{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.ShowCustomer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_ShowCustomerByName(t *testing.T) {
	type fields struct {
		repo CustomerRepository
	}
	type args struct {
		c *gin.Context
	}
	mocking := NewCustomerRepository(t)
	mocking.On("ShowCustomerByName", mock.AnythingOfType("*gin.Context")).Return([]Customer{}, nil)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Customer
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				repo: mocking,
			},
			args: args{
				c: &gin.Context{},
			},
			want:    []Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.ShowCustomerByName(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowCustomerByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowCustomerByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_AddCustomer(t *testing.T) {
	type fields struct {
		repo CustomerRepository
	}
	type args struct {
		c *gin.Context
	}
	mocking := NewCustomerRepository(t)
	mocking.On("AddCustomer", mock.AnythingOfType("*gin.Context")).Return(Customer{}, nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Customer
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				repo: mocking,
			},
			args: args{
				c: &gin.Context{},
			},
			want:    Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.AddCustomer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_EditCustomer(t *testing.T) {
	type fields struct {
		repo CustomerRepository
	}
	type args struct {
		c *gin.Context
	}
	mocking := NewCustomerRepository(t)
	mocking.On("EditCustomer", mock.AnythingOfType("*gin.Context")).Return(Customer{}, nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Customer
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				repo: mocking,
			},
			args: args{
				c: &gin.Context{},
			},
			want:    Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.EditCustomer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EditCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_DeleteCustomer(t *testing.T) {
	type fields struct {
		repo CustomerRepository
	}
	type args struct {
		c *gin.Context
	}
	mocking := NewCustomerRepository(t)
	mocking.On("DeleteCustomer", mock.AnythingOfType("*gin.Context")).Return(Customer{}, nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Customer
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				repo: mocking,
			},
			args: args{
				c: &gin.Context{},
			},
			want:    Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.DeleteCustomer(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUseCase(t *testing.T) {
	type args struct {
		repo CustomerRepository
	}
	mocking := NewCustomerRepository(t)
	mocking.On("UseCase", mock.AnythingOfType("*CustomerRepository")).Return(mocking, nil)
	tests := []struct {
		name string
		args args
		want *UseCase
	}{
		{
			name: "Test Case 1",
			args: args{
				repo: mocking,
			},
			want: &UseCase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUseCase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
