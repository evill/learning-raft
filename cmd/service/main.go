package main

import (
	raftController "github.com/evill/learning-raft/internal/core/controller"
	"github.com/evill/learning-raft/internal/core/node"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	pid := os.Getpid()
	log.Printf("Process ID: %d", pid)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	InitConfig()
	config := ReadConfig()

	nc := raftController.NewNodeController(config)
	go func() {
		sig := <-sigs
		log.Printf("Received signal: %v", sig)
		nc.Stop()
	}()

	nc.Start()
}

func InitConfig() {
	viper.SetEnvPrefix("RAFTIG")
	viper.AutomaticEnv()

	viper.BindEnv("NODE_INTERNAL_URI")
	viper.RegisterAlias("node.internal.uri", "NODE_INTERNAL_URI")
	//viper.SetDefault("node.uri", "http://:8080") // Host can be used as Id letter so no default value

	viper.BindEnv("API_INTERNAL_PORT")
	viper.RegisterAlias("api.internal.port", "API_INTERNAL_PORT")

	viper.BindEnv("PEERS_URIS")
	viper.RegisterAlias("peers.uris", "PEERS_URIS")
}

func ReadConfig() raftController.Config {
	// Try to unmarshal config from file
	config := raftController.Config{
		Api: raftController.ApiConfig{
			InternalPort: viper.GetUint16("api.internal.port"),
		},
		Node: node.Config{
			InternalUri: viper.GetString("node.internal.uri"),
		},
		Peers: node.PeersConfig{
			Uris: viper.GetStringSlice("peers.uris"),
		},
	}

	return config
}
