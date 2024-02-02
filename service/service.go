package service

type HelloService struct{}

func (s *HelloService) Hello(request *UserInfo, reply *UserInfo) error {
	reply.Name = request.GetName()

	return nil
}
