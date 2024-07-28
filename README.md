# GitHub Actions Sample

GitHub Actions を使ってみるリポジトリ。

## 内容

master ブランチへのプルリクエスト時に `go test` を行う。

## 設定ファイル

`.github/workflows` 配下を確認すること。

```text
$ tree -a -I "\.git|images"
.
├── .github
│   └── workflows
│       └── go-test.yml
├── README.md
├── go.mod
├── main.go
└── sqrt
    ├── sqrt.go
    └── sqrt_test.go
```

## テスト対象プログラム

Go を用いてニュートン法で平方根を求めるプログラム。

## テスト

### 成功例

![image](/images/success1.png)
![image](/images/success2.png)

### 失敗例

![image](/images/fail1.png)
![image](/images/fail2.png)
