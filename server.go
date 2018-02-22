//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
package scgost

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//
// TLS configuration.
//
type _TlsConfig struct {
	CertFile string
	KeyFile  string
}

//
// HTTP/s server.
//
type _SCServer struct {
	Port      int
	TlsConfig *_TlsConfig
}

//
// Creates a new HTTP server with default HTTP port.
//
func NewSCServer() *_SCServer {
	return &_SCServer{Port: HTTP_PORT}
}

//
// Creates a new HTTP server with an informed port.
//
func NewSCServerWithValues(__port int) *_SCServer {
	return &_SCServer{Port: __port}
}

//
// Creates a new HTTPS server with the informed parameters.
//
func NewSCServerTLS(__port int, __cert_file, __key_file string) *_SCServer {
	return &_SCServer{
		Port: __port,
		TlsConfig: &_TlsConfig{
			CertFile: __cert_file,
			KeyFile:  __key_file,
		},
	}
}

//
// Configure a certificate and a key for the TLS server.
//
func (__obj *_SCServer) ConfigureTLS(__cert_file, __key_file string) *_SCServer {
	__obj.TlsConfig = &_TlsConfig{
		CertFile: __cert_file,
		KeyFile:  __key_file,
	}
	return __obj
}

//
// Based on the values, it starts a HTTP/HTTPS server with the informed routes.
//
func (__obj *_SCServer) ListenAndServe(__router *mux.Router) error {

	__srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", __obj.Port),
		Handler:      __router,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	if __obj.TlsConfig != nil {

		_TLSConfig := &tls.Config{
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		__srv.TLSConfig = _TLSConfig
		__srv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

		return __srv.ListenAndServeTLS(__obj.TlsConfig.CertFile, __obj.TlsConfig.KeyFile)
	} else {

		return __srv.ListenAndServe()
	}

}

//
// ListenAndServe listens on the TCP network address srv.Addr and
// then calls Serve to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
func ListenAndServe(_router RouteFactory, __port int) error {

	return NewSCServerWithValues(__port).ListenAndServe(_router.Get())

}

//
// ListenAndServeTLS listens on the TCP network address srv.Addr and
// then calls Serve to handle requests on incoming TLS connections.
// Accepted connections are configured to enable TCP keep-alives.
//
func ListenAndServeTLS(_router RouteFactory, __port int, __cert_file, __key_file string) error {

	return NewSCServerTLS(__port, __cert_file, __key_file).ListenAndServe(_router.Get())

}
