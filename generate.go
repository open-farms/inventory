package generate

//go:generate kratos proto client api
//go:generate kratos proto server api/equipment/service/v1/equipment.proto -t internal/service
//go:generate kratos proto server api/vehicles/service/v1/vehicles.proto -t internal/service
