package biz

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/open-farms/inventory/ent"
)

var (
	ErrFailedCreateVehicle = "failed to create vehicle %w"
	ErrFailedListVehicles  = "failed to list vehicles %w"
)

func CreateVehicle(ctx context.Context, client *ent.Client, vehicle ent.Vehicle) (*ent.Vehicle, error) {
	resp, err := client.Vehicle.
		Create().
		SetMake(vehicle.Make).
		SetModel(vehicle.Model).
		SetYear(vehicle.Year).
		SetActive(vehicle.Active).
		SetTags(vehicle.Tags).
		SetCondition(vehicle.Condition).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedCreateVehicle, err)
	}

	return resp, nil
}

func GetVehicle(ctx context.Context, client *ent.Client, id uint64) (*ent.Vehicle, error) {
	resp, err := client.Vehicle.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DeleteVehicle(ctx context.Context, client *ent.Client, id uint64) (*ent.Vehicle, error) {
	err := client.Vehicle.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &ent.Vehicle{
		ID: id,
	}, nil
}

func ListVehicles(ctx context.Context, client *ent.Client) ([]*ent.Vehicle, error) {
	resp, err := client.Vehicle.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedListVehicles, err)
	}

	return resp, nil
}

func UpdateVehicle(ctx context.Context, client *ent.Client, id uint64, vehicle ent.Vehicle) (*ent.Vehicle, error) {
	prev, err := client.Vehicle.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if prev == &vehicle {
		return nil, errors.New("no differences found in update request")
	}

	update := client.Vehicle.UpdateOneID(id)
	if prev.Make != vehicle.Make && vehicle.Make != "" {
		update.SetMake(vehicle.Make)
	}

	if prev.Model != vehicle.Model && vehicle.Model != "" {
		update.SetModel(vehicle.Model)
	}

	if prev.Year != vehicle.Year && vehicle.Year != "" {
		update.SetYear(vehicle.Year)
	}

	if prev.Active != vehicle.Active {
		update.SetActive(vehicle.Active)
	}

	if !reflect.DeepEqual(prev.Tags, vehicle.Tags) && !reflect.DeepEqual(vehicle.Tags, []string{}) {
		update.SetTags(vehicle.Tags)
	}

	if prev.Condition != vehicle.Condition && vehicle.Condition != "" {
		update.SetCondition(vehicle.Condition)
	}

	update.SetUpdatedAt(time.Now())
	resp, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
