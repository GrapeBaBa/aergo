/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package rpc

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/aergoio/aergo/types"

	"github.com/aergoio/aergo-actor/actor"
	"github.com/aergoio/aergo/config"
	"github.com/aergoio/aergo/message"
	"github.com/aergoio/aergo/pkg/component"
	aergorpc "github.com/aergoio/aergo/types"
	"google.golang.org/grpc"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/soheilhy/cmux"
)

// RPC is actor for providing rpc service
type RPC struct {
	conf *config.Config

	*component.BaseComponent

	hub *component.ComponentHub

	grpcServer    *grpc.Server
	grpcWebServer *grpcweb.WrappedGrpcServer
	actualServer  aergorpc.AergoRPCServiceServer
	httpServer    *http.Server

	ca types.ChainAccessor
}

//var _ component.IComponent = (*RPCComponent)(nil)

// NewRPC create an rpc service
func NewRPC(hub *component.ComponentHub, cfg *config.Config, chainAccessor types.ChainAccessor) *RPC {
	actualServer := &AergoRPCService{
		hub:       hub,
		msgHelper: message.GetHelper(),
	}
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 256),
	}

	grpcServer := grpc.NewServer(opts...)

	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
			return true
		}))

	rpcsvc := &RPC{
		conf: cfg,
		hub:  hub,

		grpcServer:    grpcServer,
		grpcWebServer: grpcWebServer,
		actualServer:  actualServer,
	}
	rpcsvc.BaseComponent = component.NewBaseComponent("rpc", rpcsvc, logger)
	actualServer.actorHelper = rpcsvc

	rpcsvc.httpServer = &http.Server{
		Handler:        rpcsvc.grpcWebHandlerFunc(grpcWebServer, http.DefaultServeMux),
		ReadTimeout:    4 * time.Second,
		WriteTimeout:   4 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return rpcsvc
}

// Start start rpc service.
func (ns *RPC) BeforeStart() {
	aergorpc.RegisterAergoRPCServiceServer(ns.grpcServer, ns.actualServer)
}

func (ns *RPC) AfterStart() {
	go ns.serve()
}

// Stop stops rpc service.
func (ns *RPC) BeforeStop() {
	ns.httpServer.Close()
	ns.grpcServer.Stop()
}

func (ns *RPC) Statistics() *map[string]interface{} {
	return nil
}

func (ns *RPC) Receive(context actor.Context) {
}

// Create HTTP handler that redirects matching requests to the grpc-web wrapper.
func (ns *RPC) grpcWebHandlerFunc(grpcWebServer *grpcweb.WrappedGrpcServer, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if grpcWebServer.IsAcceptableGrpcCorsRequest(r) || grpcWebServer.IsGrpcWebRequest(r) || grpcWebServer.IsGrpcWebSocketRequest(r) {
			grpcWebServer.ServeHTTP(w, r)
		} else {
			ns.Info().Msg("Request handled by other hanlder. is this correct?")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// Serve GRPC server over TCP
func (ns *RPC) serveGRPC(l net.Listener, server *grpc.Server) {
	if err := server.Serve(l); err != nil {
		switch err {
		case cmux.ErrListenerClosed:
			// sometimes this is occured when this is killed by signals
			// but this is normal case, not bug. so skip
		default:
			panic(err)
		}
	}
}

// Serve HTTP server over TCP
func (ns *RPC) serveHTTP(l net.Listener, server *http.Server) {
	if err := server.Serve(l); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

// Serve TCP multiplexer
func (ns *RPC) serve() {
	ipAddr := net.ParseIP(ns.conf.RPC.NetServiceAddr)
	if ipAddr == nil {
		panic("Wrong IP address format in RPC.NetServiceAddr")
	}

	addr := fmt.Sprintf("%s:%d", ipAddr, ns.conf.RPC.NetServicePort)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	// Setup TCP multiplexer
	tcpm := cmux.New(l)
	grpcL := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := tcpm.Match(cmux.HTTP1Fast())

	ns.Info().Msg(fmt.Sprintf("Starting RPC server listening on %s, with TLS: %v", addr, ns.conf.RPC.NSEnableTLS))

	// Server both servers
	go ns.serveGRPC(grpcL, ns.grpcServer)
	go ns.serveHTTP(httpL, ns.httpServer)

	// Serve TCP multiplexer
	if err := tcpm.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		ns.Fatal().Msg(fmt.Sprintf("%v", err))
	}

	return
}

const defaultTTL = time.Second * 4

// SendRequest implement interface method of ActorService
func (ns *RPC) SendRequest(actor string, msg interface{}) {
	ns.RequestTo(actor, msg)
}

// FutureRequest implement interface method of ActorService
func (ns *RPC) FutureRequest(actor string, msg interface{}) *actor.Future {
	return ns.RequestToFuture(actor, msg, defaultTTL)
}

// CallRequest implement interface method of ActorService
func (ns *RPC) CallRequest(actor string, msg interface{}) (interface{}, error) {
	future := ns.RequestToFuture(actor, msg, defaultTTL)

	return future.Result()
}

// GetChainAccessor implment interface method of ActorService
func (ns *RPC) GetChainAccessor() types.ChainAccessor {
	return ns.ca
}

func convertError(err error) types.CommitStatus {
	switch err {
	case nil:
		return types.CommitStatus_TX_OK
	case types.ErrTxNonceTooLow:
		return types.CommitStatus_TX_NONCE_TOO_LOW
	case types.ErrTxAlreadyInMempool:
		return types.CommitStatus_TX_ALREADY_EXISTS
	case types.ErrTxHasInvalidHash:
		return types.CommitStatus_TX_INVALID_HASH
	case types.ErrTxFormatInvalid:
		return types.CommitStatus_TX_INVALID_FORMAT
	case types.ErrInsufficientBalance:
		return types.CommitStatus_TX_INSUFFICIENT_BALANCE
	default:
		//logger.Info().Str("hash", err.Error()).Msg("RPC encountered unconvertable error")
		return types.CommitStatus_TX_INTERNAL_ERROR
	}
}
