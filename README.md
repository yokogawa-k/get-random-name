# ランダムな名前を生成してくれるプログラム

[![Build Status](https://travis-ci.org/yokogawa-k/get-random-name.svg?branch=master)](https://travis-ci.org/yokogawa-k/get-random-name)
[![Coverage Status](https://coveralls.io/repos/github/yokogawa-k/get-random-name/badge.svg?branch=master)](https://coveralls.io/github/yokogawa-k/get-random-name?branch=master)

### 特徴

[moby/pkg/namesgenerator](https://github.com/moby/moby/tree/master/pkg/namesgenerator)を利用しているが、heroku の名前をつけれるように区切り文字は "_" ではなく "-" を利用。

### 使い方

release から自分の環境にあったものをダウンロードしてきて実行

```console
$ wget https://github.com/yokogawa-k/get-random-name/releases/download/v0.0.1/get-random-name_linux_amd64
$ chmod +x get-random-name_linux_amd64
$ ./get-random-name_linux_amd64
quirky-pare
$ ./get-random-name_linux_amd64
festive-albattani
```

### 参考

[Dockerコンテナのおもしろい名前 | SOTA](http://deeeet.com/writing/2014/07/15/docker-container-name/) を参考に作成

