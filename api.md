# V1

## Context

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextDeleteResponse">V1ContextDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddResponse">V1ContextAddResponse</a>
- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextSearchResponse">V1ContextSearchResponse</a>

Methods:

- <code title="post /api/v1/context/delete">client.V1.Context.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextDeleteParams">V1ContextDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextDeleteResponse">V1ContextDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/context/add">client.V1.Context.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddParams">V1ContextAddParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddResponse">V1ContextAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/context/search">client.V1.Context.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextSearchParams">V1ContextSearchParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextSearchResponse">V1ContextSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Traces

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceListResponse">V1ContextTraceListResponse</a>
- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceDeleteResponse">V1ContextTraceDeleteResponse</a>

Methods:

- <code title="get /api/v1/context/traces">client.V1.Context.Traces.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceListParams">V1ContextTraceListParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceListResponse">V1ContextTraceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/context/traces/{traceId}/delete">client.V1.Context.Traces.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, traceID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextTraceDeleteResponse">V1ContextTraceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### View

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewGetResponse">V1ContextViewGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewDocsResponse">V1ContextViewDocsResponse</a>

Methods:

- <code title="get /api/v1/context/view">client.V1.Context.View.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewGetParams">V1ContextViewGetParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewGetResponse">V1ContextViewGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/context/view/docs">client.V1.Context.View.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewService.Docs">Docs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewDocsParams">V1ContextViewDocsParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextViewDocsResponse">V1ContextViewDocsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Memory

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextMemoryUpdateResponse">V1ContextMemoryUpdateResponse</a>

Methods:

- <code title="post /api/v1/context/memory/update">client.V1.Context.Memory.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextMemoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextMemoryUpdateParams">V1ContextMemoryUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextMemoryUpdateResponse">V1ContextMemoryUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/context/memory/delete">client.V1.Context.Memory.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextMemoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextMemoryDeleteParams">V1ContextMemoryDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### AddAsync

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncNewResponse">V1ContextAddAsyncNewResponse</a>
- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncCancelResponse">V1ContextAddAsyncCancelResponse</a>

Methods:

- <code title="post /api/v1/context/add-async">client.V1.Context.AddAsync.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncNewParams">V1ContextAddAsyncNewParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncNewResponse">V1ContextAddAsyncNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /api/v1/context/add-async/{id}/cancel">client.V1.Context.AddAsync.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncCancelResponse">V1ContextAddAsyncCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Status

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusGetResponse">V1ContextAddAsyncStatusGetResponse</a>
- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusListResponse">V1ContextAddAsyncStatusListResponse</a>

Methods:

- <code title="get /api/v1/context/add-async/{id}/status">client.V1.Context.AddAsync.Status.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusGetResponse">V1ContextAddAsyncStatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /api/v1/context/add-async/status">client.V1.Context.AddAsync.Status.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusListParams">V1ContextAddAsyncStatusListParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1ContextAddAsyncStatusListResponse">V1ContextAddAsyncStatusListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Org

### Context

Response Types:

- <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1OrgContextViewResponse">V1OrgContextViewResponse</a>

Methods:

- <code title="post /api/v1/org/context/view">client.V1.Org.Context.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1OrgContextService.View">View</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1OrgContextViewParams">V1OrgContextViewParams</a>) (<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang">alchemystai</a>.<a href="https://pkg.go.dev/github.com/Alchemyst-ai/alchemyst-sdk-golang#V1OrgContextViewResponse">V1OrgContextViewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
