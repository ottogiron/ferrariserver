package config

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	elastic "gopkg.in/olivere/elastic.v3"

	"time"

	"github.com/ferrariframework/ferrariserver/grpc/gen"
	rpcservices "github.com/ferrariframework/ferrariserver/grpc/services"
	jobservice "github.com/ferrariframework/ferrariserver/services/job"
	"github.com/ferrariframework/ferrariserver/store"
	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	//RPCPortKey key for the grpc port configuration
	RPCPortKey = "rpc-port"
	//CertFileKey key for the grpc file key configuration
	CertFileKey = "cert-file"
	//KeyFileKey key for the grpc key file
	KeyFileKey = "key-file"
	//TLSKey key for the grpc tls flag
	TLSKey = "tls"
	//RecordLogsIntervalKey record logs  interval key
	RecordLogsIntervalKey = "record-logs-interval"
	//ElasticURLSKeys key for the elastic urls
	ElasticURLSKeys = "elastic-urls"
	//ElasticSetSniffKey key for the set sniff for elastic flag
	ElasticSetSniffKey = "elastic-set-sniff"
)

//GRPCServer creates a configured instance of a rpc server
func GRPCServer() (*grpc.Server, error) {

	tls := viper.GetBool(TLSKey)
	certFile := viper.GetString(CertFileKey)
	keyFile := viper.GetString(KeyFileKey)
	recordLogsInterval := viper.GetInt64(RecordLogsIntervalKey)
	elasticURLs := viper.GetString(ElasticURLSKeys)
	elasticSetSniff := viper.GetBool(ElasticSetSniffKey)

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

	elasticClient, err := ElasticClient(elasticSetSniff, elasticURLs)

	if err != nil {
		logger.Crit("Error creating elastic client", "error", err)
		return nil, errors.Wrap(err, "Failed to create elastic client")
	}

	jobStore, err := JobStore(ctx, "job", "job", elasticClient)

	if err != nil {

		return nil, errors.Wrap(err, "Failed to create JobStore")
	}

	jobLogStore, err := JobLogStore(ctx, "joblog", "joblog", elasticClient)

	if err != nil {
		return nil, errors.Wrap(err, "Failed to create JobLogStore")
	}
	jobService := JobService(ctx, logger, jobStore, jobLogStore, true, time.Duration(recordLogsInterval))
	gen.RegisterJobServiceServer(grpcServer, rpcservices.NewJobService(jobService))

	return grpcServer, nil
}

//TCPListener returns an initialized rpc listener in the viper configured port
func TCPListener() (net.Listener, error) {
	port := viper.GetInt(RPCPortKey)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to listen")
	}

	return lis, nil
}

//ElasticClient returns a new instance of an elastic client
func ElasticClient(setSniff bool, urls ...string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(urls...),
		elastic.SetSniff(setSniff),
	)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create elastic client urls=%v setSniff=%v", urls, setSniff)
	}
	return client, nil
}

//JobService Configures a new instance of a job service
func JobService(ctx context.Context, logger log15.Logger, jobStore store.Job, jobLogStore store.JobLog, recordLogs bool, recordLogsInterval time.Duration) jobservice.Service {
	clogger := logger.New("service", "job")

	return jobservice.New(
		jobservice.SetContext(ctx),
		jobservice.SetLogger(clogger),
		jobservice.SetJobStore(jobStore),
		jobservice.SetJobLogStore(jobLogStore),
		jobservice.SetRecordLogs(recordLogs),
		jobservice.SetRecordLogsInterval(recordLogsInterval),
	)
}
