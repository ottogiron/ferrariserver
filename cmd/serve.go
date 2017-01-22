package cmd

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/ferrariframework/ferrariserver/config"
	"github.com/ferrariframework/ferrariserver/grpc/gen"
	rpcservices "github.com/ferrariframework/ferrariserver/grpc/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	rpcPortKey = "rpc_port"
	//TLS
	certFileKey = "cert_file"
	keyFileKey  = "key_file"
	tlsKey      = "tls"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the ferrariworker server",
	Long:  `Starts the ferrariworker server`,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt(rpcPortKey)
		tls := viper.GetBool(tlsKey)
		certFile := viper.GetString(certFileKey)
		keyFile := viper.GetString(keyFileKey)

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			grpclog.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		if tls {
			creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
			if err != nil {
				grpclog.Fatalf("Failed to generate credentials %v", err)
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}
		grpcServer := grpc.NewServer(opts...)
		gen.RegisterJobServiceServer(grpcServer, rpcservices.NewJobService(config.JobService()))
		grpcServer.Serve(lis)

	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP(rpcPortKey, "r", 4051, "Port for the rpc service")
	serveCmd.Flags().BoolP(tlsKey, "t", false, "Connection uses TLS if true, else plain TCP")
	serveCmd.Flags().StringP(certFileKey, "c", "server.pem", "The TLS cert file")
	serveCmd.Flags().StringP(keyFileKey, "k", "server.key", "The TLS key file")

	viper.BindPFlag(rpcPortKey, serveCmd.Flags().Lookup(rpcPortKey))
	viper.BindPFlag(tlsKey, serveCmd.Flags().Lookup(tlsKey))
	viper.BindPFlag(certFileKey, serveCmd.Flags().Lookup(certFileKey))
	viper.BindPFlag(keyFileKey, serveCmd.Flags().Lookup(keyFileKey))
}
