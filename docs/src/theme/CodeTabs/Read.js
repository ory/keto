import React from 'react'
import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'
import CodeBlock from '@theme/CodeBlock'

const generateGoQuery = (params, indent) => {
  const entries = Object.entries(params).filter(([_, v]) => v)
  const longestKeyLen = entries.reduce((l, [k]) => l > k.length ? l : k.length, 0)

  return entries.reduce(
      (query, [key, value]) => query + "\n" + " ".repeat(indent + 2) + key.charAt(0).toUpperCase() + key.slice(1) + ": " + " ".repeat(longestKeyLen - key.length) + "\"" + value + "\",",
      ""
    )
}

const ReadRequest = (relationTuple) => (
  <Tabs
    values={[
      {label: 'gRPC Go', value: 'grpc-go'},
      {label: 'gRPC node.js', value: 'grpc-nodejs'},
      {label: 'REST', value: 'rest'},
      {label: 'Keto Client CLI', value: 'cli'}
    ]}
    defaultValue="grpc-go"
  >
    <TabItem value="grpc-go">
      <CodeBlock className="language-go">
        {`func Main() {
  conn := grpc.Dial(ketoRemote)
  client := acl.NewReadServiceClient(conn)

  resp, err := client.ListRelationTuples(context.Background(), &acl.ListRelationTuplesRequest{
    Query: &acl.ListRelationTuplesRequest_Query{${generateGoQuery(relationTuple, 4)}
    },
  })
}
`}
      </CodeBlock>
    </TabItem>
    <TabItem value="grpc-nodejs">
      <CodeBlock className="language-js">foo</CodeBlock>
    </TabItem>
    <TabItem value="rest">
      <CodeBlock className="language-js">foo</CodeBlock>
    </TabItem>
    <TabItem value="cli">
      <CodeBlock className="language-js">foo</CodeBlock>
    </TabItem>
  </Tabs>
)

export default ReadRequest
