account:
	#-protoc  --proto_path=F:\workspace\goproject\ipfsdisk\service\account\rpc user.proto  \
	#--go_out=plugins=grpc:F:\workspace\goproject\ipfsdisk\service\account\rpc \
	#--go_opt=Muser.proto=./pb
	- goctl rpc proto -src F:\workspace\goproject\ipfsdisk\service\account\rpc\pb\user.proto -dir \
 		F:\workspace\goproject\ipfsdisk\service\account\rpc user.proto