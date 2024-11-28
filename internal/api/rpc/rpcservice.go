package rpc

import (
	"context"

	"github.com/Dmitriy-M1319/vacation-api/internal/models"
	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	pb "github.com/Dmitriy-M1319/vacation-api/pkg/vacation_api/v1"
)

type RpcService struct {
	pb.UnimplementedVacationsServiceServer
	empRepo  interfaces.EmployeeRepository
	vacRepo  interfaces.VacationRepository
	normRepo interfaces.VacationNormRepository
}

func NewRpcService(eRepo interfaces.EmployeeRepository,
	vRepo interfaces.VacationRepository,
	nRepo interfaces.VacationNormRepository) *RpcService {
	return &RpcService{empRepo: eRepo, vacRepo: vRepo, normRepo: nRepo}
}

func (sr *RpcService) GetAllEmployeers(ctx context.Context, req *pb.EmptyResponse) (*pb.ManyEmployeersResponse, error) {
	emps, err := sr.empRepo.GetAll()

	if err != nil {
		return nil, err
	}

	res := make([]*pb.Employee, len(emps))
	for i, e := range emps {
		var emp = &pb.Employee{
			Id:         e.ID,
			FirstName:  e.FirstName,
			LastName:   e.LastName,
			Patronymic: e.Patronymic,
		}
		res[i] = emp
	}

	return &pb.ManyEmployeersResponse{Employees: res}, nil
}

func (sr *RpcService) GetEmployeeById(ctx context.Context, req *pb.EmployeeId) (*pb.EmployeeResponse, error) {
	emp, err := sr.empRepo.GetById(uint64(req.GetId()))
	if err != nil {
		return nil, err
	}

	res := &pb.Employee{
		Id:         emp.ID,
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
		Id:         val.ID,
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

	err := sr.empRepo.Update(uint64(req.Id), val)

	if err != nil {
		return nil, err
	}

	result := &pb.Employee{
		Id:         req.Id,
		FirstName:  val.FirstName,
		LastName:   val.LastName,
		Patronymic: val.Patronymic,
	}

	return &pb.EmployeeResponse{Emp: result}, nil
}

func (sr *RpcService) DeleteEmployee(ctx context.Context, req *pb.EmployeeId) (*pb.EmptyResponse, error) {
	err := sr.empRepo.Delete(uint64(req.Id))

	if err != nil {
		return nil, err
	} else {
		return &pb.EmptyResponse{}, nil
	}
}

func (sr *RpcService) GetVacationsByEmployee(ctx context.Context, req *pb.EmployeeId) (*pb.ManyVacationsResponse, error) {
	vacations, err := sr.vacRepo.GetByEmployeeId(uint64(req.Id))

	if err != nil {
		return nil, err
	}

	result := make([]*pb.Vacation, len(vacations))
	for i, v := range vacations {
		result[i] = &pb.Vacation{Id: v.ID, EmpId: v.EmployeeID, StartDate: v.StartDate, EndDate: v.EndDate, DaysCount: v.DaysCount}
	}

	return &pb.ManyVacationsResponse{Vacations: result}, nil
}

func (sr *RpcService) GetVacationById(ctx context.Context, req *pb.VacationId) (*pb.VacationResponse, error) {
	v, err := sr.vacRepo.GetById(uint64(req.Id))

	if err != nil {
		return nil, err
	}

	res := &pb.Vacation{Id: v.ID, EmpId: v.EmployeeID, StartDate: v.StartDate, EndDate: v.EndDate, DaysCount: v.DaysCount}

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
		Id:        val.ID,
		EmpId:     val.EmployeeID,
		StartDate: val.StartDate,
		EndDate:   val.EndDate,
		DaysCount: val.DaysCount,
	}}, nil
}

func (sr *RpcService) UpdateVacation(ctx context.Context, req *pb.UpdateVacRequest) (*pb.VacationResponse, error) {
	val := &models.Vacation{
		EmployeeID: uint64(req.Vac.EmpId),
		StartDate:  req.Vac.StartDate,
		EndDate:    req.Vac.EndDate,
		DaysCount:  uint64(req.Vac.DaysCount),
	}

	err := sr.vacRepo.Update(uint64(req.Id), val)

	if err != nil {
		return nil, err
	}

	result := &pb.Vacation{
		Id:        val.ID,
		EmpId:     val.EmployeeID,
		StartDate: val.StartDate,
		EndDate:   val.EndDate,
		DaysCount: val.DaysCount,
	}

	return &pb.VacationResponse{Vac: result}, nil
}

func (sr *RpcService) DeleteVacation(ctx context.Context, req *pb.VacationId) (*pb.EmptyResponse, error) {
	err := sr.vacRepo.Delete(uint64(req.Id))

	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}

func (sr *RpcService) GetAllVacationNorms(ctx context.Context, req *pb.EmployeeResponse) (*pb.ManyVacationNormsResponse, error) {
	norms, err := sr.normRepo.GetAll()

	if err != nil {
		return nil, err
	}

	result := make([]*pb.VacationNorm, len(norms))
	for i, n := range norms {
		result[i] = &pb.VacationNorm{Id: n.ID, Month: n.Month, VacationsCount: n.VacationsCount}
	}

	return &pb.ManyVacationNormsResponse{Norms: result}, nil
}

func (sr *RpcService) GetVacationNormById(ctx context.Context, req *pb.VacationNormId) (*pb.VacationNormResponse, error) {
	n, err := sr.normRepo.GetById(uint64(req.Id))

	if err != nil {
		return nil, err
	}

	res := &pb.VacationNorm{Id: n.ID, Month: n.Month, VacationsCount: n.VacationsCount}

	return &pb.VacationNormResponse{Norm: res}, nil
}

func (sr *RpcService) InsertVacationNorm(ctx context.Context, req *pb.VacationNorm) (*pb.VacationNormResponse, error) {
	val := &models.VacationNorm{
		Month:          req.Month,
		VacationsCount: req.VacationsCount,
	}
	err := sr.normRepo.Insert(val)

	if err != nil {
		return nil, err
	}

	return &pb.VacationNormResponse{Norm: &pb.VacationNorm{
		Id:             val.ID,
		Month:          val.Month,
		VacationsCount: val.VacationsCount,
	}}, nil
}

func (sr *RpcService) UpdateVacationNorm(ctx context.Context, req *pb.UpdateVacNormRequest) (*pb.VacationNormResponse, error) {
	val := &models.VacationNorm{
		Month:          req.Norm.Month,
		VacationsCount: req.Norm.VacationsCount,
	}

	err := sr.normRepo.Update(req.Id, val)

	if err != nil {
		return nil, err
	}

	result := &pb.VacationNorm{
		Id:             val.ID,
		Month:          val.Month,
		VacationsCount: val.VacationsCount,
	}

	return &pb.VacationNormResponse{Norm: result}, nil
}

func (sr *RpcService) DeleteVacationNorm(ctx context.Context, req *pb.VacationNormId) (*pb.EmptyResponse, error) {
	err := sr.normRepo.Delete(req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.EmptyResponse{}, nil
}
