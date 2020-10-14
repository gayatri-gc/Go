package main

import(
	"google.golang.org/grpc"
//	pb "home/neosoft/go/src/shippy/shippy-service-consignment/proto/consignment"
	pb "../shippy-service-consignment/proto/consignment"
	
)

const(
	address = "localhost:5001"
	defaultFileName = "consignment.json"
)

func parseFile(file)(*pb.Consignment, error){
	var consignment *pb.Consignment
	data, err := ioutils.ReadFile(file)
	if err != nil{
		fmt.Println("Error while reading file", err)
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		fmt.Println("couldnt make connection : ", err)
	}
	defer conn.Close()
	pb.NewShippingServiceClient(conn)

	file := defaultFileName
	consignment, err := parseFile(file)
	if err != nil{
		fmt.Println("Error parsing file ", err)
	}
	response, err := pb.CreateConsignment(context.Background(), consignment)
	if err != nil{
		fmt.Println("error while getting response ", err)
	}
	fmt.Println("Created :", response.Created)
}