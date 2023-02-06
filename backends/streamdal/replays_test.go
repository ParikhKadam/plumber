package streamdal

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
)

var _ = Describe("Replays", func() {
	defer GinkgoRecover()

	Context("listReplays", func() {
		It("returns an error when empty", func() {
			b := StreamdalWithMockResponse(200, `[]`)

			output, err := b.listReplays()
			Expect(err).To(Equal(errNoReplays))
			Expect(len(output)).To(Equal(0))
		})

		It("returns error on a bad response", func() {
			b := StreamdalWithMockResponse(200, `{}`)

			output, err := b.listReplays()
			Expect(err).To(Equal(errReplayListFailed))
			Expect(len(output)).To(Equal(0))
		})

		It("lists replays with collection source", func() {
			apiResponse := `[{
			  "id": "44da12e6-6dfd-4f11-a952-6863958acf05",
			  "name": "Test Replay",
			  "type": "single",
			  "query": "*",
			  "paused": false,
			  "collection": {"name": "Test Collection"},
			  "destination": {"name": "Test Destination"}
			}]`

			b := StreamdalWithMockResponse(200, apiResponse)

			output, err := b.listReplays()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(output)).To(Equal(1))
			Expect(output[0].ID).To(Equal("44da12e6-6dfd-4f11-a952-6863958acf05"))
			Expect(output[0].Name).To(Equal("Test Replay"))
			Expect(output[0].Type).To(Equal("Single"))
			Expect(output[0].Query).To(Equal("*"))
			Expect(output[0].Paused).To(BeFalse())
			Expect(output[0].Source).To(Equal("Collection - Test Collection"))
			Expect(output[0].Destination).To(Equal("Test Destination"))
		})

		It("lists replays with dead letter stage source", func() {
			apiResponse := `[{
			  "id": "44da12e6-6dfd-4f11-a952-6863958acf05",
			  "name": "Test Replay",
			  "type": "single",
			  "query": "*",
			  "paused": false,
			  "stage": {"name": "Test Stage"},
			  "destination": {"name": "Test Destination"}
			}]`

			b := StreamdalWithMockResponse(200, apiResponse)

			output, err := b.listReplays()
			Expect(err).ToNot(HaveOccurred())
			Expect(len(output)).To(Equal(1))
			Expect(output[0].ID).To(Equal("44da12e6-6dfd-4f11-a952-6863958acf05"))
			Expect(output[0].Name).To(Equal("Test Replay"))
			Expect(output[0].Type).To(Equal("Single"))
			Expect(output[0].Query).To(Equal("*"))
			Expect(output[0].Paused).To(BeFalse())
			Expect(output[0].Source).To(Equal("Dead Letter Stage - Test Stage"))
			Expect(output[0].Destination).To(Equal("Test Destination"))
		})
	})

	Context("Create Replay", func() {
		apiResponse := `{
    "id": "729e524d-6801-4660-8cc1-2c75999727ad",
    "name": "foo",
    "type": "single",
    "query": "*",
    "notes": "",
    "status": "running",
    "archived": false,
    "team": {
        "id": "fa2b3482-2a00-43ea-a0ed-341974a5f6ef",
        "name": "test-123"
    },
    "destination": {
        "id": "14a4af72-d119-457c-b572-863794a7e9e1",
        "name": "test pubsub",
        "notes": "",
        "type": "redis_streams",
        "metadata": {
            "address": "localhost:6379",
            "streams": [
                "plumbertest"
            ],
            "password": "test",
            "username": "mark"
        }
    },
    "collection": {
        "id": "cc512f2e-51e8-4a3a-994b-1dffc8d78108",
        "name": "newtest2",
        "notes": "",
        "token": "c745eae3-8ed5-47d7-a27e-a7d4c876257e",
        "paused": false,
        "schema_id": "a2dd7760-5907-4505-8813-42543b40ebec"
    },
    "author": {
        "id": "1503e318-9c87-4e86-ab6c-fda235cb4bde",
        "name": "mark test",
        "email": "mark-new@streamdal.com"
    },
    "inserted_at": "2021-04-01T14:35:00.840143Z",
    "updated_at": "2021-04-01T14:35:00.840143Z"
}`
		b := StreamdalWithMockResponse(200, apiResponse)
		replay, err := b.createReplay("*")

		Expect(err).ToNot(HaveOccurred())
		Expect(replay.ID).To(Equal("729e524d-6801-4660-8cc1-2c75999727ad"))
	})

	Context("generateReplayQuery", func() {
		It("handles wildcard", func() {
			b := &Streamdal{
				Opts: &opts.CLIOptions{
					Streamdal: &opts.StreamdalOptions{
						Create: &opts.StreamdalCreateOptions{
							Replay: &opts.StreamdalCreateReplayOptions{
								Name:          "",
								Type:          0,
								Notes:         "",
								CollectionId:  "",
								DestinationId: "",
								Query:         "*",
								FromTimestamp: "2021-04-01T12:13:14Z",
								ToTimestamp:   "2021-04-02T12:13:14Z",
							},
						},
					},
				},
			}

			query, err := b.generateReplayQuery()
			Expect(err).ToNot(HaveOccurred())
			Expect(query).To(Equal("batch.info.date_human: [2021-04-01T12:13:14Z TO 2021-04-02T12:13:14Z]"))
		})

		It("handles query with a field", func() {
			b := &Streamdal{
				Opts: &opts.CLIOptions{
					Streamdal: &opts.StreamdalOptions{
						Create: &opts.StreamdalCreateOptions{
							Replay: &opts.StreamdalCreateReplayOptions{
								Name:          "",
								Type:          0,
								Notes:         "",
								CollectionId:  "",
								DestinationId: "",
								Query:         "batch.team.id: 'foo'",
								FromTimestamp: "2021-04-01T12:13:14Z",
								ToTimestamp:   "2021-04-02T12:13:14Z",
							},
						},
					},
				},
			}

			query, err := b.generateReplayQuery()
			Expect(err).ToNot(HaveOccurred())
			Expect(query).To(Equal("batch.team.id: 'foo' AND batch.info.date_human: [2021-04-01T12:13:14Z TO 2021-04-02T12:13:14Z]"))
		})
	})
})
