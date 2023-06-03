# kubernetes-audit-log
Shipping Kubernetes Audit Logs to Slack

Based on the official documentation says “Kubernetes Audit logs provide a security-relevant, chronological set of records documenting the sequence of actions in a cluster.”, so by collecting and analyzing them we can answer these following questions:

- What happened?
- When did it happen?
- Who initiated it?
- On what did it happen?

The important thing we need to know is that request is recorded with an associated stage, these stages actually define the time during which audit logs should occur as below stages:

- `RequestReceived`: The stage for events generated as soon as the audit handler receives the request, and before it is delegated down the handler chain

- `ResponseStarted`: Once the response headers are sent, but before the response body is sent. This stage is only generated for long-running requests (e.g. watch)

- `ResponseComplete`: The response body has been completed and no more bytes will be sent

- `Panic`: Events generated when a panic occurred

### Notes:

- The Audit feature is disabled by default, so we will need to enable it in the API Server parameters, But we also need to defind the [Audit Policy](https://kubernetes.io/docs/tasks/debug/debug-cluster/audit/#audit-policy) because by default it sends huge JSON content and Audit Policy is needed to just configure what is actually needed as Audit content.

- The defined audit levels are:
  - None
  - Metadata: log request metadata (requesting user, timestamp, resource, verb, etc.) but not request or response body
  - Request: log event metadata and request body but not response body. This does not apply to non-resource requests
  - RequestResponse: log event metadata, request and response bodies. This does not apply to non-resource requests
- The simple Audit Policy may look like this:

```yaml
# Log all requests at the Metadata level.
apiVersion: audit.k8s.io/v1
kind: Policy
rules:
- level: Metadata
```

## Audit Backend:
- Audit backend, simply defines where we want to send audit logs. The API Server provides two backends to send logs and these are the following
  - Log Backend: writes events into the filesystem
  - Webhook backend: sends events to an external HTTP API
- To configure which backend enabled with the audit, it should be using the below flag:

```bash
## defines the details of the Audit Event, what they should include
--audit-policy-file=/etc/kubernetes/audit-policy.yaml
## specifies the log file path that log backend uses to write audit events
--audit-log-path=/var/log/audit.log
```

https://hooks.slack.com/services/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