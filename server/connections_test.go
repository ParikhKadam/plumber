package server

//import (
//	"context"
//	"io/ioutil"
//	"sync"
//
//	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//	uuid "github.com/satori/go.uuid"
//	"github.com/sirupsen/logrus"
//
//	"github.com/batchcorp/plumber-schemas/build/go/protos"
//	"github.com/batchcorp/plumber-schemas/build/go/protos/common"
//
//	"github.com/batchcorp/plumber/config"
//	"github.com/batchcorp/plumber/embed/etcd/etcdfakes"
//)
//
//var logger = &logrus.Logger{Out: ioutil.Discard}
//
//var _ = Describe("Connection", func() {
//
//	var p *Server
//
//	BeforeEach(func() {
//		fakeEtcd := &etcdfakes.FakeIEtcd{}
//
//		p = &Server{
//			Bus:      fakeEtcd,
//			AuthToken: "batchcorp",
//			PersistentConfig: &config.Config{
//				ConnectionsMutex: &sync.RWMutex{},
//				Connections:      map[string]*opts.ConnectionOptions{},
//			},
//			Log: logrus.NewEntry(logger),
//		}
//	})
//
//	Context("GetConnection", func() {
//		It("check auth token", func() {
//			_, err := p.GetConnection(context.Background(), &protos.GetConnectionRequest{
//				Auth: &common.Auth{Token: "incorrect token"},
//			})
//
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(ContainSubstring(ErrInvalidToken.Error()))
//		})
//
//		It("returns a specific connection", func() {
//			connID := uuid.NewV4().String()
//
//			conn := &opts.ConnectionOptions{
//				Name:  "testing",
//				Notes: "test connection",
//				XId:   connID,
//				Conn: &opts.Conn_Kafka{Kafka: &conns.Kafka{
//					Address: []string{"127.0.0.1:9200"},
//				}},
//			}
//
//			p.PersistentConfig.SetConnection(connID, conn)
//
//			resp, err := p.GetConnection(context.Background(), &protos.GetConnectionRequest{
//				Auth:         &common.Auth{Token: "batchcorp"},
//				ConnectionId: connID,
//			})
//
//			Expect(err).ToNot(HaveOccurred())
//			Expect(resp.Connection).To(Equal(conn))
//		})
//	})
//
//	Context("GetAllConnections", func() {
//		It("check auth token", func() {
//			_, err := p.GetConnection(context.Background(), &protos.GetConnectionRequest{
//				Auth: &common.Auth{Token: "incorrect token"},
//			})
//
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(ContainSubstring(ErrInvalidToken.Error()))
//		})
//
//		It("returns all specific connections", func() {
//
//			for i := 0; i < 10; i++ {
//				conn := &opts.ConnectionOptions{
//					Name:  "testing",
//					Notes: "test connection",
//					XId:   uuid.NewV4().String(),
//					Conn: &opts.Conn_Kafka{Kafka: &conns.Kafka{
//						Address: []string{"127.0.0.1:9200"},
//					}},
//				}
//				p.PersistentConfig.SetConnection(conn.Id, conn)
//			}
//
//			resp, err := p.GetAllConnections(context.Background(), &protos.GetAllConnectionsRequest{
//				Auth: &common.Auth{Token: "batchcorp"},
//			})
//
//			Expect(err).ToNot(HaveOccurred())
//			Expect(len(resp.Connections)).To(Equal(10))
//		})
//	})
//
//	Context("TestConnection", func() {
//		It("check auth token", func() {
//			_, err := p.TestConnection(context.Background(), &protos.TestConnectionRequest{
//				Auth: &common.Auth{Token: "incorrect token"},
//			})
//
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(ContainSubstring(ErrInvalidToken.Error()))
//		})
//		It("tests connection", func() {
//			// TODO
//		})
//	})
//
//	Context("CreateConnection", func() {
//		It("check auth token", func() {
//			_, err := p.CreateConnection(context.Background(), &protos.CreateConnectionRequest{
//				Auth: &common.Auth{Token: "incorrect token"},
//			})
//
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(ContainSubstring(ErrInvalidToken.Error()))
//		})
//
//		It("creates a connection", func() {
//			conn := &opts.ConnectionOptions{
//				Name:  "testing",
//				Notes: "test connection",
//				Conn: &opts.Conn_Kafka{Kafka: &conns.Kafka{
//					Address: []string{"127.0.0.1:9200"},
//				}},
//			}
//
//			resp, err := p.CreateConnection(context.Background(), &protos.CreateConnectionRequest{
//				Auth:       &common.Auth{Token: "batchcorp"},
//				Connection: conn,
//			})
//
//			Expect(err).ToNot(HaveOccurred())
//			Expect(resp.ConnectionId).ToNot(BeEmpty())
//		})
//	})
//
//	Context("UpdateConnection", func() {
//		It("check auth token", func() {
//			_, err := p.UpdateConnection(context.Background(), &protos.UpdateConnectionRequest{
//				Auth: &common.Auth{Token: "incorrect token"},
//			})
//
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(ContainSubstring(ErrInvalidToken.Error()))
//		})
//
//		It("update a connection", func() {
//			connID := uuid.NewV4().String()
//
//			fakeEtcd := &etcdfakes.FakeIEtcd{}
//			p.Bus = fakeEtcd
//
//			conn := &opts.ConnectionOptions{
//				XId:   connID,
//				Name:  "testing",
//				Notes: "test connection",
//				Conn: &opts.Conn_Kafka{Kafka: &conns.Kafka{
//					Address: []string{"127.0.0.1:9200"},
//				}},
//			}
//
//			p.PersistentConfig.SetConnection(conn.Id, conn)
//
//			newConn := &opts.ConnectionOptions{
//				XId:   connID,
//				Name:  "updated",
//				Notes: "test connection",
//				Conn: &opts.Conn_Kafka{Kafka: &conns.Kafka{
//					Address: []string{"1.2.3.4:9200"},
//				}},
//			}
//
//			_, err := p.UpdateConnection(context.Background(), &protos.UpdateConnectionRequest{
//				Auth:          &common.Auth{Token: "batchcorp"},
//				ConnectionXId: connID,
//				Connection:    newConn,
//			})
//
//			updateConn := p.PersistentConfig.GetConnection(connID)
//
//			Expect(err).ToNot(HaveOccurred())
//			Expect(updateConn).To(Equal(newConn))
//			Expect(fakeEtcd.PutCallCount()).To(Equal(1))
//			Expect(fakeEtcd.PublishUpdateConnectionCallCount()).To(Equal(1))
//		})
//	})
//
//	Context("DeleteConnection", func() {
//		It("check auth token", func() {
//			_, err := p.DeleteConnection(context.Background(), &protos.DeleteConnectionRequest{
//				Auth: &common.Auth{Token: "incorrect token"},
//			})
//
//			Expect(err).To(HaveOccurred())
//			Expect(err.Error()).To(ContainSubstring(ErrInvalidToken.Error()))
//		})
//
//		It("deletes a connection", func() {
//			connID := uuid.NewV4().String()
//
//			fakeEtcd := &etcdfakes.FakeIEtcd{}
//			p.Bus = fakeEtcd
//
//			conn := &opts.ConnectionOptions{
//				XId:   connID,
//				Name:  "testing",
//				Notes: "test connection",
//				Conn: &opts.Conn_Kafka{Kafka: &conns.Kafka{
//					Address: []string{"127.0.0.1:9200"},
//				}},
//			}
//
//			p.PersistentConfig.SetConnection(conn.Id, conn)
//
//			resp, err := p.DeleteConnection(context.Background(), &protos.DeleteConnectionRequest{
//				Auth:          &common.Auth{Token: "batchcorp"},
//				ConnectionXId: connID,
//			})
//
//			Expect(err).ToNot(HaveOccurred())
//			Expect(resp.Status.Code).To(Equal(common.Code_OK))
//			Expect(fakeEtcd.DeleteCallCount()).To(Equal(1))
//			Expect(fakeEtcd.PublishDeleteConnectionCallCount()).To(Equal(1))
//		})
//	})
//})
