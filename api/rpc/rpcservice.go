package rpc

import (
	"context"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	pb "github.com/Dmitriy-M1319/vacation-api/proto"
)

type RpcService struct {
	pb.UnimplementedVacationsServiceServer
	empRepo interfaces.EmployeeRepository
	vacRepo interfaces.VacationRepository
	// normRepo interfaces.VacationNormRepository
}

func NewRpcService(eRepo interfaces.EmployeeRepository,
	vRepo interfaces.VacationRepository) *RpcService {
	return &RpcService{empRepo: eRepo, vacRepo: vRepo}
}

func (sr *RpcService) GetAllEmployeers(ctx context.Context, req *pb.EmptyResponse) (*pb.ManyEmployeersResponse, error) {
	emps, err := sr.empRepo.GetAll()

	if err != nil {
		return nil, err
	}

	res := make([]*pb.Employee, len(emps))
	for i, e := range emps {
		var emp = &pb.Employee{
			Id:         int32(e.ID),
			FirstName:  e.FirstName,
			LastName:   e.LastName,
			Patronymic: e.Patronymic,
		}
		res[i] = emp
	}

	return &pb.ManyEmployeersResponse{Employees: res}, nil
}

func (sr *RpcService) GetEmployeeById(ctx context.Context, req *pb.EmployeeId) (*pb.EmployeeResponse, error) {
	emp, err := sr.empRepo.GetById(uint64(req.GetEmpId()))
	if err != nil {
		return nil, err
	}

	res := &pb.Employee{
		Id:         int32(emp.ID),
		FirstName:  emp.FirstName,
		LastName:   emp.LastName,
		Patronymic: emp.Patronymic,
	}

	return &pb.EmployeeResponse{Emp: res}, nil
}

func (sr *RpcService) InsertEmployee(ctx context.Context, res *pb.Employee) (*pb.EmployeeResponse, error) {
	val := &models.Employee{
		FirstName:  res.FirstName,
		LastName:   res.LastName,
		Patronymic: res.Patronymic,
	}
	err := sr.empRepo.Insert(val)

	if err != nil {
		return nil, err
	}

	return &pb.EmployeeResponse{Emp: &pb.Employee{
		Id:         int32(val.ID),
		FirstName:  val.FirstName,
		LastName:   val.LastName,
		Patronymic: val.Patronymic,
	}}, nil
}

func (sr *RpcService) UpdateEmployee(ctx context.Context, req *pb.UpdateEmpRequest) (*pb.EmployeeResponse, error) {
	val := &models.Employee{
		FirstName:  req.Emp.FirstName,
		LastName:   req.Emp.LastName,
		Patronymic: req.Emp.Patronymic,
	}

	err := sr.empRepo.Update(uint64(req.EmpId), val)

	if err != nil {
		return nil, err
	}

	result := &pb.Employee{
		Id:         req.EmpId,
		FirstName:  val.FirstName,
		LastName:   val.LastName,
		Patronymic: val.Patronymic,
	}

	return &pb.EmployeeResponse{Emp: result}, nil
}

func (sr *RpcService) DeleteEmployee(ctx context.Context, req *pb.EmployeeId) (*pb.EmptyResponse, error) {
	err := sr.empRepo.Delete(uint64(req.EmpId))

	if err != nil {
		return nil, err
	} else {
		return &pb.EmptyResponse{}, nil
	}
}

func (sr *RpcService) GetVacationsByEmployee(ctx context.Context, req *pb.EmployeeId) (*pb.ManyVacationsResponse, error) {
	vacations, err := sr.vacRepo.GetByEmployeeId(uint64(req.EmpId))

	if err != nil {
		return nil, err
	}

	result := make([]*pb.Vacation, len(vacations))
	for i, v := range vacations {
		result[i] = &pb.Vacation{Id: int32(v.ID), EmpId: int32(v.EmployeeID), StartDate: v.StartDate, EndDate: v.EndDate, DaysCount: int32(v.DaysCount)}
	}

	return &pb.ManyVacationsResponse{Vacations: result}, nil
}

func (sr *RpcService) GetVacationById(ctx context.Context, req *pb.VacationId) (*pb.VacationResponse, error) {
	v, err := sr.vacRepo.GetById(uint64(req.VacId))

	if err != nil {
		return nil, err
	}

	res := &pb.Vacation{Id: int32(v.ID), EmpId: int32(v.EmployeeID), StartDate: v.StartDate, EndDate: v.EndDate, DaysCount: int32(v.DaysCount)}

	return &pb.VacationResponse{Vac: res}, nil
}

func (sr *RpcService) InsertVacation(ctx context.Context, req *pb.Vacation) (*pb.VacationResponse, error) {
	val := &models.Vacation{
		EmployeeID: uint64(req.EmpId),
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		DaysCount:  uint64(req.DaysCount),
	}
	err := sr.vacRepo.Insert(val)

	if err != nil {
		return nil, err
	}

	return &pb.VacationResponse{Vac: &pb.Vacation{
		Id:        int32(val.ID),
		EmpId:     int32(val.EmployeeID),
		StartDate: val.StartDate,
		EndDate:   val.EndDate,
		DaysCount: int32(val.DaysCount),
	}}, nil
}

func (sr *RpcService) UpdateVacation(ctx context.Context, req *pb.UpdateVacRequest) (*pb.VacationResponse, error) {
	val := &models.Vacation{
		EmployeeID: uint64(req.Vac.EmpId),
		StartDate:  req.Vac.StartDate,
		EndDate:    req.Vac.EndDate,
		DaysCount:  uint64(req.Vac.DaysCount),
	}

	err := sr.vacRepo.Update(uint64(req.VacId), val)

	if err != nil {
		return nil, err
	}

	result := &pb.Vacation{
		Id:        int32(val.ID),
		EmpId:     int32(val.EmployeeID),
		StartDate: val.StartDate,
		EndDate:   val.EndDate,
		DaysCount: int32(val.DaysCount),
	}

	return &pb.VacationResponse{Vac: result}, nil
}

func (sr *RpcService) DeleteVacation(ctx context.Context, req *pb.VacationId) (*pb.EmptyResponse, error) {
	err := sr.vacRepo.Delete(uint64(req.VacId))

	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}
