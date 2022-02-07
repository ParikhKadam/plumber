package config

import (
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"

	"github.com/batchcorp/plumber/server/types"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
)

var _ = Describe("Options", func() {
	Context("Connections", func() {
		cfg := &Config{
			ConnectionsMutex: &sync.RWMutex{},
			Connections:      make(map[string]*types.Connection),
		}
		It("Can save/get/delete", func() {
			id := uuid.NewV4().String()

			conn := &types.Connection{
				Connection: &opts.ConnectionOptions{
					XId: id,
				},
			}

			// Save
			cfg.SetConnection(conn.Connection.XId, conn)

			// Get
			stored := cfg.GetConnection(conn.Connection.XId)
			Expect(stored).To(Equal(conn))

			// Delete
			cfg.DeleteConnection(conn.Connection.XId)
			notThere := cfg.GetConnection(conn.Connection.XId)
			Expect(notThere).To(BeNil())
		})
	})

	Context("Services", func() {
		cfg := &Config{
			ServicesMutex: &sync.RWMutex{},
			Services:      make(map[string]*protos.Service),
		}
		It("Can save/get/delete", func() {
			id := uuid.NewV4().String()

			svc := &protos.Service{
				Id: id,
			}

			// Save
			cfg.SetService(svc.Id, svc)

			// Get
			stored := cfg.GetService(svc.Id)
			Expect(stored).To(Equal(svc))

			// Delete
			cfg.DeleteService(svc.Id)
			notThere := cfg.GetService(svc.Id)
			Expect(notThere).To(BeNil())
		})
	})

	Context("Relays", func() {
		cfg := &Config{
			RelaysMutex: &sync.RWMutex{},
			Relays:      make(map[string]*types.Relay),
		}
		It("Can save/get/delete", func() {
			id := uuid.NewV4().String()

			// TODO: this needs an ID
			relay := &types.Relay{
				Id:      id,
				Options: &opts.RelayOptions{},
			}

			// Save
			cfg.SetRelay(relay.Id, relay)

			// Get
			stored := cfg.GetRelay(relay.Id)
			Expect(stored).To(Equal(relay))

			// Delete
			cfg.DeleteRelay(relay.Id)
			notThere := cfg.GetRelay(relay.Id)
			Expect(notThere).To(BeNil())
		})
	})

	Context("Schemas", func() {
		cfg := &Config{
			SchemasMutex: &sync.RWMutex{},
			Schemas:      make(map[string]*protos.Schema),
		}
		It("Can save/get/delete", func() {
			id := uuid.NewV4().String()

			schema := &protos.Schema{
				Id: id,
			}

			// Save
			cfg.SetSchema(schema.Id, schema)

			// Get
			stored := cfg.GetSchema(schema.Id)
			Expect(stored).To(Equal(schema))

			// Delete
			cfg.DeleteSchema(schema.Id)
			notThere := cfg.GetSchema(schema.Id)
			Expect(notThere).To(BeNil())
		})
	})

	Context("Reads", func() {
		cfg := &Config{
			ReadsMutex: &sync.RWMutex{},
			Reads:      make(map[string]*types.Read),
		}
		It("Can save/get/delete", func() {
			id := uuid.NewV4().String()

			read := &types.Read{ReadOptions: &opts.ReadOptions{
				XId: id,
			}}

			// Save
			cfg.SetRead(read.ReadOptions.XId, read)

			// Get
			stored := cfg.GetRead(read.ReadOptions.XId)
			Expect(stored).To(Equal(read))

			// Delete
			cfg.DeleteRead(read.ReadOptions.XId)
			notThere := cfg.GetRead(read.ReadOptions.XId)
			Expect(notThere).To(BeNil())
		})
	})
})
