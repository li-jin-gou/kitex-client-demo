// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello1");
    HelloResp HelloMethod1(1: HelloReq request) (api.get="/hello2");
    HelloResp HelloMethod2(1: HelloReq request) (api.get="/hello3");
}
