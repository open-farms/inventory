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
	ErrFailedCreateEquipment = "failed to create equipment"
	ErrFailedListEquipment   = "failed to list equipment"
	ErrNoDifference          = "no differences found in update request"
)

func CreateEquipment(ctx context.Context, client *ent.Client, equipment ent.Equipment) (*ent.Equipment, error) {
	resp, err := client.Equipment.
		Create().
		SetName(equipment.Name).
		SetTags(equipment.Tags).
		SetCondition(equipment.Condition).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedCreateEquipment, err)
	}

	return resp, nil
}

func GetEquipment(ctx context.Context, client *ent.Client, id uint64) (*ent.Equipment, error) {
	resp, err := client.Equipment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DeleteEquipment(ctx context.Context, client *ent.Client, id uint64) (*ent.Equipment, error) {
	err := client.Equipment.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &ent.Equipment{
		ID: id,
	}, nil
}

func ListEquipments(ctx context.Context, client *ent.Client) ([]*ent.Equipment, error) {
	resp, err := client.Equipment.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedListEquipment, err)
	}

	return resp, nil
}

func UpdateEquipment(ctx context.Context, client *ent.Client, id uint64, equipment ent.Equipment) (*ent.Equipment, error) {
	prev, err := client.Equipment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if prev == &equipment {
		return nil, errors.New(ErrNoDifference)
	}

	update := client.Equipment.UpdateOneID(id)
	if prev.Name != equipment.Name && equipment.Name != "" {
		update.SetName(equipment.Name)
	}

	if !reflect.DeepEqual(prev.Tags, equipment.Tags) && !reflect.DeepEqual(equipment.Tags, []string{}) {
		update.SetTags(equipment.Tags)
	}

	if prev.Condition != equipment.Condition && equipment.Condition != "" {
		update.SetCondition(equipment.Condition)
	}

	update.SetUpdatedAt(time.Now())
	resp, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
