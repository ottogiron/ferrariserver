package cmd

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"time"

	"github.com/ferrariframework/ferrariserver/config"
	"github.com/ferrariframework/ferrariserver/grpc/gen"
	rpcservices "github.com/ferrariframework/ferrariserver/grpc/services"
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	rpcPortKey = "rpc_port"
	//TLS
	certFileKey           = "cert_file"
	keyFileKey            = "key_file"
	tlsKey                = "tls"
	recordLogsIntervalKey = "record_logs_interval"
	elasticURLSKeys       = "elastic_urls"
	elasticSetSniffKey    = "elastic_set_sniff"
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
		recordLogsInterval := viper.GetInt64(recordLogsIntervalKey)
		elasticURLs := viper.GetString(elasticURLSKeys)
		elasticSetSniff := viper.GetBool(elasticSetSniffKey)

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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		logger := log15.New("app", "ferrari-worker")
		idGenerator, err := config.SnowFlakeGenerator()

		if err != nil {
			logger.Crit("Error creating idGenerator", err)
		}

		elasticClient, err := config.ElasticClient(elasticSetSniff, elasticURLs)

		if err != nil {
			logger.Crit("Error creating elastic client", err)
		}

		jobStore, err := config.JobStore(ctx, "job", "job", elasticClient, idGenerator)

		if err != nil {
			logger.Crit("Error creating jobStore", err)
		}

		jobLogStore, err := config.JobLogStore(ctx, "joblog", "joblog", elasticClient, idGenerator)

		if err != nil {
			logger.Crit("Error creating jobLogStore", err)
		}
		jobService := config.JobService(ctx, logger, jobStore, jobLogStore, true, time.Duration(recordLogsInterval))
		gen.RegisterJobServiceServer(grpcServer, rpcservices.NewJobService(jobService))
		grpcServer.Serve(lis)

	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP(rpcPortKey, "r", 4051, "Port for the rpc service")
	serveCmd.Flags().BoolP(tlsKey, "t", false, "Connection uses TLS if true, else plain TCP")
	serveCmd.Flags().StringP(certFileKey, "c", "server.pem", "The TLS cert file")
	serveCmd.Flags().StringP(keyFileKey, "k", "server.key", "The TLS key file")
	serveCmd.Flags().Int64(recordLogsIntervalKey, 500, "Interval to record logs to the underlying store in milliseconds")
	serveCmd.Flags().String(elasticURLSKeys, "http://localhost:9200", "Coma separated list of elastic url's")
	serveCmd.Flags().Bool(elasticSetSniffKey, false, "the elastic client  sniffes the cluster via the Nodes Info API")

	viper.BindPFlag(rpcPortKey, serveCmd.Flags().Lookup(rpcPortKey))
	viper.BindPFlag(tlsKey, serveCmd.Flags().Lookup(tlsKey))
	viper.BindPFlag(certFileKey, serveCmd.Flags().Lookup(certFileKey))
	viper.BindPFlag(keyFileKey, serveCmd.Flags().Lookup(keyFileKey))
	viper.BindPFlag(recordLogsIntervalKey, serveCmd.Flags().Lookup(recordLogsIntervalKey))
	viper.BindPFlag(elasticURLSKeys, serveCmd.Flags().Lookup(elasticURLSKeys))
	viper.BindPFlag(elasticSetSniffKey, serveCmd.Flags().Lookup(elasticSetSniffKey))
}
