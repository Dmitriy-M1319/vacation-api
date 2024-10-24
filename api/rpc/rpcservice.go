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
	// vacRepo  interfaces.VacationRepository
	// normRepo interfaces.VacationNormRepository
}

func NewRpcService(eRepo interfaces.EmployeeRepository) *RpcService {
	return &RpcService{empRepo: eRepo}
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

	//TODO: Пофиксить респонс для процедуры
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
