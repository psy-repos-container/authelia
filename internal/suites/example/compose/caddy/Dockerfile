FROM caddy:2.11.2-builder@sha256:92adff27ec7e96a6dc42324c3b89a5b1402bfe9ea95b214639ab3015c4f4f6f5 AS builder

RUN xcaddy build fix-empty-copy-headers

FROM caddy:2.11.2@sha256:22e1d921a7dd98ea722ebd6819de785fd71abdab7f7fed8a2378e96d29bb923a

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
