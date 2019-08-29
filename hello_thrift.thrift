namespace go hello

struct HelloReq {
    1: string msg;
}

struct HelloRes {
    1: string msg;
}

service Hello {
    HelloRes echo(1: HelloReq req);
}
