import { CheckService } from "@ory/keto-grpc-client/ory/keto/v1beta/check_service_connectweb"
import {
  createConnectTransport,
  createPromiseClient,
} from "@bufbuild/connect-web";

const transport = createConnectTransport({
  baseUrl: "127.0.0.1:4466",
});
const checkClient = createPromiseClient(CheckService, transport);

try {
  const response = await checkClient.check({namespace: "messages", object: "02y_15_4w350m3", relation: "decypher"})
  console.log(response.allowed ? "Allowed" : "Denied")
} catch (error) {
  console.log("Encountered error:", error)
}
