import React from 'react'
import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'
import CodeBlock from '@theme/CodeBlock'

const ReadRequest = ({ namespace, object, relation, subject }) => (
  <Tabs
    values={[
      {label: 'gRPC', key: 'grpc'},
      {label: 'REST', key: 'rest'},
      {label: 'Keto Client CLI', key: 'cli'}
    ]}
  >
    <TabItem value="grpc">
      <CodeBlock className="language-proto">
        ReadService.ListRelationTuples(ListRelationTuplesRequest)
      </CodeBlock>
    </TabItem>
  </Tabs>
)
