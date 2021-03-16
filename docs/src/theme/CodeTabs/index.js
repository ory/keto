import React from 'react'
import Tabs from '@theme/Tabs'
import TabItem from '@theme/TabItem'
import CodeFromRemote from '../CodeFromRemote'

const CodeTabs = ({ sampleId, version }) => (
  <>
    <Tabs
      values={[
        { label: 'gRPC Go', value: 'grpc-go' },
        { label: 'gRPC node.js', value: 'grpc-nodejs' },
        { label: 'REST', value: 'rest' },
        { label: 'Keto Client CLI', value: 'cli' }
      ]}
      defaultValue="grpc-go"
    >
      <TabItem value="grpc-go">
        <CodeFromRemote
          src={`https://github.com/ory/keto/blob/${version}/contrib/docs-code-samples/${sampleId}/main.go`}
        />
      </TabItem>
      <TabItem value="grpc-nodejs">
        <CodeFromRemote
          src={`https://github.com/ory/keto/blob/${version}/contrib/docs-code-samples/${sampleId}/index.js`}
        />
      </TabItem>
      <TabItem value="rest">
        <CodeFromRemote
          src={`https://github.com/ory/keto/blob/${version}/contrib/docs-code-samples/${sampleId}/curl.sh`}
        />
      </TabItem>
      <TabItem value="cli">
        <CodeFromRemote
          src={`https://github.com/ory/keto/blob/${version}/contrib/docs-code-samples/${sampleId}/cli.sh`}
        />
      </TabItem>
    </Tabs>
    <CodeFromRemote
      src={`https://github.com/ory/keto/blob/${version}/contrib/docs-code-samples/${sampleId}/expected_output.txt`}
      title="Result"
    />
  </>
)

export default CodeTabs
