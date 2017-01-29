package cmd

import (
	"github.com/ferrariframework/ferrariserver/config"
	"github.com/inconshreveable/log15"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:           "serve",
	Short:         "Starts the ferrariworker server",
	Long:          `Starts the ferrariworker server`,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		lis, err := config.TCPListener()
		if err != nil {
			log15.Crit("Failed to create listener", "err", err)
			return err
		}
		grpcServer, err := config.GRPCServer()

		if err != nil {
			log15.Crit("Failed to create a grpc server", "err", err)
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log15.Crit("Failed to start grpc server", "err", err)
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP(config.RPCPortKey, "r", 4051, "Port for the rpc service")
	serveCmd.Flags().BoolP(config.TLSKey, "t", false, "Connection uses TLS if true, else plain TCP")
	serveCmd.Flags().StringP(config.CertFileKey, "c", "server.pem", "The TLS cert file")
	serveCmd.Flags().StringP(config.KeyFileKey, "k", "server.key", "The TLS key file")
	serveCmd.Flags().Int64(config.RecordLogsIntervalKey, 500, "Interval to record logs to the underlying store in milliseconds")
	serveCmd.Flags().String(config.ElasticURLSKeys, "http://localhost:9200", "Coma separated list of elastic url's")
	serveCmd.Flags().Bool(config.ElasticSetSniffKey, false, "the elastic client  sniffes the cluster via the Nodes Info API")

	viper.BindPFlag(config.RPCPortKey, serveCmd.Flags().Lookup(config.RPCPortKey))
	viper.BindPFlag(config.TLSKey, serveCmd.Flags().Lookup(config.TLSKey))
	viper.BindPFlag(config.CertFileKey, serveCmd.Flags().Lookup(config.CertFileKey))
	viper.BindPFlag(config.KeyFileKey, serveCmd.Flags().Lookup(config.KeyFileKey))
	viper.BindPFlag(config.RecordLogsIntervalKey, serveCmd.Flags().Lookup(config.RecordLogsIntervalKey))
	viper.BindPFlag(config.ElasticURLSKeys, serveCmd.Flags().Lookup(config.ElasticURLSKeys))
	viper.BindPFlag(config.ElasticSetSniffKey, serveCmd.Flags().Lookup(config.ElasticSetSniffKey))
}
