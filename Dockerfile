# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/youtubeWeb774inc
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/youtubeWeb774inc
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/youtubeWeb774inc