package app

import (
	"context"
	"linebot/app/shutdown"
	"linebot/infras"
	web "linebot/router"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5 // 讓程式最多等待 n 秒時間，如果超過 n 秒就強制關閉所有連線
)

// Server struct
type LinebotServer struct {
	infra         *infras.Options
	LinebotServer web.IRouter
}

// New Server constructor
// @param opts infrastructure options
// @param q1 router interface
// @result q1 server instance
func NewLineBotServer(opts *infras.Options, iWeb web.IRouter) *LinebotServer {
	opts.OnConfigChange()
	return &LinebotServer{infra: opts, LinebotServer: iWeb}
}

// region public methods
func (svc *LinebotServer) Run() error {
	port := svc.infra.Config.Server.Port
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        svc.LinebotServer.InitRouter(),
		ReadTimeout:    time.Second * svc.infra.Config.Server.ReadTimeout,
		WriteTimeout:   time.Second * svc.infra.Config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	go func() {
		svc.infra.Logger.Infof("Starting q1 server v%s which ip is %s:%s", svc.infra.Config.Server.AppVersion, svc.getOutboundIP(), port)
		if err := server.ListenAndServe(); err != nil {
			svc.infra.Logger.Errorf("Error starting Server: %s", err)
		}
	}()
	shutdown.Gracefully()
	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second) // 讓程式最多等待 5 秒時間，如果超過 5 秒就強制關閉所有連線
	defer shutdown()
	if err := server.Shutdown(ctx); err != nil { // 1. 關閉連接埠及2. 等待所有連線處理結束
		return err // handle err
	}
	svc.infra.Logger.Info("Server Exited Properly")
	return server.Shutdown(ctx)
}

// endregion

// region private methods
// Get preferred outbound ip of this machine
func (srv *LinebotServer) getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		srv.infra.Logger.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

// endregion
