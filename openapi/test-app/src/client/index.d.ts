declare const injectedRtkApi: import("@reduxjs/toolkit/query").Api<import("@reduxjs/toolkit/query").BaseQueryFn<string | import("@reduxjs/toolkit/query").FetchArgs, unknown, import("@reduxjs/toolkit/query").FetchBaseQueryError, {}, import("@reduxjs/toolkit/query").FetchBaseQueryMeta>, {
    pubSubUiEcho: import("@reduxjs/toolkit/query").MutationDefinition<PubSubUiEchoApiArg, import("@reduxjs/toolkit/query").BaseQueryFn<string | import("@reduxjs/toolkit/query").FetchArgs, unknown, import("@reduxjs/toolkit/query").FetchBaseQueryError, {}, import("@reduxjs/toolkit/query").FetchBaseQueryMeta>, never, TopicSubscription, "api">;
}, "api", never, typeof import("@reduxjs/toolkit/query").coreModuleName | typeof import("@reduxjs/toolkit/dist/query/react").reactHooksModuleName>;
export { injectedRtkApi as pubsubui };
export type PubSubUiEchoApiResponse = TopicSubscription;
export type PubSubUiEchoApiArg = {
    topicSubscription: TopicSubscription;
};
export type TopicSubscription = {
    gcpProjectId?: string;
    pubsubTopicName?: string;
};
export type GoogleProtobufAny = {
    /** The type of the serialized message. */
    "@type"?: string;
    [key: string]: any;
};
export type Status = {
    /** The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code]. */
    code?: number;
    /** A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. */
    message?: string;
    /** A list of messages that carry the error details.  There is a common set of message types for APIs to use. */
    details?: GoogleProtobufAny[];
};
export declare const usePubSubUiEchoMutation: import("@reduxjs/toolkit/dist/query/react/buildHooks").UseMutation<import("@reduxjs/toolkit/query").MutationDefinition<PubSubUiEchoApiArg, import("@reduxjs/toolkit/query").BaseQueryFn<string | import("@reduxjs/toolkit/query").FetchArgs, unknown, import("@reduxjs/toolkit/query").FetchBaseQueryError, {}, import("@reduxjs/toolkit/query").FetchBaseQueryMeta>, never, TopicSubscription, "api">>;
