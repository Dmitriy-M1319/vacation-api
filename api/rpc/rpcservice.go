package rpc

import (
	"context"

	"github.com/Dmitriy-M1319/vacation-api/internal/repository/interfaces"
	pb "github.com/Dmitriy-M1319/vacation-api/proto"
)

type RpcService struct {
	pb.UnimplementedVacationsServiceServer
	EmpRepo interfaces.EmployeeRepository
	// vacRepo  interfaces.VacationRepository
	// normRepo interfaces.VacationNormRepository
}

func (sr *RpcService) GetAllEmployeers(ctx context.Context, req *pb.EmptyResponse) (*pb.ManyEmployeersResponse, error) {
	emps, err := sr.EmpRepo.GetAll()

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
