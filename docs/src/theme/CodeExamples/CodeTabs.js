import React from 'react'
import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'
import CodeFromRemote from "../CodeFromRemote";

const CodeTabs = ({sampleId}) =>
  <Tabs
    values={[
      {label: 'gRPC Go', value: 'grpc-go'},
      {label: 'gRPC node.js', value: 'grpc-nodejs'},
      {label: 'REST', value: 'rest'},
      {label: 'Keto Client CLI', value: 'cli'}
    ]}
    defaultValue="grpc-go">
    <TabItem value="grpc-go">
      <CodeFromRemote src={`https://github.com/ory/keto/blob/docs-guides/contrib/docs-code-samples/${sampleId}/main.go`}/>
    </TabItem>
  </Tabs>

export default CodeTabs
