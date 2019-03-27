package parser

// func TestParse(t *testing.T) {
// 	asyncapi := []byte(`
// asyncapi: '2.0.0'
// id: myapi
// info:
//   title: My API
//   version: '1.0.0'
// channels:
//   mychannel:
//     publish:
//       message:
//         payload:
//           type: object
//           properties:
//             name:
//               type: string
// `)

// 	doc, err := Parse(asyncapi)
// 	assert.Assert(t, is.Nil(err))
// 	assert.Equal(t, doc.Asyncapi, "2.0.0")
// 	assert.Equal(t, doc.Id, "myapi")
// 	assert.Equal(t, doc.Info.Title, "My API")
// 	assert.Equal(t, doc.Info.Version, "1.0.0")
// 	assert.Equal(t, len(doc.Channels), 1)
// 	msgs, _ := doc.ListMessages()
// 	assert.Equal(t, len(msgs), 1)
// 	pm, _ := doc.ListProducedMessages()
// 	assert.Equal(t, len(pm), 1)
// 	cm, _ := doc.ListConsumedMessages()
// 	assert.Equal(t, len(cm), 0)
// }
