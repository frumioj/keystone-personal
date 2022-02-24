package main

import (
	"context"
	"flag"
	"log"
	"net"
	"plugin"

	"google.golang.org/grpc"

	pb "github.com/regen-network/keystone2/keystone"
	krplugin "github.com/regen-network/keystone2/plugin"
	
)

func discoverKeyring( plugins []*krplugin.Plugin ) (*krplugin.Plugin, error){
	// return the first plugin as the keyring for now
	return plugins[0], nil
}

type pluginFlags []string

func (i *pluginFlags) String() string {
	return "a list of keyserving plugin paths"
}

func (i *pluginFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Server = pb.KeyringServer

type server struct {
	pb.UnimplementedKeyringServer
	ServerAddress string
	ChainID       string
	KeyringType   string
	KeyringDir    string
	RpcURI        string
	Plugins       []*krplugin.Plugin
}

func New() (Server, error) {
	return &server{}, nil
}

// NewKey implements the method given in the protobuf definition for
// the Keystone service (proto/keystone.proto)
func (s *server) NewKey(ctx context.Context, in *pb.KeySpec) (*pb.KeyRef, error) {
	log.Printf("Receive message body from client: %v", in)

	kr, err := discoverKeyring( s.Plugins )

	if err != nil {
		return nil, err
	}
	
	// spec := pb.KeySpec{
	// 	Label: "acbde12334",
	// 	Algo: pb.KeygenAlgorithm_KEYGEN_SECP256R1,
	// }
	
	ref, err := (*kr).NewKey(in)

	if err != nil {
		return nil, err
	} else {
		return ref, nil
	}
}

func (s *server) PubKey(ctx context.Context, in *pb.KeySpec) (*pb.PublicKey, error) {
	log.Printf("Receive message body from client: %v", in)

	kr, err := discoverKeyring( s.Plugins )

	if err != nil {
		return nil, err
	}
	
	// spec := pb.KeySpec{
	// 	Label: "acbde12334",
	// }
	
	ref, err := (*kr).PubKey(in)

	if err != nil {
		return nil, err
	} else {
		return ref, nil
	}
}

func (s *server) Sign(ctx context.Context, msg *pb.Msg) (*pb.Signed, error) {

	log.Printf("Receive message body from client: %v", msg)

	kr, err := discoverKeyring( s.Plugins )

	if err != nil {
		return nil, err
	}

	signed, err := (*kr).Sign( msg )

	log.Printf("Send message body to client: %v", signed)

	if err != nil {
		return nil, err
	} else {
		return signed, nil
	}

}

func main() {
	var plugins pluginFlags
	// Retrieve the command line parameters passed in to configure the server
	// Most have likely-reasonable defaults.
	keystoneAddress := flag.String("key-addr", "", "the address associated with the key used to sign transactions on behalf of Keystone")
	blockchain := flag.String("chain-id", "test-chain", "the blockchain that Keystone should connect to")
	keyringType := flag.String("keyring-type", "test", "the keyring backend type where keys should be read from")
	keyringDir := flag.String("keyring-dir", "~/.regen/", "the directory where the keys are")
	chainRpcURI := flag.String("chain-rpc", "tcp://localhost:26657", "the address of the RPC endpoint to communicate with the blockchain")
	grpcListenPort := flag.String("listen-port", "8080", "the port where the server will listen for connections")
	pkcs11KeyringConfig := flag.String("pkcsll-cfg", "./pkcs11-config", "configuration file for PKCS11 HSM connection")
	fileKeyringConfig := flag.String("file-cfg", "./keys", "configuration file for PKCS11 HSM connection")
	flag.Var(&plugins, "key-plugin", "one or more key-serving plugins")

	flag.Parse()

	if len(*keystoneAddress) <= 0 {
		log.Fatalln("Keystone server blockchain address may not be left empty")
		return
	}

	if len(plugins) <= 0 {
		log.Fatalln("At least one key-serving plugin libraries MUST be given with -plugin")
	}

	var pluginList []*krplugin.Plugin

	for _, s := range plugins {
		p, err := plugin.Open(s)
		if err != nil {
			log.Fatalf("Plugin could not be loaded from %s\n", s)
		}
		v, err := p.Lookup("TypeIdentifier")

		typeId, ok := v.(func() string)

		if !ok || len(typeId()) < 1 {
			log.Printf("No type identifier for the plugin, so not keeping it!")
		} else {

			v, err = p.Lookup("Init")

			var kr krplugin.Plugin = nil
			
			if err == nil &&
				typeId() == krplugin.Plugin_Type_File_Id {
				log.Printf("Init is %v of type %t", v, v)
				
				kr, err = v.(func(string) (kr krplugin.Plugin, err error))( *fileKeyringConfig)
			} else {
				if err == nil && typeId() == krplugin.Plugin_Type_Pkcs11_Id {
					kr, err = v.(func(string) (kr krplugin.Plugin, err error))( *pkcs11KeyringConfig)
				}
			}
			
			if err == nil {
				pluginList = append(pluginList, &kr)
			} else {
				// move on
			}
		}
	}

	lis, err := net.Listen("tcp", ":"+*grpcListenPort)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create new server context, used for passing server-global state
	ss := server{
		ServerAddress: *keystoneAddress,
		ChainID:       *blockchain,
		KeyringType:   *keyringType,
		KeyringDir:    *keyringDir,
		RpcURI:        *chainRpcURI,
		Plugins:       pluginList,
	}

	s := grpc.NewServer()
	pb.RegisterKeyringServer(s, &ss)

	s.Serve(lis)
	return

}
