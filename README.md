# demo-quicgo

## 今日やること

1. `quic-go/example` の動作確認
    - https://github.com/quic-go/quic-go
    - Server 上のコンテンツを Client (CLI, WebBrowser) で開けるか確認

2. quic-go を用いて，自前で http3 サーバを立ててみる
    - https://zenn.dev/satoken/articles/golang-hajimete-http3
        - `openssl`/`mkcert` を利用したキーペア作成
        - Wireshark を利用した Packet Capture（`keylog` ファイルを利用）
        - `qlog` ファイルの生成，`qviz` を用いたパケットの visualization
            - https://qvis.edm.uhasselt.be/#/files

3. （QUIC DATAGRAM を触ってみる）
    - https://tech.aptpod.co.jp/entry/2021/01/28/100000
        - SiDUCK サーバの実装
        - `quic-go` を用いた SiDUCK クライアントの実装
