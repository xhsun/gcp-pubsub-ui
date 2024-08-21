"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.usePubSubUiEchoMutation = exports.pubsubui = void 0;
const config_1 = require("../services/config");
const injectedRtkApi = config_1.emptySplitApi.injectEndpoints({
    endpoints: (build) => ({
        pubSubUiEcho: build.mutation({
            query: (queryArg) => ({
                url: `/pubsubui.PubSubUI/Echo`,
                method: "POST",
                body: queryArg.topicSubscription,
            }),
        }),
    }),
    overrideExisting: false,
});
exports.pubsubui = injectedRtkApi;
exports.usePubSubUiEchoMutation = injectedRtkApi.usePubSubUiEchoMutation;
